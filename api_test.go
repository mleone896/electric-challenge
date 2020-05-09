package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadEndpoint(t *testing.T) {
	file := readFile()

	var b bytes.Buffer
	var fw io.Writer
	var err error

	writer := multipart.NewWriter(&b)

	if fw, err = writer.CreateFormFile("file", file.Name()); err != nil {
		t.Errorf("Error creating writer %v", err)
	}

	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("error with io.Copy: %v", err)
	}
	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/upload", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	router := InitRoutes()

	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
