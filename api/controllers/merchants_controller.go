package controllers

import (
	"net/http"
	"strconv"
	"github.com/wentyocta/majoo/api/models"
	"github.com/wentyocta/majoo/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetMerchant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["merchant_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	merchant := models.Merchant{}

	merchantReceived, err := merchant.FindMerchantByUserID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, merchantReceived)
}
