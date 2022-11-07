package main

import (
	"projet/restserveragent"
	"fmt"
	"log"
	_ "time"
)

func main() {

	const url = ":8080"

	servAgt := restserveragent.NewRestServerAgent(url)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	fmt.Scanln()
}
