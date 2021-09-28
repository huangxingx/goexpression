package operate

var _ IOperate = MultiOperate{}

type MultiOperate struct {
}

func (a MultiOperate) GetRegexMatch() string {
	return "\\*"
}

func (a MultiOperate) GetPriority() int {
	return Arithmetic2
}

func (a MultiOperate) GetOperateSymbol() []string {
	return []string{"*"}
}

func (a MultiOperate) Execute(v1 interface{}, v2 interface{}) interface{} {
	float1 := ensureFloat64(v1)
	float2 := ensureFloat64(v2)
	return float1 * float2
}

func init() {
	Register("multi", MultiOperate{})
}
