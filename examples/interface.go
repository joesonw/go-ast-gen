package main

import (
	"go/ast"
	"go/parser"
	"go/token"

	astgen "github.com/joesonw/go-ast-gen"
)

type interfaceType struct {
	typ  *ast.InterfaceType
	name string
}

func main() {
	var (
		input = `
		package test
		import "time"

		type StructA struct {}

		type A interface {
			A(a *StructA, b time.Time) (int64, error) //test
			B(arr [][]string) error
		}
	`
	)

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", input, parser.AllErrors|parser.ParseComments)
	die(err)

	var interfaces []*interfaceType
	for _, d := range f.Decls {
		decl, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		spec, ok := decl.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		if typ, ok := spec.Type.(*ast.InterfaceType); ok {
			interfaces = append(interfaces, &interfaceType{
				typ:  typ,
				name: spec.Name.Name,
			})
		}
	}

	for _, in := range interfaces {
		methods, err := astgen.ParseInterface(in.typ, func(name string) bool { return true })
		die(err)

		println("")
		println("interface ", in.name, " {")
		for _, method := range methods {
			println("    ", method.String())
		}
		println("}")
		println("")
	}

}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
