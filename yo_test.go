package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testClient *Client
)

func init() {
	testClient = NewClient("some_token")
	return
}

func TestYoAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(yoAllHandler))
	defer server.Close()
	YO_API = server.URL

	if err := testClient.YoAll(); err != nil {
		t.Fatal(err)
	}
}

func yoAllHandler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusCreated)
}

func TestYoUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(yoUserHandler))
	defer server.Close()
	YO_API = server.URL

	if err := testClient.YoUser("some_user"); err != nil {
		t.Fatal(err)
	}
}

func yoUserHandler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusCreated)
}