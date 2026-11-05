package main

import (
	"lukasrar/go-rest-api-with-db/db"
	"lukasrar/go-rest-api-with-db/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080 in dev mode

}
