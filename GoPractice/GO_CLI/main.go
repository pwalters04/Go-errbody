package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)
var app = cli.NewApp()
var pizza = []string{"You Ordered: "}



func info(){
	app.Name = "Shells Pizza Company"
	app.Usage="Order Some Pizza Please via CLI"
	app.Version="1.0"
}



func main() {

	app := &cli.App{
		Name:" Shells Pizza Company",
		Usage: "Pizza CLI Application",
		Commands: []*cli.Command{
			{
				Name:    "peppers",
				Aliases: []string{"p"},
				Usage:   "Add peppers to your pizza",
				Action:  func(c *cli.Context) error  {
					pe := "peppers"
					peppers := append(pizza, pe)
					m := strings.Join(peppers, " ")
					fmt.Println(m)
					return nil
				},
			},
			{
				Name:    "pineapple",
				Aliases: []string{"pa"},
				Usage:   "Add pineapple to your pizza",
				Action: func(c *cli.Context) error {
					pa := "pineapple"
					pineapple := append(pizza, pa)
					m := strings.Join(pineapple, " ")
					fmt.Println(m)
					return nil
				},
			},
		},
	}


	err :=app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

}
