package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start(c *gin.Engine) {
	// Obtener puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	if err := c.Run(":" + port); err != nil {
		log.Fatalf("Error al inicial el servidor: %v", err)
	}

}
