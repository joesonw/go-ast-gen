package astgen

type Field struct {
	name string
	typ Type
	comments []string
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Type() Type {
	return f.typ
}

func (f Field) Comments() []string {
	return f.comments[:]
}
