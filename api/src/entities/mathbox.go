package entities

import "math"

type IMathBox interface {
	cen() ICoordenada
	km() float64
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

func NewMatematicaBox(side float64, center ICoordenada) IMathBox {
	return &MathBox{
		side:   side,
		center: center,
	}
}

func NewMathCaixaSquare(c IMathBox) (ICoordenada, ICoordenada) {
	center := c.cen()
	halfKm := c.km() / 2

	deltaLat := halfKm / 111

	cosAngle := (c.km() * math.Pi) / 180
	deltaLon := halfKm / (111 * math.Cos(cosAngle))

	lat := NewCoordenada(center.Latitude()+deltaLat, center.Latitude()-deltaLat)
	lon := NewCoordenada(center.Longitude()+deltaLon, center.Longitude()-deltaLon)

	return lat, lon
}
