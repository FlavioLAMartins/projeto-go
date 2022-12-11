package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func RegisterAPI(r *mux.Router, n *negroni.Negroni) {

	api := r.PathPrefix("/api/v1").Subrouter()

	n.Use(applicationJSON())
	n.Use(isAuth())

	api.Handle("/user/login", n.With(
		negroni.Wrap(authLogin()),
	)).Methods("POST", "OPTIONS")

	api.HandleFunc("/products", CreateProduct).Methods("POST")

	api.HandleFunc("/products", GetAllProducts).Methods("GET")

	api.HandleFunc("/products/{id}", GetById).Methods("GET")

	api.HandleFunc("/products/{id}", DeleteById).Methods("DELETE")

	api.HandleFunc("/products/{id}", UpdateById).Methods("PUT")

}
