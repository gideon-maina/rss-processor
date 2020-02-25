package serverss

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJWTToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-token", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetToken)

	// Does the server respond well
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code %v", rr.Code)
	}

	// Does the response body exist and contain the token
	var tokenResponse Token
	respBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(respBytes, &tokenResponse)
	if tokenResponse.JWTTokenValue == "" {
		t.Errorf("Expected a valid JWT Token got %v", tokenResponse.JWTTokenValue)
	}

	// Also check the response header is application/jsopn
	expectedContentHeader := "application/json"
	actualContentHeader := rr.Header().Get("Content-Type")
	if expectedContentHeader != actualContentHeader {
		t.Errorf("Expected %v content header got %v", expectedContentHeader, actualContentHeader)
	}
}
