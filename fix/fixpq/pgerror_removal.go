package main

import (
	"github.com/tmc/fix"
	"go/ast"
)

func init() {
	fix.Register(pgerror_removalfix)
}

var pgerror_removalfix = fix.Fix{
	"pgerror_removal",
	"2013-10-08",
	pgerror_removal,
	`Convert PGError interface use to Error struct`,
}

func pgerror_removal(f *ast.File) bool {
	if !fix.Imports(f, "github.com/lib/pq") {
		return false
	}

	fixed := false
	fix.Walk(f, func(n interface{}) {
		cl, ok := n.(*ast.TypeAssertExpr)
		if !ok {
			return
		}
		se, ok := cl.Type.(*ast.SelectorExpr)
		if !ok {
			return
		}

		if se.X.(*ast.Ident).Name == "pq" && se.Sel.Name == "PGError" {
			se.X.(*ast.Ident).Name = "*pq"
			se.Sel.Name = "Error"
			fixed = true
		}
	})
	return fixed
}
