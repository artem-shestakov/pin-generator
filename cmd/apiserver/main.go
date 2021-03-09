package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/artem-shestakov/pin-generator/env"
	"github.com/artem-shestakov/pin-generator/internal/app/apiserver"
)

var (
	listenAddres string = env.GetEnv("LISTEN_ADDRESS", "0.0.0.0")
	listenPort   string = env.GetEnv("LISTEN_PORT", "8080")
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
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
