package operate

import "fmt"

//LogicOperate logic ">", ">=", "<", "<="
type LogicOperate struct {
	baseOperate
}

func (l LogicOperate) GetOperateSymbol() []string {
	return []string{">", ">=", "<", "<="}
}

func (l LogicOperate) Execute(op string, v1 interface{}, v2 interface{}) interface{} {
	switch op {
	case ">":
		return ensureFloat64(v1) > ensureFloat64(v2)
	case "<":
		return ensureFloat64(v1) < ensureFloat64(v2)
	case ">=":
		return ensureFloat64(v1) >= ensureFloat64(v2)
	case "<=":
		return ensureFloat64(v1) <= ensureFloat64(v2)
	}
	panic(fmt.Sprintf("op err: %s", op))
}

func (l LogicOperate) GetPriority() int {
	return LogicPriority3
}

func (l LogicOperate) GetRegexMatch() string {
	return "<=|>=|>|<"
}

//LogicNotOperate logic !
type LogicNotOperate struct {
	baseOperate
}

func (l LogicNotOperate) GetOperateNumber() int {
	return 1
}

func (l LogicNotOperate) GetOperateSymbol() []string {
	return []string{"!"}
}

func (l LogicNotOperate) Execute(op string, v1 interface{}, v2 interface{}) interface{} {
	return !ensureBool(v1)
}

func (l LogicNotOperate) GetPriority() int {
	return LogicPriority4
}

func (l LogicNotOperate) GetRegexMatch() string {
	return "!"
}
