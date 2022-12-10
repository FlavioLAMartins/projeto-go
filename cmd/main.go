package main

import (
	"deposito/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	handler.RegisterAPI(r)
	fmt.Println("Server on :5000 ")
	log.Fatal(http.ListenAndServe(":5000", r))

}
