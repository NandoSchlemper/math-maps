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
	cos_angle := (c.km() * math.Pi) / 180.0
	delta_latitude := (c.km() / 2) / 111
	delta_longitude := (c.km() / 2) / (111 * math.Cos(cos_angle))

	latsum := c.cen().Latitude() + delta_latitude
	latdif := c.cen().Latitude() - delta_latitude
	lonsum := c.cen().Longitude() + delta_longitude
	londif := c.cen().Longitude() - delta_longitude

	lat := NewCoordenada(latsum, latdif)
	lon := NewCoordenada(lonsum, londif)

	return lat, lon
}
