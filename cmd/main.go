package main

import (
	"emergencyNotification/internal/config"
	"emergencyNotification/pkg/logger"
)

func main() {

	// Инициализация конфига

	cfg := config.MustLoad()

	// Инифиализация логгера

	log := logger.SetupLogger(cfg.Env)

	// Инициализация базы данных

	// Инициализация роутеров

	// Запуск сервера

}
