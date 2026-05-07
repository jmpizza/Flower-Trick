package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serverError(c *gin.Context, result APIResponse){
	result.Code = 500
	result.Message = "Server Error"
	c.JSON(http.StatusInternalServerError, result)
}