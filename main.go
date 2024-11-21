package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicação := app.Gerar()
	erro := aplicação.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
