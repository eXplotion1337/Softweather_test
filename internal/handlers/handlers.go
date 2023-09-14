package handlers

import (
	"Softweather_test/internal/handlers/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ApiCalculateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := service.Сalculate(string(body))
	if err != nil {
		log.Println("Ошибка вычислений:", err)
		http.Error(w, "Ошибка вычислений", http.StatusBadRequest)
		return
	}

	response := struct {
		Result int `json:"result"`
	}{
		Result: result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Ошибка маршалинга JSON:", err)
		http.Error(w, "Ошибка маршалинга JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
