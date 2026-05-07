package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmpizza/Flower-Trick/internal/models"
)

func Pagination(c *gin.Context) {
	var obj models.NamedResponse
	var result APIResponse
	path := c.Request.RequestURI
	err := doRequest(path, &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}
