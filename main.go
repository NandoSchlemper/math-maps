package main

import (
	e "math-coordenadas/src/entities"
	u "math-coordenadas/src/utils"
)

func main() {
	centro_cliente1 := e.NewCoordenada(-22.49096, -47.58047)
	cliente1 := e.NewMathBox(1.0, centro_cliente1)
	lat, lon := e.NewMathBoxSquare(cliente1)

	coordenada_alvo3 := e.NewCoordenada(-80.4367065, -67.5699109)
	coordenada_alvo2 := e.NewCoordenada(-22.4367065, -47.5699109)
	coordenada_alvo := e.NewCoordenada(-22.49096, -47.58047)
	u.IsAt(coordenada_alvo, lat, lon)  // true
	u.IsAt(coordenada_alvo2, lat, lon) // false
	u.IsAt(coordenada_alvo3, lat, lon) // false
}
