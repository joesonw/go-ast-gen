package astgen

import (
	"go/ast"
)

func ParseInterface(expr *ast.InterfaceType, importer Importer) ([]Method, error) {
	var methods []Method
	for _, field := range expr.Methods.List {
		function := field.Type.(*ast.FuncType)
		var comments []string
		if field.Comment != nil {
			for _, comment := range field.Comment.List {
				comments = append(comments, comment.Text)
			}
		}

		var ins []Field
		var outs []Field
		for _, param := range function.Params.List {
			typ, err := ParseType(param.Type, importer)
			if err != nil {
				return nil, err
			}
			if len(param.Names) > 0 {
				for _, ident := range param.Names {
					ins = append(ins, Field{
						name: ident.Name,
						typ:  typ,
					})
				}
			} else {
				ins = append(ins, Field{
					typ: typ,
				})
			}
		}
		for _, out := range function.Results.List {
			typ, err := ParseType(out.Type, importer)
			if err != nil {
				return nil, err
			}
			if len(out.Names) > 0 {
				for _, ident := range out.Names {
					outs = append(outs, Field{
						name: ident.Name,
						typ:  typ,
					})
				}
			} else {
				outs = append(outs, Field{
					typ: typ,
				})
			}
		}
		for _, ident := range field.Names {
			methods = append(methods, Method{
				name:     ident.Name,
				comments: comments,
				ins:      ins[:],
				outs:     outs[:],
			})
		}
	}
	return methods, nil
}
