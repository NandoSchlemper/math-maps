package main

import (
	"fmt"
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
	initial_page := pages.Initial()
	home_page := pages.Home()
	mux.Handle("/", templ.Handler(initial_page))
	mux.Handle("/home", templ.Handler(home_page))
	fmt.Println("Iniciando Template.")
	http.ListenAndServe(":3030", mux)
}
