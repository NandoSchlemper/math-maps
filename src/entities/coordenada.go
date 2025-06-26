package entities

type ICoordenada interface {
	MinValue() float64
	MaxValue() float64
	Latitude() float64
	Longitude() float64
}

type Coordenada struct {
	lat float64
	lon float64
}

func (c Coordenada) Latitude() float64 {
	return c.lat
}

func (c Coordenada) Longitude() float64 {
	return c.lon
}

func (c Coordenada) MinValue() float64 {
	if c.lat < c.lon {
		return c.lat
	}
	return c.lon
}

func (c Coordenada) MaxValue() float64 {
	if c.lat > c.lon {
		return c.lat
	}
	return c.lon
}

func NewCoordenada(lat, lon float64) ICoordenada {
	return &Coordenada{
		lat: lat,
		lon: lon,
	}
}
