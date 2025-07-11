package entities

type ICoordenada interface {
	MinValue() float64
	MaxValue() float64
	Latitude() float64
	Longitude() float64
}

type Coordenada struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (c Coordenada) Latitude() float64 {
	return c.Lat
}

func (c Coordenada) Longitude() float64 {
	return c.Lon
}

func (c Coordenada) MinValue() float64 {
	if c.Lat < c.Lon {
		return c.Lat
	}
	return c.Lon
}

func (c Coordenada) MaxValue() float64 {
	if c.Lat > c.Lon {
		return c.Lat
	}
	return c.Lon
}

func NewCoordenada(lat, lon float64) ICoordenada {
	return &Coordenada{
		Lat: lat,
		Lon: lon,
	}
}
