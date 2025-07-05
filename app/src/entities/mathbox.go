package entities

type ISquare interface {
	Cen() ICoordenada
	Km() float64
}

type Square struct {
	side   float64
	center ICoordenada
}

func (m Square) Cen() ICoordenada {
	return m.center
}

func (m Square) Km() float64 {
	return m.side
}

func NewSquare(side float64, center ICoordenada) ISquare {
	return &Square{
		side:   side,
		center: center,
	}
}
