package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	mux := http.NewServeMux()

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	log.Fatal(s.Shutdown(ctx))
}
