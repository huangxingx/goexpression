package operate

var _ IOperate = AndOperate{}

type AndOperate struct {
}

func (a AndOperate) GetRegexMatch() string {
	return "and|&&"
}

func (a AndOperate) GetPriority() int {
	return LogicPriority2
}

func (a AndOperate) GetOperateSymbol() []string {
	return []string{"and"}
}

func (a AndOperate) Execute(v1 interface{}, v2 interface{}) interface{} {
	part1 := ensureBool(v1)
	part2 := ensureBool(v2)
	return part1 && part2
}

func init() {
	Register("and", AndOperate{})
}
