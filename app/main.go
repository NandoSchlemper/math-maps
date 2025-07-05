package main

import (
	"math-coordenadas/frontend/pages"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// src.App()
	// ticker := time.NewTicker(5 * time.Minute)
	// defer ticker.Stop()

	// for {
	// 	<-ticker.C
	// 	src.App()
	// }

	mux := http.NewServeMux()
	component := pages.Hello("Thor")
	mux.Handle("/", templ.Handler(component))
	http.ListenAndServe(":3030", mux)
}
