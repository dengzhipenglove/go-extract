package goextract

import (
	"go/ast"
	"go/constant"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
)

type ConstIdentItem struct {
	Name        string
	TypeName    string
	Value       int64
	ValueString string
	IsInteger   bool // interger or string
	Comment     string
}

func ExtractGoFileConst(filePath string, typeName string) (string, []*ConstIdentItem, error) {
	var pkgName string
	var res = []*ConstIdentItem{}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	info := types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check(".", fset, []*ast.File{astFile}, &info)
	if err != nil {
		return pkgName, nil, err
	}

	pkgName = pkg.Name()

	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			vspec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			var typ string
			if vspec.Type == nil && len(vspec.Values) > 0 {
				typ = ""
				//  const OK = T(2)
				ce, ok := vspec.Values[0].(*ast.CallExpr)
				if ok {
					id, ok := ce.Fun.(*ast.Ident)
					if ok {
						typ = id.Name
					}
				}

			}
			if vspec.Type != nil {
				// "X T". We have a type. Remember it.
				ident, ok := vspec.Type.(*ast.Ident)
				if ok {
					typ = ident.Name
				}

			}

			if typeName != "" && typeName != typ {
				continue
			}

			// extra const data
			for _, name := range vspec.Names {
				if name.Name == "_" {
					continue
				}

				obj, ok := info.Defs[name]
				if !ok {
					panic("fatal: obj not exist" + name.Name)
				}

				comment := ""
				if vspec.Comment != nil {
					comment = vspec.Comment.Text()
				}

				resItem := ConstIdentItem{}

				kst := obj.(*types.Const)

				basic := obj.Type().Underlying().(*types.Basic)

				resItem.Name = name.Name
				resItem.ValueString = kst.Val().ExactString()
				resItem.TypeName = typ

				if basic.Info()&types.IsInteger > 0 {
					resItem.Value, ok = constant.Int64Val(kst.Val())
					resItem.IsInteger = true
				} else if basic.Info()&types.IsString > 0 {
					resItem.ValueString = strings.Trim(kst.Val().ExactString(), "\"")
				} else {
					//panic("ident must be interger or string" + name.Name)
					break
				}

				resItem.Comment = comment
				res = append(res, &resItem)
			}
		}
	}
	return pkgName, res, nil
}
