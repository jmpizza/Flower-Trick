package app

// import (
// 	//"github.com/jmpizza/Flower-Trick/internal/handlers"
// )

func Run() {
	// Routing
	r := NewRouter()

	//fmt.Println(handlers.Ability("1"))
	// //data, err := handlers.Pokemon("meowscarada")
	//data, err := handlers.Characteristic()
	// fmt.Printf("%#v\n\n\n\n", data)
	// fmt.Println(data)
	// fmt.Println(err)
	// Server start
	Start(r)

}
