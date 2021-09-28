package operate

type sub struct {
	baseOperate
}

func init() {
	s := sub{}
	s.name = "sub"
	s.symbol = []string{"-", "sub"}
	s.regexMatch = "\\-|(?i:sub)"
	s.priority = Arithmetic1
	s.execute = func(v1, v2 interface{}) interface{} {
		return ensureFloat64(v1) - ensureFloat64(v2)
	}

	Register(s.name, s)
}
