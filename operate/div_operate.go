package operate

type divOperate struct {
	baseOperate
}

func init() {
	s := divOperate{}
	s.name = "div"
	s.symbol = []string{"/", "div"}
	s.regexMatch = "\\/"
	s.priority = Arithmetic2
	s.execute = func(v1, v2 interface{}) interface{} {
		return ensureFloat64(v1) / ensureFloat64(v2)
	}

	Register(s.name, s)
}
