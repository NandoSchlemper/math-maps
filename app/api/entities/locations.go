package entities

type LocationType int64

const (
	Carregamento LocationType = iota
	Posto
	Descarga
	RFM
)

type Cerco struct {
	Name string
	Type LocationType
	Lat  ICoordenada
	Lon  ICoordenada
}

type Cercos struct {
	Cercos []Cerco
}

func NewCerco(name string, localizacaotype LocationType, lat, lon ICoordenada) *Cerco {
	return &Cerco{
		Name: name,
		Type: localizacaotype,
		Lat:  lat,
		Lon:  lon,
	}
}

func NewCercos() []Cerco {
	return []Cerco{}
}
