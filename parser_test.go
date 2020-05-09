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

func getTestMap() map[string]int {
	testData := make(map[string]int)
	testData["burger.electric.ai"] = 2
	testData["burger.letters.com"] = 1
	testData["d104.electric.ai"] = 3
	testData["unicomp6.electric.ai"] = 4

	return testData
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

func TestCalculatedUrlCounts(t *testing.T) {

	contents := readFile()
	counts := getCounts(contents)

	if !eq(counts, getTestMap()) {
		t.Errorf("Exptected %v to equal %v", counts, getTestMap())
	}

}

// helper func
func eq(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}
