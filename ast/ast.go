package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func GetAST(code string) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, "", code, 0)
	if err != nil {
		panic(err)
	}
	_ = ast.Print(fileSet, f)
}
