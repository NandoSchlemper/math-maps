package main

import (
	e "math-coordenadas/src/entities"
	"math-coordenadas/src/utils"
)

func main() {
	cord := e.NewCoordenada(-22.49265586923846, -47.58156334720672)
	l := e.NewLocations()
	l.RegisterCoordenadas()
	for x := range l.Localizacoes {
		result := utils.IsAt(cord, l.Localizacoes[x].Lat, l.Localizacoes[x].Lon)
		if result {
			println("Está dentro de X cliente")
		} else {
			println("Não está dentro de X cliente")
		}
	}
}
