package handlers

import (
	"fmt"

	"github.com/jmpizza/Flower-Trick/internal/models"
)

func Ability(id string) (result models.Ability, err error) {
	err = do(fmt.Sprintf("ability/%s", id), &result)
	return result, err
}

func Characteristic(id string) {}

func EggGroup(id string) {}

func Gender(id string) {}

func GrowthRate(id string) {}

func Nature(id string) {}

func PokeathlonStat(id string) {}

func Pokemon(id string) {}

func PokemonColor(id string) {}

func PokemonForm(id string) {}

func PokemonHabitad(id string) {}

func PokemonShape(id string) {}

func PokemonSpecies(id string) {}

func Stats(id string) {}

func Type(id string) {}
