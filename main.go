package main

import (
	"fmt"
	"log"

	"github.com/msanvarov/gin-rest-prisma-boilerplate/config"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/router"
)

// Main method to act as a starting point for the web server.
func main() {
	yaml := config.Configuration("config")
	port := yaml.GetString("server.port")

	routing := router.Router(yaml)
	fmt.Printf("🚀 Preparing to listen on port %s\n", port)
	err := routing.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
