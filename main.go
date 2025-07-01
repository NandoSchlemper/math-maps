package main

import (
	"math-coordenadas/src"
	"time"
)

/* Função responsável por inicializar a aplicação.*/
func main() {
	src.App()
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		src.App()
	}
}
