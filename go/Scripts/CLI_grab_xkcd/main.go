package main

import (
	"flag"
	"fmt"
	"github.com/Danr17/GO_scripts/tree/master/Scripts/CLI_grab_xkcd/client"
	"log"
	"time"
)

func main() {
	comicNo := flag.Int(
		"n", int(client.LatestComic), "Comic number to fetch (default latest)",
	)
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	saveImage := flag.Bool(
		"s", false, "Save image to current directory",
	)
	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)
	flag.Parse()

	xkcdClient := client.NewXKCDClient()
	xkcdClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	comic, err := xkcdClient.Fetch(client.ComicNumber(*comicNo), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}
