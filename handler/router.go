package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigServer() {
	router := mux.NewRouter().StrictSlash(true)
	ConfigRoutes(router)
	fmt.Println("Server on :5000 ")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func ConfigRoutes(router *mux.Router) {

	router.HandleFunc("/products", CreateProduct).Methods("POST")

	router.HandleFunc("/products", GetAllProducts).Methods("GET")

	router.HandleFunc("/products/{id}", GetById).Methods("GET")

	router.HandleFunc("/products/{id}", DeleteById).Methods("DELETE")

	router.HandleFunc("/products/{id}", UpdateById).Methods("PUT")

}
