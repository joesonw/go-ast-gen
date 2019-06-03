package astgen

import (
	"fmt"
	"go/ast"
	"reflect"
)

var builtInFlatKind = map[string]Kind{
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

func ParseStruct(expr *ast.StructType) ([]Field, error) {
	var fields []Field
	for _, field := range expr.Fields.List {
		name := field.Names[0].Name
		var comments []string
		if field.Comment != nil {
			for _, comment := range field.Comment.List {
				comments = append(comments, comment.Text)
			}
		}
		typ, err := ParseType(field.Type)
		if err != nil {
			return nil, err
		}
		fields = append(fields, Field{
			name:     name,
			typ:      typ,
			comments: comments,
		})

	}
	return fields, nil
}

func MustParseStruct(expr *ast.StructType) []Field {
	fields, err := ParseStruct(expr)
	if err != nil {
		panic(err)
	}
	return fields
}

func ParseType(expr ast.Expr) (Type, error) {
	switch expr.(type) {
	case *ast.Ident:
		{
			ident := expr.(*ast.Ident)
			if kind, ok := builtInFlatKind[ident.Name]; ok {
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
			t, err := ParseType(expr.(*ast.StarExpr).X)
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
			key, err := ParseType(m.Key)
			if err != nil {
				return key, err
			}
			value, err := ParseType(m.Value)
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
			ele, err := ParseType(array.Elt)
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
			ele, err := ParseType(ch.Value)
			if err != nil {
				return ele, err
			}
			return Type{
				kind:  Chan,
				types: []Type{ele},
			}, nil
		}
	default:
		return InvalidType, fmt.Errorf("type %s is not supported at this moment", reflect.TypeOf(expr).String())
	}
}
