package errcheck

// This file I wrote myself. It was a mini-task from Yandex Practicum.

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var DeferErrCheckAnalyzer = &analysis.Analyzer{
	Name: "defererrcheck",
	Doc:  "check for unchecked errors from goroutines",
	Run:  runDeferErrCheck,
}

func runDeferErrCheck(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if goSt, ok := n.(*ast.DeferStmt); ok {
				if isReturnError(pass, goSt.Call) {
					pass.Reportf(goSt.Pos(), "unchecked error in defer call")
				}
			}
			return true
		})
	}
	return nil, nil
}
