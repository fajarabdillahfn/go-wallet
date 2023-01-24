package http

import (
	"net/http"
	"strconv"

	cWrapper "github.com/fajarabdillahfn/go-wallet/common/wrapper"
	"github.com/fajarabdillahfn/go-wallet/internal/model"
	"github.com/julienschmidt/httprouter"
)

func GetWalletDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	wallet_id, err := strconv.Atoi(ps.ByName("wallet_id"))
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	threshold, ok := model.ThresholdList[int32(wallet_id)]
	if !ok {
		cWrapper.ErrorJSON(w, err, "wallet not found", http.StatusNotFound)
		return
	}

	balance, ok := model.WalletsBalance[int32(wallet_id)]
	if !ok {
		cWrapper.ErrorJSON(w, err, "wallet not found", http.StatusNotFound)
		return
	}

	walletDetail := model.WalletDetails{
		WalletId:       wallet_id,
		Balance:        balance,
		AboveThreshold: threshold.AboveThreshold,
	}

	cWrapper.WriteJSON(w, http.StatusOK, walletDetail, "")
}
