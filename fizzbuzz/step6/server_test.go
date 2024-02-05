package step6

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/?number=15", nil)
	w := httptest.NewRecorder()

	// Act
	FizzBuzzHandler(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Fatalf("HTTP Code = %d; want %d", w.Code, http.StatusOK)
	}
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}

	result := string(data)
	if result != "FizzBuzz" {
		t.Fatalf("GET /?number=15 = %s; want FizzBuzz", result)
	}
}
