package serverss

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSanitizeInput(t *testing.T) {
	testUserInput := "//$hello"
	expectedQuery := "hello"

	actualQuery := sanitizeInput(testUserInput)
	if actualQuery != expectedQuery {
		t.Errorf("Sanitize input failed want %v, got %v", expectedQuery, actualQuery)
	}
}

func TestWebServerHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/search?q=$$$", nil) // Make search return empty sicne $ chars are not allowed in search string
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchAndRespond)

	// Does the server respond well
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code %v", rr.Code)
	}

	// Does the response body exist
	expected := `{"Query":"","Result":[]}` // expecting an empty result since we only sent $$$ as query param
	actual := rr.Body.String()
	if expected != actual {
		t.Errorf("Body is incorrect expcted %v got %v", expected, actual)
	}

	// Also check the response header is application/jsopn
	expectedContentHeader := "application/json"
	actualContentHeader := rr.Header().Get("Content-Type")
	if expectedContentHeader != actualContentHeader {
		t.Errorf("Expected %v content header got %v", expectedContentHeader, actualContentHeader)
	}
}

func TestSearchOnlyAcceptsGET(t *testing.T) {
	req, err := http.NewRequest("POST", "/search?q=$$$", nil) // Make search return empty sicne $ chars are not allowed in search string
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchAndRespond)

	// Does the server respond well
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Handler returned a wrong status code %v", rr.Code)
	}

}
