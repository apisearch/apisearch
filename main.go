package main

import (
	"github.com/apisearch/importer/commands/elastic"
	"github.com/apisearch/importer/commands/importer"
	"github.com/apisearch/importer/routers"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "apisearch importer"
	app.Usage = "download xml files and import them into elasticsearch"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start http server",
			Action: func(c *cli.Context) error {
				log.Println("Starting HTTP server...")
				router := routers.NewRouter()
				log.Fatal(http.ListenAndServe(":8080", router))
				return nil
			},
		},
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "download and import xml files",
			Action: func(c *cli.Context) error {
				log.Println("Starting XML importer...")
				importer.ImportXmlFiles()
				return nil
			},
		},
		{
			Name:    "createIndex",
			Aliases: []string{"c"},
			Usage:   "create index in elasticsearch",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "force"},
			},
			Action: func(c *cli.Context) error {
				log.Println("Creating index...")
				elastic.CreateIndex(c.Bool("force"))
				return nil
			},
		},
	}
	app.Run(os.Args)
}
