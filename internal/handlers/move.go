package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmpizza/Flower-Trick/internal/models"
)

func Move(c *gin.Context) {
	var obj models.Move
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("move/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}