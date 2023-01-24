package http

import (
	"io/ioutil"
	"log"
	"net/http"

	cWrapper "github.com/fajarabdillahfn/go-wallet/common/wrapper"
	bKafka "github.com/fajarabdillahfn/go-wallet/internal/broker/kafka"

	"github.com/fajarabdillahfn/go-wallet/internal/model"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/julienschmidt/httprouter"
)

func DepositMoney(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var inputDeposit model.Deposit

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		cWrapper.ErrorJSON(w, err, "failed to read data", http.StatusInternalServerError)
		return
	}

	err = protojson.Unmarshal(data, &inputDeposit)
	if err != nil {
		cWrapper.ErrorJSON(w, err, "", http.StatusBadRequest)
		return
	}

	err = bKafka.Emit(ctx, &inputDeposit)
	if err != nil {
		log.Println(err)
		cWrapper.ErrorJSON(w, err, "failed to save data", http.StatusInternalServerError)
		return
	}

	cWrapper.WriteJSON(w, http.StatusOK, "deposit success", "")
}
