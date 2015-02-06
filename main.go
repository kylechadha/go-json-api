package main

import (
	"os"

	"github.com/codegangsta/negroni"
)

func main() {

	// CONFIG
	// --------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		// negroni.HandlerFunc(AppMiddleware),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	// SERVER
	// --------------
	n.Run(":" + port)

}
