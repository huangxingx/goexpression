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

//func parse2mpn(expression string) []string {
//	result := make([]string, 0)
//	s := scanner.Scanner{}
//	s.Init(strings.NewReader(expression))
//	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
//		result = append(result, s.TokenText())
//	}
//	return result
//}

func parseSuffixExpression(expressionList []string) []string {
	suffixExpressionList := make([]string, 0, len(expressionList))
	stack := NewStack() // lower priority operate stack
	for _, v := range expressionList {
		// 字面量
		if IsNum(v) {
			suffixExpressionList = append(suffixExpressionList, v)
			continue
		}
		// v is op or var
		switch v {
		case "(":
			stack.Push(v)
		case ")":
			for stack.Peek() != "(" {
				suffixExpressionList = append(suffixExpressionList, stack.Pop().(string))
			}
			stack.Pop() // 移除 (
		default:
			// keyword
			if isKeyWork(v) {
				suffixExpressionList = append(suffixExpressionList, v)
				continue
			}

		cc:
			// check v is op?
			currentOperate := operate.GetOperate(v)
			if currentOperate == nil {
				//panic(fmt.Sprintf("不支持操作符: %s", currentOperate))
				// v is var
				suffixExpressionList = append(suffixExpressionList, v)
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
				suffixExpressionList = append(suffixExpressionList, item)
				goto cc
			}
		}
	}
	for !stack.IsEmpty() {
		suffixExpressionList = append(suffixExpressionList, stack.Pop().(string))
	}
	return suffixExpressionList
}
