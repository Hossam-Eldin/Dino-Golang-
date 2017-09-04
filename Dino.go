package main

import (
	"os"
	"log"
	"encoding/json"

	"github.com/Hossam-Eldin/Dino/dinowebportal"
	"github.com/Hossam-Eldin/Dino/databaselayer"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("starting the server at ", config.ServerAddress)
	err = dinowebportal.RunWebPortal(databaselayer.MONGODB, config.ServerAddress, "mongodb://127.0.0.1", "dinowebportal/dinoTemplate")
	if err != nil {
		log.Panic(err)
	}
}
