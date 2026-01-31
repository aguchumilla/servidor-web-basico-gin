package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola mundo!!") //200
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s!", nombre)

	})

	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario

		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error al decodificar el json"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre y email no pueden ser vacios"})
			return
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"Mensaje": "Usuario registrado ", "datos": usuarios})

	})
}
