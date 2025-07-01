package entities

import "net/http"

// api v0.1
type Dado struct {
	ID    string `json:"id"`
	Placa string `json:"placa"`
	Lat   string `json:"lat"`
	Lon   string `json:"lon"`
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

type IWrsatAPI interface {
	fetchAPI(string, string, http.Client)
	getAddr() string
	getPlate() string
}
type API struct {
	User     string
	Password string
}

type WrsatAPI struct {
	addr string
}

// fetchAPI implements IWrsatAPI.
func (w *WrsatAPI) fetchAPI(user, password string, client http.Client) {
	panic("unimplemented")
}

// getAddr implements IWrsatAPI.
func (w *WrsatAPI) getAddr() string {
	return w.addr
}

// getPlate implements IWrsatAPI.
func (w *WrsatAPI) getPlate() string {
	panic("unimplemented")
}

func NewWrsastAPI(addr string) IWrsatAPI {
	return &WrsatAPI{
		addr: addr,
	}
}
