package utils

import (
	"fmt"
	"math-coordenadas/src/entities"
)

func SquareVerify(alvo entities.ICoordenada, latmax, latmin, lonmax, lonmin float64) {
	if alvo.Latitude() >= latmin &&
		alvo.Latitude() <= latmax &&
		alvo.Longitude() >= lonmin &&
		alvo.Longitude() <= lonmax {
		fmt.Println("Está dentro")
	} else {
		fmt.Println("Está fora")
	}
}
