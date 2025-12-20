package app

func Run() {
	// Direccionar rutas
	router := NewRouter()

	// Iniciar el servidor
	Start(router)

}
