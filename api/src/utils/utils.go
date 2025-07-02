package utils

import (
	"fmt"
	"math"
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

func NewMathCaixaSquare(c e.ISquare) (e.ICoordenada, e.ICoordenada) {
	center := c.Cen()
	halfKm := c.Km() / 2

	deltaLat := halfKm / 111

	cosAngle := (c.Km() * math.Pi) / 180
	deltaLon := halfKm / (111 * math.Cos(cosAngle))

	lat := e.NewCoordenada(center.Latitude()+deltaLat, center.Latitude()-deltaLat)
	lon := e.NewCoordenada(center.Longitude()+deltaLon, center.Longitude()-deltaLon)

	return lat, lon
}

func RegisterCoordenadas(l *e.Locations) {
	delta_lat, delta_lon := NewClient(1.0, -22.49096, -47.58047)
	salto_lat, salto_lon := NewClient(1.5, -23.74781, -47.58103)

	l.Localizacoes = append(l.Localizacoes, *e.NewQuadradoLocations("Delta", delta_lat, delta_lon))
	l.Localizacoes = append(l.Localizacoes, *e.NewQuadradoLocations("Salto", salto_lat, salto_lon))
}

func NewClient(km float64, lat, lon float64) (e.ICoordenada, e.ICoordenada) {
	return NewMathCaixaSquare(e.NewSquare(km, e.NewCoordenada(lat, lon)))
}
