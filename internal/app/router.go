package app

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Router creation
	router := gin.Default()

	return router
}
