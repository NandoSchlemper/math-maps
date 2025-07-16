package utils

import (
	"fmt"
	"math"
	"math-coordenadas/api/entities"
	"time"
)

func IsAt(alvo_lat, alvo_lon float64, lat, lon entities.ICoordenada) bool {
	return alvo_lat >= lat.MinValue() && alvo_lat <= lat.MaxValue() &&
		alvo_lon >= lon.MinValue() && alvo_lon <= lon.MaxValue()
}

// Adicionar passos lógicos à essa
func ConcatenateCoordenadas(entities *entities.Response, l *[]entities.Cerco) {
	for x := 0; x < len(entities.Dados); x++ {
		for y := 0; y < len(*l); y++ {
			if IsAt(entities.Dados[x].Lat, entities.Dados[x].Lon, (*l)[y].Lat, (*l)[y].Lon) {
				entities.Dados[x].Location = (*l)[y].Name
			} else {
				entities.Dados[x].Location = "None"
			}
		}
	}
}

func Timer(start time.Time, nome string) {
	elapsed := time.Since(start)
	fmt.Printf("%s executado em: %.3f/n", nome, elapsed.Seconds())
}

func CalcularCoordenadas(c entities.ISquare) (entities.ICoordenada, entities.ICoordenada) {
	center := c.Cen()
	halfKm := c.Km() / 2

	deltaLat := halfKm / 111

	cosAngle := (c.Km() * math.Pi) / 180
	deltaLon := halfKm / (111 * math.Cos(cosAngle))

	lat := entities.NewCoordenada(center.Latitude()+deltaLat, center.Latitude()-deltaLat)
	lon := entities.NewCoordenada(center.Longitude()+deltaLon, center.Longitude()-deltaLon)

	return lat, lon
}

// Fazer um formulario para criação e inclusão de mais posições, salvando no DB
func RegisterCoordenadas(l *[]entities.Cerco) {
	NewClient(1.0, -22.49096, -47.58047, "Delta", l)
	NewClient(5.0, -23.74781, -47.58103, "Salto", l)
	NewClient(0.2, -23.74781, -47.58103, "Zanforlin", l)
	NewClient(5.0, -24.16354, -48.90452, "Braganceiro", l)
}

func NewClient(km, lat, lon float64, name string, L *[]entities.Cerco) {
	latitude, longitude := CalcularCoordenadas(entities.NewSquare(km, entities.NewCoordenada(lat, lon)))
	*L = append(*L, *entities.NewCerco(name, latitude, longitude))
}

// NewClient vai receber:
// Nome, Descrição, Tipo, KM, Lat, Lon e []Cercos
