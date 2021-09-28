package operate

var _ IOperate = AddOperate{}

type AddOperate struct {
}

func (a AddOperate) GetRegexMatch() string {
	return "\\+"
}

func (a AddOperate) GetPriority() int {
	return Arithmetic1
}

func (a AddOperate) GetOperateSymbol() []string {
	return []string{"+"}
}

func (a AddOperate) Execute(v1 interface{}, v2 interface{}) interface{} {
	float1 := ensureFloat64(v1)
	float2 := ensureFloat64(v2)
	return float1 + float2
}

func init() {
	Register("increase", AddOperate{})
}
