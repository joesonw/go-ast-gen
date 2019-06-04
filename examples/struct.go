package main

import (
	"go/ast"
	"go/parser"
	"go/token"

	astgen "github.com/joesonw/go-ast-gen"
)

type structType struct {
	typ  *ast.StructType
	name string
}

func main() {
	var (
		input = `
		package test

		type AllFlatTypes struct {
			String string //test
			Bool bool
			Byte byte
			Int int
			Int8 int8
			Int16 int16
			Int32 int32
			Uint uint
			Uint16 uint16
			Uint32 uint32
			Uint64 uint64
			Float32 float32
			Float64 float64
			Complex64 complex64
			Complex128 complex128
			Time time.Time
		}

		type ReferenceStruct struct {
			A AllFlatTypes
			B *AllFlatTypes
		}

		type MapTypes struct {
			A map[string]string
			B map[string]int
			C map[string]ReferenceStruct
			D map[string]*ReferenceStruct
			E map[string]map[int]bool
		}

		type SliceTypes struct {
			A []string
			B [][]string
			C []ReferenceStruct
			D []*ReferenceStruct
		}

		type MixedMapSliceTypes struct {
			A map[int64][]string
			B []map[string]*ReferenceStruct
			C []map[string][]*ReferenceStruct
			D map[int64][]map[string]int64
		}

		type InterfaceTypes struct {
			A interface{}
			B []interface{}
			C map[interface{}]interface{}
		}

		type ChanTypes struct {
			A chan string 
			B []chan string
			C chan []string
			D chan chan string
			E map[string]chan string
			F chan map[string]string
			G chan map[string]chan string
		}
	`
	)

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", input, parser.AllErrors|parser.ParseComments)
	die(err)

	var structs []*structType
	for _, d := range f.Decls {
		decl, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		spec, ok := decl.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		if typ, ok := spec.Type.(*ast.StructType); ok {
			structs = append(structs, &structType{
				typ:  typ,
				name: spec.Name.Name,
			})
		}
	}

	for _, s := range structs {
		fields, err := astgen.ParseStruct(s.typ, "time.Time")
		die(err)

		println("")
		println("struct ", s.name, " {")
		for _, field := range fields {
			println("    ", field.String())
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
