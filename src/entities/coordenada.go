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
	var result float64
	if c.lat > c.lon {
		result = c.lon
	} else if c.lat < c.lon {
		result = c.lat
	}

	return result
}

func (c Coordenada) MaxValue() float64 {
	var result float64
	if c.lat > c.lon {
		result = c.lat
	} else if c.lat < c.lon {
		result = c.lon
	}

	return result
}

func NewCoordenada(lat, lon float64) ICoordenada {
	return &Coordenada{
		lat: lat,
		lon: lon,
	}
}
