package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error al inicial el servidor: %v", err)
	}

}
