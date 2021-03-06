// Package classification PinSaltHash API.
//
// Generator of pin code, salt and SHA-1 hash
//
//     Schemes: http
//     Host: localhost
//     BasePath: /api/v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Artem Shestakov<artem.s.shestakov@gmail.com.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - application/json
// swagger:meta
package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/artem-shestakov/pin-generator/internal/app/apiserver"
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
