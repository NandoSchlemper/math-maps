package main

import "math"

type Coordenada struct {
	lat float64
	lon float64
}

type ICoordenada interface {
	minValue() float64
	maxValue() float64
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
	inTheBox(Coordenada) (bool, error)
}

type MathBox struct {
	side   float64
	center Coordenada
}

// inTheBox implements IMathBox.
func (m *MathBox) inTheBox(Coordenada) (bool, error) {
	panic("unimplemented")
}

func NewCoordenada(lat, lon float64) ICoordenada {
	return &Coordenada{
		lat: lat,
		lon: lon,
	}
}

func NewMathBox(side float64, center Coordenada) IMathBox {
	return &MathBox{
		side:   side,
		center: center,
	}
}

func CalcularMinimoMaximo(c MathBox) ([]ICoordenada, error) {
	cos_angle := (c.side * math.Pi) / 180.0
	delta_latitude := (c.side / 2) / 111
	delta_longitude := (c.side / 2) / (111 * math.Cos(cos_angle))

	latsum := c.center.lat + delta_latitude
	latdif := c.center.lat - delta_latitude
	lonsum := c.center.lon + delta_longitude
	londif := c.center.lon - delta_longitude

	lat := NewCoordenada(latsum, latdif)
	lon := NewCoordenada(lonsum, londif)

	return []ICoordenada{lat, lon}, nil
}
