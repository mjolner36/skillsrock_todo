package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mjolner36/skillsrock_todo/config"
	"github.com/mjolner36/skillsrock_todo/internal/handler"
	"github.com/mjolner36/skillsrock_todo/internal/repo"
	"github.com/mjolner36/skillsrock_todo/internal/routes"
	"github.com/mjolner36/skillsrock_todo/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := godotenv.Load(".env")

	var cfg config.AppConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err, "failed to load .env file")
	}

	repository := repo.NewInMemoryRepo()
	service := services.NewService(repository)
	taskHandler := handler.NewTaskHandler(service)

	app := fiber.New()
	routes.TasksRoutes(app, taskHandler)

	//app.Listen(cfg.PORT)
	go func() {
		if err := app.Listen(cfg.PORT); err != nil {
			log.Panicf("server failed to start: %v", err)
		}
	}()
	log.Println("Server started on port", cfg.PORT)

	// Ожидание сигнала завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	log.Println("Gracefully shutting down...")

}
