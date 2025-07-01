package entities

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Dado struct {
	ID     string  `json:"id"`
	Placa  string  `json:"placa"`
	Lat    float64 `json:"lat,string"`
	Lon    float64 `json:"lon,string"`
	Status string
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

type WrsatAPI struct {
	Usuario  string
	Password string
	Url      string
}

type IWrsatAPI interface {
	LoadVariables()
	FetchData() Response
}

// fetchData implements IWrsatAPI.
func (w *WrsatAPI) FetchData() Response {
	client := http.Client{Timeout: time.Second * 2}

	payload := map[string]string{
		"usuario":   w.Usuario,
		"senha":     w.Password,
		"ordem":     "ASC",
		"limit":     "100",
		"pagina":    "1",
		"descricao": "",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Erro ao Marshal do JSON")
	}

	req, err := http.NewRequest(http.MethodPost, w.Url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(w.Usuario+":"+w.Password)))
	req.Header.Set("Content-Type", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal("Erro ao realizar a requisição: \n", getErr)
	}

	defer res.Body.Close()

	if req.Body == nil {
		log.Fatal("Body = nil, wtf brother")
	}

	jsonBytes, _ := io.ReadAll(res.Body)
	res.Body.Read(jsonBytes)
	var resp Response
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		panic(err)
	}

	return resp
}

// loadVariables implements IWrsatAPI
func (w *WrsatAPI) LoadVariables() {
	w.Usuario = os.Getenv("WRSAT_USER")
	w.Password = os.Getenv("WRSAT_PASSWORD")
	w.Url = os.Getenv("WRSAT_API_URL")
}

func NewWrsatAPI() IWrsatAPI {
	return &WrsatAPI{}
}
