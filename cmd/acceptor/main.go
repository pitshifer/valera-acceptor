package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pitshifer/valera-acceptor/internal/app/apiserver"
	"github.com/pitshifer/valera-acceptor/internal/migrate"
)

var configPath string
var migrateAction string

func init() {
	flag.StringVar(&configPath, "c", "configs/acceptor.toml", "path to config file")
	flag.StringVar(&migrateAction, "migrate", "", "apply migrates up or down")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Apply migrations
	if len(migrateAction) > 0 {
		err := migrate.Do(config, migrateAction)
		if err != nil {
			log.Println(err)
			os.Exit(0)
		}
		log.Println("migrations applied successfully.")
		os.Exit(0)
	}

	if err = apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
