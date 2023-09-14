package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAccess := r.Header.Get("User-Access")
		if userAccess != "superuser" {
			log.Println("Доступ запрещен.")
			http.Error(w, "Доступ запрещен.", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
