package astgen

import (
	"fmt"
	"go/ast"
	"reflect"
)

var flatTypes = map[string]Kind{
	"string":     String,
	"bool":       Bool,
	"byte":       Byte,
	"int":        Int,
	"int8":       Int8,
	"int16":      Int16,
	"int32":      Int32,
	"int64":      Int64,
	"uint":       Uint,
	"uint8":      Uint8,
	"uint16":     Uint16,
	"uint32":     Uint32,
	"uint64":     Uint64,
	"float32":    Float32,
	"float64":    Float64,
	"complex64":  Complex64,
	"complex128": Complex128,
}

func ParseStruct(expr *ast.StructType, importer Importer) ([]Field, error) {
	var fields []Field
	for _, field := range expr.Fields.List {
		var comments []string
		if field.Comment != nil {
			for _, comment := range field.Comment.List {
				comments = append(comments, comment.Text)
			}
		}
		typ, err := ParseType(field.Type, importer)
		if err != nil {
			return nil, err
		}
		for _, ident := range field.Names {
			fields = append(fields, Field{
				name:     ident.Name,
				typ:      typ,
				comments: comments,
			})
		}
	}
	return fields, nil
}

func MustParseStruct(expr *ast.StructType, importer Importer) []Field {
	fields, err := ParseStruct(expr, importer)
	if err != nil {
		panic(err)
	}
	return fields
}

func ParseType(expr ast.Expr, importer Importer) (Type, error) {
	switch expr.(type) {
	case *ast.Ident:
		{
			ident := expr.(*ast.Ident)
			if kind, ok := flatTypes[ident.Name]; ok {
				return Type{
					kind: kind,
				}, nil
			} else {
				return Type{
					kind: Reference,
					name: ident.Name,
				}, nil
			}
		}
	case *ast.StarExpr:
		{
			t, err := ParseType(expr.(*ast.StarExpr).X, importer)
			if err != nil {
				return t, err
			}
			return Type{
				kind:  Pointer,
				types: []Type{t},
			}, nil
		}
	case *ast.MapType:
		{
			m := expr.(*ast.MapType)
			key, err := ParseType(m.Key, importer)
			if err != nil {
				return key, err
			}
			value, err := ParseType(m.Value, importer)
			if err != nil {
				return value, err
			}
			return Type{
				kind:  Map,
				types: []Type{key, value},
			}, nil
		}
	case *ast.ArrayType:
		{
			array := expr.(*ast.ArrayType)
			ele, err := ParseType(array.Elt, importer)
			if err != nil {
				return ele, err
			}
			return Type{
				kind:  Slice,
				types: []Type{ele},
			}, nil
		}
	case *ast.InterfaceType:
		{
			return Type{
				kind: Interface,
			}, nil
		}
	case *ast.ChanType:
		{
			ch := expr.(*ast.ChanType)
			ele, err := ParseType(ch.Value, importer)
			if err != nil {
				return ele, err
			}
			return Type{
				kind:  Chan,
				types: []Type{ele},
			}, nil
		}
	case *ast.SelectorExpr:
		{
			sel := expr.(*ast.SelectorExpr)
			pkg := sel.X.(*ast.Ident)
			fullName := fmt.Sprintf("%s.%s", pkg.Name, sel.Sel.Name)
			if !importer(fullName) {
				return InvalidType, fmt.Errorf("imported '%s' is not allowed", fullName)
			}
			return Type{
				kind: Imported,
				name: pkg.Name,
				types: []Type{{
					kind: Reference,
					name: sel.Sel.Name,
				}},
			}, nil
		}
	default:
		return InvalidType, fmt.Errorf("type %s is not supported at this moment", reflect.TypeOf(expr).String())
	}
}
