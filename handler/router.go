package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func RegisterAPI(r *mux.Router) {

	n := negroni.Classic()
	n.Use(applicationJSON())
	n.Use(basicAuth())

	n.UseHandler(r)

	r.HandleFunc("/products", CreateProduct).Methods("POST")

	r.HandleFunc("/products", GetAllProducts).Methods("GET")

	r.HandleFunc("/products/{id}", GetById).Methods("GET")

	r.HandleFunc("/products/{id}", DeleteById).Methods("DELETE")

	r.HandleFunc("/products/{id}", UpdateById).Methods("PUT")

}
