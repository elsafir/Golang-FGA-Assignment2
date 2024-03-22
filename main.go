package main

import (
	"fmt"
	"log"

	"Golang-FGA-Assignment2/config"
	routes "Golang-FGA-Assignment2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	viper := config.NewViper()
	db := config.NewDatabase(viper)

	app := gin.Default()
	routes.NewRouter(db, app)
	port := viper.GetString("web.port")

	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
