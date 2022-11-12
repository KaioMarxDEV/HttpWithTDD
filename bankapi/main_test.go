package main

import (
	"net/http"
	"testing"
)

func TestStatementsEndpoint(t *testing.T) {
	resp, _ := http.Get("http://localhost:8000/statement?number=1001")

	if resp.StatusCode != 200 {
		t.Error("Should be able to get statement from account created")
	}
}
