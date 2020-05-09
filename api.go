package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// InitRoutes : returns a router with endpoints mounted
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api/v1")
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(ReceiveFile)

	return router

}

// ReceiveFile : HttpHandler that accepts a file parses it and runs calculations
// on its contents
func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // limit your max input length!

	file, header, err := r.FormFile("file")

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	defer file.Close()

	name := strings.Split(header.Filename, ".")

	log.Println("File Name:", name)

	counts := getCounts(file)

	if json.NewEncoder(w).Encode(counts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// maybe we want to save the file?
func createFile(buf bytes.Buffer, name string) error {

	fo, err := os.Create(name)
	if err != nil {
		return err
	}

	defer fo.Close()

	if _, err := buf.WriteTo(fo); err != nil {
		return err
	}

	return err
}

func getCounts(contents io.Reader) map[string]int {
	lines, err := Parse(contents)

	if err != nil {
		log.Println(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		//  If we were to run this concurrently I would put a mutex to protect this
		counts[line.RemoteHost]++
	}

	return counts

}
