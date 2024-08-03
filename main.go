package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santosh-shetty/blog/pkg/config"
	"github.com/santosh-shetty/blog/pkg/routes"
)

func main() {
	config.DBConnect()

	defer config.DBClose()

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	http.Handle("/", r)
	routes.Routes(r)
	fmt.Println("localhost:9000")
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
