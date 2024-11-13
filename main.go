package main

import (
	"gin_api/config"
	"gin_api/models"
	"gin_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Usuario{})

	router := gin.Default()

	routes.ConfigurarRutasUsuario(router)

	router.Run(":8080")
}
