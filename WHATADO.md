### 1. Variaveis Constantes de Posições de Clientes
Para utilizar isso, seria interessante criar constantes com as Boxes de cada cliente, seria algo assim:

type Coordenada struct {
    lat: float64, (mlr usar double ao invés de float devido a precisão das coordenadas)
    lon: float64,
}


type IMathBox interface {
    inTheBox(Coordenada) -> (err, bool)
}

type MathBox struct {
    lado: int -> em KM
    centro: Coordenada
}

func NewCoordenada(lat, lon float64) *Coordenada {
    return &Coordenada{
        lat: lat,
        lon: lon
    }
}

func NewMinMaxCoordenada(latsum, ladif, lonsum, londif float64) *Coordenada {
    if 
}

func NewMathBox(lado int, centro Coordenada) IMathBox {
    return &MathBox{
        lado: lado
        centro: centro
    }
}

func CalcularMinimoMaximo(c: MathBox) -> (err, Coordenada) {
    delta_latitude := (c.lado/2) / 111
    delta_longitude := (c.lado/2) / (111 * math.cos(c.lado * math.Pi / 180.0))

    latsum := c.centro.lat + delta_latitude
    latdif := c.centro.lat - delta_latitude
    lonsum := c.centro.lon + delta_longitude
    londif := c.centro.lon - delta_longitude
    
}
