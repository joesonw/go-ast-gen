package astgen

import (
	"go/ast"
)

func ParseInterface(expr *ast.InterfaceType, importedNames ...string) ([]Method, error) {
	var methods []Method
	for _, field := range expr.Methods.List {
		name := field.Names[0].Name
		function := field.Type.(*ast.FuncType)
		var comments []string
		if field.Comment != nil {
			for _, comment := range field.Comment.List {
				comments = append(comments, comment.Text)
			}
		}
		method := Method{
			name:     name,
			comments: comments,
		}
		for _, param := range function.Params.List {
			typ, err := ParseType(param.Type, importedNames...)
			if err != nil {
				return nil, err
			}
			name := ""
			if len(param.Names) > 0 {
				name = param.Names[0].Name
			}
			method.ins = append(method.ins, Field{
				name: name,
				typ:  typ,
			})
		}
		for _, out := range function.Results.List {
			typ, err := ParseType(out.Type, importedNames...)
			if err != nil {
				return nil, err
			}
			name := ""
			if len(out.Names) > 0 {
				name = out.Names[0].Name
			}
			method.outs = append(method.outs, Field{
				name: name,
				typ:  typ,
			})
		}
		methods = append(methods, method)
	}
	return methods, nil
}
