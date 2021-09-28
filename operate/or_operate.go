package operate

type orOperate struct {
	baseOperate
}

func init() {
	s := orOperate{}
	s.name = "or"
	s.symbol = []string{"||", "or"}
	s.regexMatch = "\\|\\||(?i:or)"
	s.priority = LogicPriority1
	s.execute = func(v1, v2 interface{}) interface{} {
		return ensureBool(v1) || ensureBool(v2)
	}

	Register(s.name, s)
}
