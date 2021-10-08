package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler1(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = "Hello, HTTPサーバ"
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}
}

func TestHandlerTest2(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?p=Gopher", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = "Hello, Gopher"
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}
}
