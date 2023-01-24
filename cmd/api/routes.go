package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	hWallet "github.com/fajarabdillahfn/go-wallet/internal/delivery/http"
)

func routes(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte("Welcome to Wallet API"))
		if err != nil {
			return
		}
	})

	router.POST("/wallets", hWallet.DepositMoney)
	router.GET("/wallets/:wallet_id", hWallet.GetWalletDetails)
}
