package operate

import (
	"strconv"
	"strings"
)

const (
	_              int = iota
	LogicPriority1     // 逻辑运算1 or
	LogicPriority2     // 逻辑运算2 and
	LogicPriority3     // 逻辑运算3 > < >= <=
	LogicPriority4     // 逻辑运算4 !
	Arithmetic1        // 四则运算1 + -
	Arithmetic2        // 四则运算2 * / %
)

var globalOperateList []IOperate

func init() {
	Register(
		AddOperate{},
		subOperate{},
		MultiOperate{},
		divOperate{},
		AndOperate{},
		orOperate{},
		LogicOperate{},
		LogicNotOperate{},
	)
}

type IOperate interface {
	GetOperateSymbol() []string                                    // operate symbol
	Execute(op string, v1 interface{}, v2 interface{}) interface{} // execute func
	GetPriority() int
	GetRegexMatch() string //string parse by regex
	GetOperateNumber() int // 运算符数
}

//Register operate
func Register(operate ...IOperate) {
	globalOperateList = append(globalOperateList, operate...)
}

//GetOperate
//todo reflect
func GetOperate(symbol string) IOperate {
	for _, iOperate := range globalOperateList {
		for _, v := range iOperate.GetOperateSymbol() {
			if v == strings.ToLower(symbol) {
				return iOperate
			}
		}
	}
	return nil
}

func GetAllOperate() (operateList []IOperate) {
	for _, operate := range globalOperateList {
		operateList = append(operateList, operate)
	}
	return
}

type baseOperate struct {
}

func (b baseOperate) GetOperateNumber() int {
	// 默认都是双目运算符
	return 2
}

// todo reflect
func ensureFloat64(s interface{}) float64 {
	switch s.(type) {
	case string:
		floatValue, err := strconv.ParseFloat(s.(string), 64)
		if err != nil {
			panic(err)
		}
		return floatValue

	case float64:
		return s.(float64)
	case int:
		return float64(s.(int))
	case int32:
		return float64(s.(int32))
	case int64:
		return float64(s.(int64))
	}
	panic("value format err")
}

func ensureBool(s interface{}) bool {
	switch s.(type) {
	case string:
		boolValue, err := strconv.ParseBool(s.(string))
		if err != nil {
			panic(err)
		}
		return boolValue

	case bool:
		return s.(bool)
	}
	panic("value format err")
}
