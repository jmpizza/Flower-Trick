package app

import (
	"fmt"

	"github.com/jmpizza/Flower-Trick/internal/handlers"
)

func Run() {
	// Routing
	r := NewRouter()

	fmt.Println(handlers.Ability("-1"))

	// Server start
	Start(r)

}
