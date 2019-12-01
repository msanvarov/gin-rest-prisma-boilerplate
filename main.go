package main

import (
	"fmt"
	"log"

	"github.com/msanvarov/gin-rest-prisma-boilerplate/config"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/router"
)

// @title API
// @version 1.0
// @description Simple API

// @contact.name API Support
// @contact.url https://msanvarov.github.io/personal-portfolio/
// @contact.email msalanvarov@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @basepath /

// Main method to act as a starting point for the web server.
func main() {
	config.Configure("config")
	yaml := config.GetConfiguration()
	port := yaml.GetString("server.port")
	routing := router.Router(yaml)
	fmt.Printf("🚀 Preparing to listen on port %s\n", port)
	err := routing.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
