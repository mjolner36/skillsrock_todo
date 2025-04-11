package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjolner36/skillsrock_todo/internal/handler"
)

func TasksRoutes(app *fiber.App, handler *handler.TaskHandler) {
	api := app.Group("/v1")

	api.Get("/tasks", handler.GetAllTasks)
	api.Post("/tasks", handler.CreateTask)

	api.Get("/tasks/:id", handler.GetTask)
	api.Put("/tasks/:id", handler.UpdateTask)
	api.Delete("/tasks/:id", handler.DeleteTask)
}
