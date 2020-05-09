package main

import (
	"fmt"
	"log"
	"net/http"
)

// A data structure to hold a key/value pair.
type pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by Value.
type pairList []pair

func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	router := InitRoutes()

	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}
