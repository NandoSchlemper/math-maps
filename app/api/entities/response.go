package entities

import (
	"fmt"
)

type Operacao struct {
	ID   string
	Nome string
}

type Dado struct {
	ID         string   `json:"id"`
	Placa      string   `json:"placa"`
	Lat        float64  `json:"lat,string"`
	Lon        float64  `json:"lon,string"`
	Velocidade int      `json:"velocidade,string"`
	Operação   Operacao `json:"web_grupo"`
	Estado     string
	Location   string
	Motorista  string
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

func (r Dado) Coordenadas() string {
	return fmt.Sprintf("%.6f, %.6f", r.Lat, r.Lon)
}

func (r Dado) Velocidades() string {
	return fmt.Sprintf("%d km/h", r.Velocidade)
}

func (r Dado) LinkPosition() string {
	return fmt.Sprintf("https://www.google.com/maps/search/%f,+%f?sa=X&ved=1t:242&ictx=111", r.Lat, r.Lon)
}
