package main

import (
	"bin-term-chat/engine"
	"flag"
	"log"
	"strings"
)

var f = flag.String("url", "", "url for connected to server")

func main() {
	flag.Parse()

	if *f == "" {
		log.Fatal("url required")
	}

	*f = strings.TrimRight(*f, "/")

	err := engine.NewEngine(*f).Run()
	if err != nil {
		log.Fatal(err.Error())
	}

}
