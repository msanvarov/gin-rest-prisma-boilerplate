package main

import (
	"fmt"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/config"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/router"
	"log"
)

func main() {
	yaml := config.Configuration("env-config")
	port := yaml.GetString("server.port")

	routing := router.Router(yaml)
	fmt.Printf("🚀 Preparing to listen on port %s\n", port)
	err := routing.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
