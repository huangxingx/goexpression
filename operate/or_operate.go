package operate

var _ IOperate = orOperate{}

type orOperate struct {
	baseOperate
}

func (o orOperate) GetOperateSymbol() []string {
	return []string{"||", "or"}
}

func (o orOperate) Execute(op string, v1 interface{}, v2 interface{}) interface{} {
	return ensureBool(v1) || ensureBool(v2)
}

func (o orOperate) GetPriority() int {
	return LogicPriority1
}

func (o orOperate) GetRegexMatch() string {
	return "\\|\\||(?i:or)"
}
