package main

import (
	"log"

	"github.com/Proudprogamer/goAuth/http/handlers"
	"github.com/Proudprogamer/goAuth/http/routes"
	"github.com/Proudprogamer/goAuth/prisma/db"
	"github.com/gin-gonic/gin"
)


func main(){

	client := db.NewClient()

	if err:= client.Prisma.Connect(); err!=nil {
		log.Fatal("failed to connect to the database", err.Error())
	}

	defer func() {
		if err:= client.Prisma.Disconnect(); err!=nil {
			panic(err)
		}
	}()

	handler := handlers.NewHandler(client)
	router:= gin.Default()

	routes.SetUpRoutes(router, handler)
	router.Run("localhost:8000")
}