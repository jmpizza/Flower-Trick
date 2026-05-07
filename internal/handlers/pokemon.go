package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmpizza/Flower-Trick/internal/models"
)

func Ability(c *gin.Context) {
	var obj models.Ability
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("ability/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Characteristic(c *gin.Context) {
	var obj models.Characteristic
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("characteristic/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func EggGroup(c *gin.Context) {
	var obj models.EggGroup
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("egg-group/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Gender(c *gin.Context) {
	var obj models.Gender
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("gender/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func GrowthRate(c *gin.Context) {
	var obj models.GrowthRates
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("growth-rate/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Nature(c *gin.Context) {
	var obj models.Nature
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("nature/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokeathlonStat(c *gin.Context) {
	var obj models.PokeathlonStat
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokeathol-stat/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Pokemon(c *gin.Context) {
	var obj models.Pokemon
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonSummary(c *gin.Context) {
	var obj models.PokemonSummary
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonColor(c *gin.Context) {
	var obj models.PokemonColor
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon-color/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonForm(c *gin.Context) {
	var obj models.PokemonForm
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon-form/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonHabitat(c *gin.Context) {
	var obj models.PokemonHabitat
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon-habitat/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonShape(c *gin.Context) {
	var obj models.PokemonShape
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon-shape/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func PokemonSpecies(c *gin.Context) {
	var obj models.PokemonSpecies
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("pokemon-species/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Stat(c *gin.Context) {
	var obj models.Stat
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("stat/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}

func Type(c *gin.Context) {
	var obj models.Type
	var result APIResponse
	id := c.Param("id")
	err := doRequest(fmt.Sprintf("type/%s", id), &result, &obj)
	if err != nil {
		serverError(c, result)
		return 
	}
	c.JSON(http.StatusOK, result)
}
