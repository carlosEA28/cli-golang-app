package app

import "github.com/urfave/cli"

//Its going to return the cli app ready to be executed
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação cli"
	app.Usage = "Busca ips e nomes de servidor"

	return app
}
