package controllers

import (
	"net/http"
	"strconv"
	"github.com/wentyocta/majoo/api/models"
	"github.com/wentyocta/majoo/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetOutlet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["outlet_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	outlet := models.Outlet{}

	outletReceived, err := outlet.FindOutletByMerchantID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, outletReceived)
}
