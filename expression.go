package goexpression

import (
	"fmt"
	"strconv"

	"github.com/huangxingx/goexpression/operate"
)

//Expression ...
type Expression struct {
	inputExpression      string   // 表达式字符串
	suffixExpressionList []string // 后缀表达式
}

// NewExpress
// usage:
//	expressStr := "1+2+3"
//	exp := goexpression.NewExpress(expressStr)
//	result := exp.Execute(nil)
//	fmt.Printf("%s = %.2f", expressStr, result.(float64))
// // 6.00
func NewExpress(expression string) *Expression {
	mpn := parse2mpn(expression)
	suffixExpress := parseSuffixExpression(mpn)

	return &Expression{
		inputExpression:      expression,
		suffixExpressionList: suffixExpress,
	}
}

//Execute
// param map[string]interface{} 参数
func (e *Expression) Execute(param map[string]interface{}) interface{} {
	stack := NewStack()
	var tmpResult interface{}
	for _, v := range e.suffixExpressionList {
		// number or keywork
		if IsNum(v) || isKeyWork(v) {
			stack.Push(v)
			continue
		}
		// op
		iOperate := operate.GetOperate(v)
		if iOperate == nil {
			// var
			value, ok := param[v]
			if !ok {
				panic(fmt.Sprintf("var %s not value", v))
			}
			stack.Push(value)
			continue
		}
		// 单目运算符
		number := iOperate.GetOperateNumber()
		if number == 1 {
			oneValue := stack.Pop()
			tmpResult = iOperate.Execute(v, oneValue, nil)
		} else {
			v2 := stack.Pop()
			v1 := stack.Pop()
			tmpResult = iOperate.Execute(v, v1, v2)
		}

		stack.Push(tmpResult)
	}
	result := stack.Pop()
	if !stack.IsEmpty() {
		stack.Print()
		panic(fmt.Sprintf("execute err: stack.len > 0"))
	}

	return result
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
