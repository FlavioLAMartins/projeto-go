package main

import (
	"deposito/handler"
)

func main() {
	handler.ConfigServer()

	//fs := http.FileServer(http.Dir("dist/spa"))
	//fs := http.FileServer(http.Dir("dist/spa/"))

	//fmt.Println("CHEGOU AQ")

	//router.Handle("/webui", http.StripPrefix("/webui", fs))

	//fmt.Println("Escutando na porta 5000")
	//log.Fatal(http.ListenAndServe(":5000", router))

}
