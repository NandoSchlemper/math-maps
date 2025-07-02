package entities

type SquareLocalizações struct {
	Name string
	Lat  ICoordenada
	Lon  ICoordenada
}

type Locations struct {
	Localizacoes []SquareLocalizações
}

func NewQuadradoLocations(name string, lat, lon ICoordenada) *SquareLocalizações {
	return &SquareLocalizações{
		Name: name,
		Lat:  lat,
		Lon:  lon,
	}
}

func NewLocations() *Locations {
	return &Locations{}
}
