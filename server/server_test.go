package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestServer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/players/Mahir", nil)
	res := httptest.NewRecorder()

	ServeHTTP(res,req)
	if res.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", res.Code, http.StatusOK)
	}
	if res.Body.String()!= 200{
		t.Errorf("got body %q, want %q", res.Body.String(), "20")
	}

}