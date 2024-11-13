package routes

import (
	"gin_api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRutasUsuario(r *gin.Engine) {
	usuario := r.Group("/usuarios")
	{
		usuario.GET("/", controllers.ObtenerUsuarios)
		usuario.GET("/:id", controllers.ObtenerUsuarioPorID)
		usuario.POST("/", controllers.CrearUsuario)
		usuario.PUT("/:id", controllers.ActualizarUsuario)
		usuario.DELETE("/:id", controllers.EliminarUsuario)
	}
}
