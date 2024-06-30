package main

import (
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
    models.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run()
}
