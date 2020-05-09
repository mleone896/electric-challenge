package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var port string
	flag.StringVar(&port, "port", ":8080", "Port to bind httpserver")

	flag.Parse()

	router := InitRoutes()

	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(port, router))
}
