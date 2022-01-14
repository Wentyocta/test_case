package controllers

import (
    "github.com/wentyocta/majoo/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Merchants routes
	s.Router.HandleFunc("/merchants/{id}", middlewares.SetMiddlewareJSON(s.GetMerchant)).Methods("GET")

	//Outlets routes
	s.Router.HandleFunc("/outlets/{merchant_id}", middlewares.SetMiddlewareJSON(s.GetOutlet)).Methods("GET")

	//Transactions routes
    s.Router.HandleFunc("/transactions/{merchant_id}/{outlet_id}", middlewares.SetMiddlewareJSON(s.GetTransaction)).Methods("GET")
}