package app

import (
	"Softweather_test/internal/config"
	"Softweather_test/internal/handlers"
	"Softweather_test/internal/middleware"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(config *config.Config) error {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware)

	r.Post("/api/calculate", func(w http.ResponseWriter, r *http.Request) {
		handlers.ApiCalculateHandler(w, r)
	})

	server := &http.Server{
		Addr:    config.HttpConfig.Port,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Запуск сервера...")
		log.Printf("Сервер запущен на порту  %s", config.HttpConfig.Port)
		log.Printf("Сервер запущен на адресе  %s", config.HttpConfig.Host)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Ошибка сервера: %v\n", err)
		}
	}()

	<-stop
	log.Println("Получен сигнал завершения. Завершение работы сервера...")

	log.Println("Завершение работы сервера через контекст...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Ошибка graceful shutdown: %v\n", err)
		return err
	}
	end := time.Now()

	log.Printf("Сервер завершил работу за %v секунд\n", end.Sub(start).Seconds())
	return nil
}
