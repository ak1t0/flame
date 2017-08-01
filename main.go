package main

import (
	"github.com/ak1t0/flame/crawler"
	"github.com/ak1t0/flame/format"
	"github.com/ak1t0/flame/reader"
	"log"
)

func main() {
	parsed := reader.ReadJson("flame.json")
	r := crawler.Scan(format.NewOnionLogs(parsed))

	log.Println(r)
}
