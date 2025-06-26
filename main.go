package main

import (
	"fmt"
	"math"
)

type Coordenada struct {
	lat float64
	lon float64
}

type ICoordenada interface {
	minValue() float64
	maxValue() float64
	latitude() float64
	longitude() float64
}

func (c Coordenada) latitude() float64 {
	return c.lat
}

func (c Coordenada) longitude() float64 {
	return c.lon
}

func (c Coordenada) minValue() float64 {
	var result float64
	if c.lat > c.lon {
		result = c.lon
	} else if c.lat < c.lon {
		result = c.lat
	}

	return result
}

func (c Coordenada) maxValue() float64 {
	var result float64
	if c.lat > c.lon {
		result = c.lat
	} else if c.lat < c.lon {
		result = c.lon
	}

	return result
}

type IMathBox interface {
	cen() ICoordenada
	km() float64
}

func SquareVerify(alvo ICoordenada, latmax, latmin, lonmax, lonmin float64) {
	if alvo.latitude() >= latmin && alvo.latitude() <= latmax && alvo.longitude() >= lonmin && alvo.longitude() <= lonmax {
		fmt.Println("Está dentro")
	} else {
		fmt.Println("Está fora")
	}
}

type MathBox struct {
	side   float64
	center ICoordenada
}

func (m MathBox) cen() ICoordenada {
	return m.center
}

func (m MathBox) km() float64 {
	return m.side
}

func NewCoordenada(lat, lon float64) ICoordenada {
	return &Coordenada{
		lat: lat,
		lon: lon,
	}
}

func NewMathBox(side float64, center ICoordenada) IMathBox {
	return &MathBox{
		side:   side,
		center: center,
	}
}

func NewMathBoxSquare(c IMathBox) (ICoordenada, ICoordenada) {
	cos_angle := (c.km() * math.Pi) / 180.0
	delta_latitude := (c.km() / 2) / 111
	delta_longitude := (c.km() / 2) / (111 * math.Cos(cos_angle))

	latsum := c.cen().latitude() + delta_latitude
	latdif := c.cen().latitude() - delta_latitude
	lonsum := c.cen().longitude() + delta_longitude
	londif := c.cen().longitude() - delta_longitude

	lat := NewCoordenada(latsum, latdif)
	lon := NewCoordenada(lonsum, londif)

	return lat, lon
}

func main() {
	centro_cliente1 := NewCoordenada(-22.49096, -47.58047)
	cliente1 := NewMathBox(1.0, centro_cliente1)
	lat, lon := NewMathBoxSquare(cliente1)

	min_lat := lat.minValue()
	max_lat := lat.maxValue()
	min_lon := lon.minValue()
	max_lon := lon.maxValue()

	coordenada_alvo3 := NewCoordenada(-80.4367065, -67.5699109)
	coordenada_alvo2 := NewCoordenada(-22.4367065, -47.5699109)
	coordenada_alvo := NewCoordenada(-22.49096, -47.58047)
	SquareVerify(coordenada_alvo, max_lat, min_lat, max_lon, min_lon)
	SquareVerify(coordenada_alvo2, max_lat, min_lat, max_lon, min_lon)
	SquareVerify(coordenada_alvo3, max_lat, min_lat, max_lon, min_lon)
}
