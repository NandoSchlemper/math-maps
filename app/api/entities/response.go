package entities

type Dado struct {
	ID       string  `json:"id"`
	Placa    string  `json:"placa"`
	Lat      float64 `json:"lat,string"`
	Lon      float64 `json:"lon,string"`
	Location string
}

type Response struct {
	Erro      bool   `json:"erro"`
	Status    int    `json:"status"`
	Mensagem  string `json:"mensagem"`
	Ordem     string `json:"ordem"`
	Limit     string `json:"limit"`
	Pagina    string `json:"pagina"`
	QtdResult int    `json:"qtd_result"`
	Dados     []Dado `json:"dados"`
}
