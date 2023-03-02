package main

import (
	"holidays-seeker/cmd/config"
	"holidays-seeker/internal/http/server"
	"holidays-seeker/internal/infraestucture/dependencies"

	"fmt"
)

// Holidays seeker api
//
// This is a simple api to filter for holidays by days and extra.
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api/v1
//	Version: 0.0.1
//
//	Consumes:
//
//	Produces:
//	- application/json
//	- application/xml
//
// swagger:meta
func main() {
	fmt.Println("Running...")

	conf, err := config.LoadConfig("./internal/shared/config")
	if err != nil {
		fmt.Println("Error loading config: ", err)
		panic(err.Error())
	}

	fmt.Println("Config loaded", conf)

	container := dependencies.NewContainer(conf)

	server.Run(container)
}
