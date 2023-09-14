package handlers

import (
	"Softweather_test/internal/middleware"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiCalculateHandlerWithAuth(t *testing.T) {
	requestBody := []byte(`2+1+1+1`)
	req, err := http.NewRequest("POST", "/api/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Access", "superuser")

	rr := httptest.NewRecorder()
	handler := middleware.AuthMiddleware(http.HandlerFunc(ApiCalculateHandler))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус код %v, получено %v", http.StatusOK, status)
	}

	expectedResponseBody := `{"result":5}`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("Ожидается тело ответа %v, получено %v", expectedResponseBody, rr.Body.String())
	}
}

func TestApiCalculateHandlerWithoutAuth(t *testing.T) {
	requestBody := []byte(`2+1+1+1`)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.AuthMiddleware(http.HandlerFunc(ApiCalculateHandler))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("Ожидается статус код %v, получено %v", http.StatusForbidden, status)
	}
}
