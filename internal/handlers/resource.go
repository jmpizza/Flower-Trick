package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmpizza/Flower-Trick/internal/models"
)

func Pagination(c *gin.Context) {
	var obj models.NamedResponse
	var result APIResponse
	err := doRequest("pokemon", &result, &obj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
	}
	c.JSON(http.StatusOK, result)
}
