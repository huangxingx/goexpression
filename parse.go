package goexpression

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/huangxingx/goexpression/operate"
)

var keyWork = []string{"true", "false", "t", "f", "null"}

func isKeyWork(s string) bool {
	for _, v := range keyWork {
		if v == strings.ToLower(s) {
			return true
		}
	}
	return false
}

func genCompileByKeyWork() string {
	strList := make([]string, 0)
	for _, k := range keyWork {
		strList = append(strList, fmt.Sprintf("(?i:%s)", k))
	}
	return strings.Join(strList, "|")
}

func genCompileByOperate() string {
	strList := make([]string, 0)
	for _, iOperate := range operate.GetAllOperate() {
		if iOperate.GetRegexMatch() != "" {
			strList = append(strList, iOperate.GetRegexMatch())
		}
	}
	return strings.Join(strList, "|")
}

func parse2mpn(express string) []string {
	compileByKeyWork := genCompileByKeyWork()
	compileByOperateSymbol := genCompileByOperate()
	compile := regexp.MustCompile("\\(|\\)|\\d+\\.?\\d+|\\w+|" + compileByKeyWork + "|" + compileByOperateSymbol)
	return compile.FindAllString(express, -1)
}

func parseSuffixExpress(expressList []string) []string {
	suffixExpressList := make([]string, 0, len(expressList))
	stack := NewStack()
	for _, v := range expressList {
		// 字面量
		if IsNum(v) {
			suffixExpressList = append(suffixExpressList, v)
			continue
		}
		// v is op or var
		switch v {
		case "(":
			stack.Push(v)
		case ")":
			for stack.Peek() != "(" {
				suffixExpressList = append(suffixExpressList, stack.Pop().(string))
			}
			stack.Pop() // 移除 (
		default:
			// keyword
			if isKeyWork(v) {
				suffixExpressList = append(suffixExpressList, v)
				continue
			}

		cc:
			// check v is op?
			currentOperate := operate.GetOperate(v)
			if currentOperate == nil {
				//panic(fmt.Sprintf("不支持操作符: %s", currentOperate))
				// v is var
				suffixExpressList = append(suffixExpressList, v)
				break
			}

			// v is op
			if stack.IsEmpty() || stack.Peek().(string) == "(" || stack.Peek().(string) == ")" {
				stack.Push(v)
				break
			}
			top := stack.Peek().(string)

			topOperate := operate.GetOperate(top)
			if topOperate == nil {
				panic(fmt.Sprintf("不支持操作符: %s", top))
			}

			if currentOperate.GetPriority() > topOperate.GetPriority() {
				stack.Push(v)
				break
			} else {
				item := stack.Pop().(string)
				suffixExpressList = append(suffixExpressList, item)
				goto cc
			}
		}
	}
	for !stack.IsEmpty() {
		suffixExpressList = append(suffixExpressList, stack.Pop().(string))
	}
	return suffixExpressList
}
