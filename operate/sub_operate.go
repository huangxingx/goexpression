package operate

type subOperate struct {
	baseOperate
}

func (s subOperate) GetOperateSymbol() []string {
	return []string{"-", "sub"}
}

func (s subOperate) Execute(op string, v1 interface{}, v2 interface{}) interface{} {
	return ensureFloat64(v1) - ensureFloat64(v2)
}

func (s subOperate) GetPriority() int {
	return Arithmetic1
}

func (s subOperate) GetRegexMatch() string {
	return "\\-|(?i:sub)"
}
