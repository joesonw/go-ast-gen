package astgen

import "fmt"

var InvalidType = Type{
	kind: Invalid,
}

type Type struct {
	kind  Kind
	name  string
	types []Type
}

func (t Type) Name() string {
	return t.name
}

func (t Type) Kind() Kind {
	return t.kind
}

func (t Type) Types() []Type {
	return t.types[:]
}

func (t Type) String() string {
	if t.kind.IsFlat() {
		return t.kind.String()
	}
	switch t.kind {
	case Reference:
		return t.name
	case Pointer:
		return "*" + t.types[0].String()
	case Slice:
		return "[]" + t.types[0].String()
	case Map:
		return fmt.Sprintf("map[%s]%s", t.types[0].String(), t.types[1].String())
	case Chan:
		return "chan " + t.types[0].String()
	case Imported:
		return t.name + "." + t.types[0].String()
	default:
		return "INVALID!"
	}
}
