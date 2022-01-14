package controllers

import (
	"net/http"
	"strconv"
	"github.com/wentyocta/majoo/api/responses"
	"github.com/gorilla/mux"
	"github.com/wentyocta/majoo/api/models"
)

func (server *Server) GetTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	merchant_id, err := strconv.ParseUint(vars["merchant_id"], 10, 32)
	outlet_id, err := strconv.ParseUint(vars["outlet_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	transaction := models.Transaction{}

	transactionReceived, err := transaction.FindTransactionByOutletID(server.DB, merchant_id, outlet_id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactionReceived)
}
