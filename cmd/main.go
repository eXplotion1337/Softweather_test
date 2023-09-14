package main

import (
	"Softweather_test/internal/app"
	"Softweather_test/internal/config"
	"log"
)

func main() {
	runConfig, err := config.InitConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфига")
	}

	err = app.Run(runConfig)
	if err != nil {
		log.Fatal("Ошибка старта сервера")
	}
}
