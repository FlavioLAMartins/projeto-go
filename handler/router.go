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

	api.HandleFunc("/product", CreateProduct).Methods("POST")

	api.HandleFunc("/products", GetAllProducts).Methods("GET")

	api.HandleFunc("/product/{id}", GetById).Methods("GET")

	api.HandleFunc("/product/{id}", DeleteById).Methods("DELETE")

	api.HandleFunc("/product/{id}", UpdateById).Methods("PUT")

}
