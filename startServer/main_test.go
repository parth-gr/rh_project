package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", "localhost:5000/", nil)

	if err != nil {
		t.Fatalf("request not created %v", err)
	}

	rec := httptest.NewRecorder()

	index(rec, request)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("excepted status ok, got %v ", res.Status)
	}

}
