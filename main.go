package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {

	aplicação := app.Gerar()
	erro := aplicação.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
