package astgen

import (
	"fmt"
	"strings"
)

type Method struct {
	name string
	ins  []Field
	outs []Field
}

func (m Method) Name() string {
	return m.name
}

func (m Method) Ins() []Field {
	return m.ins[:]
}

func (m Method) Outs() []Field {
	return m.outs[:]
}

func (m Method) String() string {
	ins := make([]string, len(m.ins))
	outs := make([]string, len(m.outs))
	for i, in := range m.ins {
		ins[i] = fmt.Sprintf("%s %s", in.name, in.typ.String())
	}
	for i, out := range m.outs {
		outs[i] = fmt.Sprintf("%s %s", out.name, out.typ.String())
	}
	return fmt.Sprintf("%s(%s) (%s)", m.name, strings.Join(ins, ","), strings.Join(outs, ","))
}
