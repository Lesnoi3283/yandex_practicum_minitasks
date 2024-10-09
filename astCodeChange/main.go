package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
    
func main() {
     ids := 77
     id := ids + 1
     fmt.Println("id равно:", id/2 )
}`

	// допишите код
	// ...

	fset := token.NewFileSet()
	node, _ := parser.ParseFile(fset, "src.go", src, parser.AllErrors)

	ast.Inspect(node, func(n ast.Node) bool {
		if v, ok := n.(*ast.Ident); ok {
			if v.Name == "id" {
				v.Name = "ident"
			}
		}
		return true
	})

	printer.Fprint(os.Stdout, fset, node)
}
