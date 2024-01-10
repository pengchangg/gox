package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/go-playground/assert/v2"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
