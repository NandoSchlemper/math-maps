package src

import (
	"log"
	"math-coordenadas/api/entities"
	"math-coordenadas/api/utils"
	"time"

	"github.com/joho/godotenv"
)

func App() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar a as variaveis de ambiente: ", err.Error())
	}

	clientApi := entities.NewWrsatAPI()
	client_locations := entities.NewCercos()

	utils.RegisterCoordenadas(&client_locations)
	clientApi.LoadVariables()
	response := clientApi.FetchData()

	defer utils.Timer(time.Now(), "Fetching data")

	utils.ConcatenateCoordenadas(&response, &client_locations)

	log.Printf("%v", response)
}
