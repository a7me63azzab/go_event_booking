package main

import (
	"azzab.com/event_booking/db"
	"azzab.com/event_booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
