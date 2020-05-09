package main

import (
	"log"
	"os"
	"testing"
)

func readFile() *os.File {

	file, err := os.Open("./fixtures/electric.log")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestParse(t *testing.T) {

	lines, err := Parse(readFile())

	if err != nil {
		log.Fatal(err)
	}

	if len(lines) != 10 {
		t.Errorf("expected 10 items to be parsed got %d", len(lines))
	}

}
