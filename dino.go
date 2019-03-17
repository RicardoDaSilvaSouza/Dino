package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/RicardoDaSilvaSouza/Dino/dynowebportal"
)

type config struct {
	WebServer string `json:"webserver,required"`
}

func main() {
	file, fErr := os.Open("config.json")
	if fErr != nil {
		log.Fatalln(fErr)
	}

	config := new(config)
	if jErr := json.NewDecoder(file).Decode(config); jErr != nil {
		log.Fatalln(jErr)
	}

	if err := dynowebportal.RunWebPortal(config.WebServer); err != nil {
		panic(err.Error())
	}
	log.Println("Shutting down Dino Web Portal...")
}
