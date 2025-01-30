package main

import (
	"github.com/Zarapinga/HealthTrack/configs"
	server "github.com/Zarapinga/HealthTrack/infra/webserver"
)

func main() {
	config, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}
	//err := databaseinit
	server.Server()
}
