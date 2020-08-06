package main

import (
	"log"
	"net/http"

	"github.com/codeformio/hello-api/api"
	"github.com/kelseyhightower/envconfig"
)

const version = "1.5.1"

func main() {
	var cfg struct {
		Addr    string `envconfig:"ADDR" default:":8080"`
		Message string `envconfig:"MESSAGE" default:"Hello!"`
	}
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)

	log.Println("Where are my pancakes?")

	log.Println("They are: here")


	log.Println("Hi from Nick 2!")

	log.Println("Hi from Dan")
	log.Println("RAWR")
	log.Println("Help! Mike sank the row boat! -Nick1")


	log.Fatal(http.ListenAndServe(cfg.Addr, &api.Handler{Message: cfg.Message, Version: version}))
}
