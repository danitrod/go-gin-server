package main

import (
	"github.com/danitrod/go-gin-server/database"
	"github.com/danitrod/go-gin-server/routes"
)

func main() {
	database.ConnectToDB()
	routes.HandleRequests()
}
