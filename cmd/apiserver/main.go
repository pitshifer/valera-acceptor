package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/pitshifer/valera-acceptor/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
