package main

import (
	"flag"
)

var (
	config = flag.String("config", "", "config file path")
)

func main() {
	flag.Parse()

	if *config == "" {
		flag.Usage()
		return
	}

}
