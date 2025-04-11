package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjolner36/skillsrock_todo/internal/dto"
	"github.com/mjolner36/skillsrock_todo/internal/entity"
	"github.com/mjolner36/skillsrock_todo/internal/services"
)

type TaskHandler struct {
	service *services.Service
}

func NewTaskHandler(service *services.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := h.service.GetAllTasks(ctx.Context())
	if err != nil {
		return dto.InternalServerError(ctx)
	}
	return dto.CorrectResponse(ctx, tasks)
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {

	var req dto.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.BadResponseError(ctx, "Invalid request body")
	}

	task := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	err := h.service.CreateTask(ctx.Context(), task)
	if err != nil {
		return dto.InternalServerError(ctx)
	}
	return dto.CorrectResponse(ctx, req)
}

func (h *TaskHandler) GetTask(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadResponseError(ctx, "Bad Request")
	}

	task, err := h.service.GetTask(ctx.Context(), id)
	if err != nil {
		return dto.InternalServerError(ctx)
	}
	return dto.CorrectResponse(ctx, task)
	//return ctx.JSON(task)
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadResponseError(ctx, "Invalid task ID")
	}

	var req dto.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.BadResponseError(ctx, "Invalid request body")
	}

	task := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	err = h.service.UpdateTask(ctx.Context(), task, id)

	if err != nil {
		return dto.InternalServerError(ctx)
	}
	return dto.CorrectResponse(ctx, req)
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadResponseError(ctx, "Invalid task ID")
	}

	err = h.service.DeleteTask(ctx.Context(), id)
	if err != nil {
		return dto.InternalServerError(ctx)
	}
	return dto.CorrectResponse(ctx, "Task deleted")
}
