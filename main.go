package main

import (
	"github.com/ak1t0/flame/crawler"
	"github.com/ak1t0/flame/format"
	"github.com/ak1t0/flame/reader"
	"github.com/urfave/cli"
	"log"
	"os"
)

var Version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "flame"
	app.Usage = "crawl onion services"
	app.Version = Version
	app.Author = "ak1t0"
	app.Email = "aktoo3097@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}

var Commands = []cli.Command{
	commandScan,
}

var commandScan = cli.Command{
	Name:    "scan",
	Usage:   "Scan onion services",
	Aliases: []string{"s"},
	Action:  doScan,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Usage: "Select log file",
		},
		cli.StringFlag{
			Name:  "o",
			Usage: "Select output file",
		},
	},
}

func doScan(c *cli.Context) error {
	var target string
	if c.String("f") != "" {
		target = c.String("f")
	} else {
		target = "log.json"
	}

	parsed, err := reader.ReadJson(target)
	if err != nil {
		log.Fatal(err)
	}

	r := crawler.Crawl(format.NewOnionLogs(parsed))

	if c.String("o") != "" {
		f, err := os.OpenFile(c.String("o"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(r)
			log.Fatal(err)
		}
		log.SetOutput(f)
		log.Println(r)
	} else {
		log.Println(r)
	}

	return nil
}
