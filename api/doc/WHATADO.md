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

func NewMinMaxCoordenada(listaValores float64) *Coordenada[] {
    
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
    
    variaveis := [latsum, latdif, lonsum, londif]
}

type Dado struct {
	ID    string `json:"id"`
	Placa string `json:"placa"`
	/*
		IDVeiculo    string `json:"idveiculo"`
		Principal    string `json:"principal"`
		WebGrupoID   string `json:"web_grupo_id"`
		Descricao    string `json:"descricao"`
		DataGPS      string `json:"datagps"`
		Localizacao  string `json:"localizacao"`
		StatusIgn    string `json:"statusign"`
		StatusGPS    string `json:"statusgps"`
		Velocidade   string `json:"velocidade"`
		Direcao      string `json:"direcao"`
	*/
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	/*
		Satelites    string `json:"satelites"`
		NivelCel     string `json:"nivelcel"`
		MovingStatus string `json:"moving_status"`
		GPRSConn     string `json:"gprs_connection"`
		GSMJamming   string `json:"gsm_jamming"`
		Ancora       string `json:"ancora"`
		WebGrupo     string `json:"web_grupo"`
		OdometerCan  string `json:"odometerCan"`
		Blocked      string `json:"blocked"`
	*/
}