package operate

type divOperate struct {
	baseOperate
}

func (d divOperate) GetOperateSymbol() []string {
	return []string{"/", "div"}
}

func (d divOperate) Execute(op string, v1 interface{}, v2 interface{}) interface{} {
	return ensureFloat64(v1) / ensureFloat64(v2)
}

func (d divOperate) GetPriority() int {
	return Arithmetic2
}

func (d divOperate) GetRegexMatch() string {
	return "\\/"
}
