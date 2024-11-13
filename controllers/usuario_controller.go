package controllers

import (
	"gin_api/config"
	"gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ObtenerUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	config.DB.Find(&usuarios)
	c.JSON(http.StatusOK, usuarios)
}

func ObtenerUsuarioPorID(c *gin.Context) {
	var usuario models.Usuario
	if err := config.DB.First(&usuario, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func CrearUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo crear el usuario"})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

func ActualizarUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := config.DB.First(&usuario, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func EliminarUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := config.DB.First(&usuario, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	config.DB.Delete(&usuario)
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
