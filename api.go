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

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(ReceiveFile)

	return router

}

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
		counts[line.RemoteHost] += 1
	}

	return counts

}
