package entities

type Cerco struct {
	Name string
	Lat  ICoordenada
	Lon  ICoordenada
}

type Cercos struct {
	Cercos []Cerco
}

func NewCerco(name string, lat, lon ICoordenada) *Cerco {
	return &Cerco{
		Name: name,
		Lat:  lat,
		Lon:  lon,
	}
}

func NewCercos() []Cerco {
	return []Cerco{}
}
