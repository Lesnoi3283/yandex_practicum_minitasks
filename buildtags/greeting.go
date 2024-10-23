//go:build 1.16

package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

//go:embed greeting.txt
var greeting string

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting)
}
