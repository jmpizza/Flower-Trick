package app

import (
	"github.com/jmpizza/Flower-Trick/internal/handlers"
)

func Run() {
	// Routing
	r := NewRouter()

	handlers.Handler()

	// Server start
	Start(r)

}
