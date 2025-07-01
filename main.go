package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"math-coordenadas/src/entities"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

/* Função responsável por inicializar a aplicação.*/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar a as variaveis de ambiente: ", err.Error())
	}

	/*
		cord := e.NewCoordenada(-22.49265586923846, -47.58156334720672)
		l := e.NewLocations()
		l.RegisterCoordenadas()
		for x := range l.Localizacoes {
			result := utils.IsAt(cord, l.Localizacoes[x].Lat, l.Localizacoes[x].Lon)
			if result {
				println("Está dentro de X cliente")
			} else {
				println("Não está dentro de X cliente")
			}
		}
	*/

	usuario := os.Getenv("WRSAT_USER")
	password := os.Getenv("WRSAT_PASSWORD")

	client := http.Client{
		Timeout: time.Second * 2,
	}
	url := os.Getenv("base_api_url")
	println(url)

	payload := map[string]string{
		"usuario":   usuario,
		"senha":     password,
		"ordem":     "ASC",
		"limit":     "100",
		"pagina":    "1",
		"descricao": "",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Erro ao jsonify o payload Go")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}

	rfm := entities.API{User: usuario, Password: password}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(rfm.User+":"+rfm.Password)))
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
	var resp entities.Response
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		panic(err)
	}

	log.Printf("Response:\n%+v", resp)

}
