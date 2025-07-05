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

func ConcatenateCoordenadas(e *e.Response, l *[]e.Cerco) {
	for x := 0; x < len(e.Dados); x++ {
		for y := 0; y < len(*l); y++ {
			if IsAt(e.Dados[x].Lat, e.Dados[x].Lon, (*l)[y].Lat, (*l)[y].Lon) {
				e.Dados[x].Status = (*l)[y].Name
			} else {
				e.Dados[x].Status = "None"
			}
		}
	}
}

func Timer(start time.Time, nome string) {
	elapsed := time.Since(start)
	fmt.Printf("%s executado em: %.3f/n", nome, elapsed.Seconds())
}

func CalcularCoordenadas(c e.ISquare) (e.ICoordenada, e.ICoordenada) {
	center := c.Cen()
	halfKm := c.Km() / 2

	deltaLat := halfKm / 111

	cosAngle := (c.Km() * math.Pi) / 180
	deltaLon := halfKm / (111 * math.Cos(cosAngle))

	lat := e.NewCoordenada(center.Latitude()+deltaLat, center.Latitude()-deltaLat)
	lon := e.NewCoordenada(center.Longitude()+deltaLon, center.Longitude()-deltaLon)

	return lat, lon
}

func RegisterCoordenadas(l *[]e.Cerco) *[]e.Cerco {
	delta_lat, delta_lon := NewClient(1.000, -22.49096, -47.58047)
	salto_lat, salto_lon := NewClient(5.0, -23.74781, -47.58103)
	zanforlin_lat, zanforlin_lon := NewClient(0.140, -23.74781, -47.58103)
	braganceiro_lat, braganceiro_lon := NewClient(5.0, -24.16354, -48.90452)

	*l = append(*l, *e.NewCerco("Delta", delta_lat, delta_lon))
	*l = append(*l, *e.NewCerco("Salto", salto_lat, salto_lon))
	*l = append(*l, *e.NewCerco("Zanforlin", zanforlin_lat, zanforlin_lon))
	*l = append(*l, *e.NewCerco("Braganceiro", braganceiro_lat, braganceiro_lon))
	return l
}

func NewClient(km float64, lat, lon float64) (e.ICoordenada, e.ICoordenada) {
	return CalcularCoordenadas(e.NewSquare(km, e.NewCoordenada(lat, lon)))
}
