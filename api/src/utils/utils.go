package utils

import (
	"fmt"
	e "math-coordenadas/src/entities"
	"time"
)

func IsAt(alvo_lat, alvo_lon float64, lat, lon e.ICoordenada) bool {
	return alvo_lat >= lat.MinValue() && alvo_lat <= lat.MaxValue() &&
		alvo_lon >= lon.MinValue() && alvo_lon <= lon.MaxValue()
}

func ConcatenateCoordenadas(e *e.Response, l *e.Locations) {
	for x := 0; x < len(e.Dados); x++ {
		for y := 0; y < len(l.Localizacoes); y++ {
			if IsAt(e.Dados[x].Lat, e.Dados[x].Lon, l.Localizacoes[y].Lat, l.Localizacoes[y].Lon) {
				e.Dados[x].Status = l.Localizacoes[y].Name
			}
			e.Dados[x].Status = "None"
		}
	}
}

func Timer(start time.Time, nome string) {
	elapsed := time.Since(start)
	fmt.Printf("%s executado em: %.3f/n", nome, elapsed.Seconds())
}
