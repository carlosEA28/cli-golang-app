package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Its going to return the cli app ready to be executed
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação cli"
	app.Usage = "Busca ips e nomes de servidor"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: BuscarIps,
		},
		{
			Name:  "servidores",
			Usage: "Buscar o nome do servidor na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: BuscarServidores,
		},
	}

	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
func BuscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servers := range servidores {
		fmt.Println(servers.Host)

	}

}
