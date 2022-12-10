package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func ConfigServer() {
	n := negroni.Classic()
	n.Use(applicationJSON())
	n.Use(basicAuth())
	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)
	RegisterAPI(r)
	fmt.Println("Server on :5000 ")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func RegisterAPI(router *mux.Router) {

	router.HandleFunc("/products", CreateProduct).Methods("POST")

	router.HandleFunc("/products", GetAllProducts).Methods("GET")

	router.HandleFunc("/products/{id}", GetById).Methods("GET")

	router.HandleFunc("/products/{id}", DeleteById).Methods("DELETE")

	router.HandleFunc("/products/{id}", UpdateById).Methods("PUT")

}
