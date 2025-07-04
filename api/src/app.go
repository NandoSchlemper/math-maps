package src

import (
	"log"
	"math-coordenadas/src/entities"
	"math-coordenadas/src/utils"
	"time"

	"github.com/joho/godotenv"
)

func App() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar a as variaveis de ambiente: ", err.Error())
	}

	clientApi := entities.NewWrsatAPI()
	client_locations := entities.NewLocations()

	clientApi.LoadVariables()
	response := clientApi.FetchData()

	utils.RegisterCoordenadas(client_locations)

	defer utils.Timer(time.Now(), "Fetching data")

	utils.ConcatenateCoordenadas(&response, client_locations)

	log.Printf("%v", response)
}
