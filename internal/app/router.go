package app

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Crear router
	router := gin.Default()

	return router
}
