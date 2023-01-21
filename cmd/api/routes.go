package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte("Welcome to Wallet API"))
		if err != nil {
			return
		}
	})
}
