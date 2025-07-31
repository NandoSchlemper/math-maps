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

// if

func ConcatenateCoordenadas(entities *entities.Response, l *[]entities.Cerco) {
	for x := 0; x < len(entities.Dados); x++ {
		found := false
		for y := 0; y < len(*l); y++ {
			if IsAt(entities.Dados[x].Lat, entities.Dados[x].Lon, (*l)[y].Lat, (*l)[y].Lon) {
				entities.Dados[x].Location = (*l)[y].Name
				found = true
				break
			}
		}
		if !found {
			entities.Dados[x].Location = "None"
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
	NewClient(1.0, -22.49096, -47.58047, "Delta Antiga", entities.Descarga, l)
	NewClient(1.1, -22.49258, -47.5676, "Delta Nova", entities.Descarga, l)
	NewClient(5.0, -23.74781, -47.58103, "Salto", entities.Descarga, l)
	NewClient(0.2, -23.74781, -47.58103, "Zanforlin", entities.Posto, l)
	NewClient(2.0, -24.1643, -48.9076, "Braganceiro", entities.Carregamento, l)
	NewClient(1.0, -22.44113, -47.56861, "Pátio", entities.RFM, l)
}

func NewClient(km, lat, lon float64, name string, tipagem entities.LocationType, L *[]entities.Cerco) {
	latitude, longitude := CalcularCoordenadas(entities.NewSquare(km, entities.NewCoordenada(lat, lon)))
	*L = append(*L, *entities.NewCerco(name, tipagem, latitude, longitude))
}
