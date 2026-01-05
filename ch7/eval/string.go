package eval

import "strconv"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'g', -1, 64)
}

func (u unary) String() string {
	switch u.x.(type) {
	case Var, literal:
		return string(u.op) + u.x.String()
	default:
		return "(" + string(u.op) + u.x.String() + ")"
	}
}

func (b binary) String() string {
	switch b.op {
	case '+', '-':
		return "(" + b.x.String() + string(b.op) + b.y.String() + ")"
	default:
		return b.x.String() + string(b.op) + b.y.String()
	}
}

func (c call) String() string {
	if len(c.args) == 0 {
		return c.fn + "()"
	}
	str := c.fn + "("
	for _, eval := range c.args {
		str += eval.String() + ","
	}
	str = str[:len(str)-1] + ")"
	return str
}

func (m minExp) String() string {
	str := "MIN" + "("
	for _, eval := range m.args {
		str += eval.String() + ","
	}
	str = str[:len(str)-1] + ")"
	return str
}
