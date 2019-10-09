package main

import (
	"fmt"
	"github.com/gin-rest-prisma-boilerplate/config"
	"github.com/gin-rest-prisma-boilerplate/router"
	"log"
)

func main() {
	// set this when moving between production and development configs
	yaml := config.Configuration("development")
	port := yaml.GetString("server.port")

	routing := router.Router()
	fmt.Printf("🚀 Preparing to listen on port %s\n", port)
	err := routing.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
