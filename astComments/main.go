package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src := `/* Тестовый пакет */
package main

// Double умножает значение на 2.
func Double(i int) int {
    return i*2
}

func main() {
   // умножаем в цикле
   for i := 1; i < 5; i++ {
      fmt.Println(Double(i))
   }
}`
	// допишите код
	// ...

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", src, parser.ParseComments)
	for _, comment := range f.Comments {
		for _, c := range comment.List {
			fmt.Printf("%v : %s\n", c.Pos(), c.Text)
		}
	}
}
