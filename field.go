package astgen

import (
	"fmt"
	"strings"
)

type Field struct {
	name     string
	typ      Type
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

func (f Field) String() string {
	return fmt.Sprintf("%s %s %s", f.name, f.typ.String(), strings.Join(f.comments, "\n"))
}
