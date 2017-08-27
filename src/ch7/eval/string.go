// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	str := strconv.FormatFloat(float64(l), 'g', 1, 64)
	return str
}

func (u unary) String() string {
	return fmt.Sprintf("%v%v", string(u.op), u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%v%v%v", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow", "min":
		return fmt.Sprintf("%v(%v, %v)", c.fn, c.args[0].String(), c.args[1].String())
	case "sin", "sqrt":
		return fmt.Sprintf("%v(%v)", c.fn, c.args[0].String())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
