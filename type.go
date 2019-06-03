package astgen

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
	output := t.kind.String()
	if t.name != "" {
		output += "<" + t.name + ">"
	}
	for _, typ := range t.types {
		output += "(" + typ.String() + ")"
	}
	return output
}
