package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"
)

const LAYOUT = "02/Jan/2006:15:04:04 - 0700"

// Parser : forward thinking for multiple formats of logs
type Parser interface {
	Parse(parser, lines string) ([]Line, error)
}

// Line : Represents a line shape coming in from handler
type Line struct {
	RemoteHost string
	Time       string
	Method     string
	Request    string
	Version    string
	Status     int
	Bytes      int
}

// String : Implement the stringer interface
func (li *Line) String() string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s\t%d\t%d",
		li.RemoteHost,
		li.Time,
		li.Method,
		li.Request,
		li.Version,
		li.Status,
		li.Bytes,
	)
}

func toInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic("Failed to convert to Integer")
	}
	return res
}

func getTime(timestamp string) time.Time {
	res, err := time.Parse(LAYOUT, timestamp)
	if err != nil {
		panic("Invalid Time")
	}
	return res
}

// first iteration isn't going to make use of it but put get line into
// its own method so that we can take advantage of channels
func getLine(line string) (Line, error) {

	regex := regexp.MustCompile(`(?P<remote_ip>\S*)\s-\s(?P<requesting_user>\S*)\s\[(?P<Timestamp>.*?)\]\s\"(?P<Method>\S*)\s*(?P<Request>\S*)\s*(HTTP\/)*(?P<http_version>.*?)\"\s(?P<response_code>\d{3})\s(?P<Size>\S*)`)
	match := regex.FindStringSubmatch(line)

	item := Line{
		RemoteHost: match[1],
		Time:       match[3],
		Method:     match[4],
		Request:    match[5],
		Version:    match[7],
		Status:     toInt(match[8]),
		Bytes:      toInt(match[9]),
	}
	return item, nil
}

// caller needs to close filehandle
func readLines(file io.Reader) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Parse : implement parser interface
func Parse(file io.Reader) ([]Line, error) {
	var entries []Line
	lines, err := readLines(file)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		entity, err := getLine(line)
		if err != nil {
			return []Line{}, err
		} else {
			entries = append(entries, entity)
		}
	}

	return entries, nil
}
