package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/calculate", nil)
	req.Header.Set("User-Access", "regularuser")

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("Обработчик должен был быть пропущен из-за отсутствия доступа superuser")
	})

	authMiddleware := AuthMiddleware(handler)
	authMiddleware.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Ожидается статус код %d, получено %d", http.StatusForbidden, w.Code)
	}
}
