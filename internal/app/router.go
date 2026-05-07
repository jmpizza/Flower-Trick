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

	//Pokemon Pagination
	router.GET("/pokemon", handlers.Pagination)
	router.GET("/ability", handlers.Pagination)
	router.GET("/type", handlers.Pagination)
	router.GET("/pokemon-habitat", handlers.Pagination)
	router.GET("/pokemon-color", handlers.Pagination)
	router.GET("/pokemon-shape", handlers.Pagination)

	//Moves Pagination
	router.GET("/move", handlers.Pagination)

	//Pokemon
	router.GET("/ability/:id", handlers.Ability)
	router.GET("/characteristic/:id", handlers.Characteristic)
	router.GET("/egg-group/:id", handlers.EggGroup)
	router.GET("/gender/:id", handlers.Gender)
	router.GET("/growt-rate/:id", handlers.GrowthRate)
	router.GET("/pokeathol-stat/:id", handlers.PokeathlonStat)
	router.GET("/pokemon/:id", handlers.Pokemon)
	router.GET("/pokemon-summary/:id", handlers.PokemonSummary)
	router.GET("/pokemon-color/:id", handlers.PokemonColor)
	router.GET("/pokemon-form/:id", handlers.PokemonForm)
	router.GET("/pokemon-habitat/:id", handlers.PokemonHabitat)
	router.GET("/pokemon-shape/:id", handlers.PokemonShape)
	router.GET("/pokemon-specie/:id", handlers.PokemonSpecies)
	router.GET("/stat/:id", handlers.Stat)
	router.GET("/type/:id", handlers.Type)

	//Moves
	router.GET("/move/:id", handlers.Move)

	return router
}
