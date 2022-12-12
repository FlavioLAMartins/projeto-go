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
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.RegisterAPI(r, n)
	fmt.Println("Server on :8080 ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
