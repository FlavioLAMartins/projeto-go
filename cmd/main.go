package main

import (
	"deposito/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.RegisterAPI(r, n)
	fmt.Println("Server on :5000 ")
	log.Fatal(http.ListenAndServe(":5000", r))
}
