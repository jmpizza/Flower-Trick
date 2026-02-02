package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmpizza/Flower-Trick/internal/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//Configuracion de CORS
	router.Use(cors.Default())
	//Pokemon
	router.GET("/pokemon", handlers.Pagination)
	router.GET("/ability/:id", handlers.Ability)
	router.GET("/characteristic/:id", handlers.Characteristic)
	router.GET("/egg-group/:id", handlers.EggGroup)
	router.GET("/gender/:id", handlers.Gender)
	router.GET("/growt-rate/:id", handlers.GrowthRate)
	router.GET("/pokeathol-stat/:id", handlers.PokeathlonStat)
	router.GET("/pokemon/:id", handlers.Pokemon)
	router.GET("/pokemon-color/:id", handlers.PokemonColor)
	router.GET("/pokemon-form/:id", handlers.PokemonForm)
	router.GET("/pokemon-habitat/:id", handlers.PokemonHabitat)
	router.GET("/pokemon-shape/:id", handlers.PokemonShape)
	router.GET("/pokemon-specie/:id", handlers.PokemonSpecies)
	router.GET("/stat/:id", handlers.Stat)
	router.GET("/type/:id", handlers.Type)

	return router
}
