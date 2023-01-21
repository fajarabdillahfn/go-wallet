package api

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func Start() {
	port := ":" + os.Getenv("APP_PORT")

	initialize()

	router := httprouter.New()
	routes(router)

	log.Println("start application on port " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic("failed to start application")
	}
}
