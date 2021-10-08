package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	paramP := r.FormValue("p")
	if paramP == "Gopher" {
		fmt.Fprint(w, "Hello, Gopher")
	} else {
		fmt.Fprint(w, "Hello, HTTPサーバ")
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
