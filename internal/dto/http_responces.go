package dto

import (
	"github.com/gofiber/fiber/v2"
)

const (
	BadResponse        = "StatusBadRequest"
	ServiceUnavailable = "SERVICE_UNAVAILABLE"
	NotFound           = "NOT_FOUND"
	InternalError      = "Service is currently unavailable. Please try again later."
	OKResponse         = "OK"
)

type Response struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Error struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func BadResponseError(ctx *fiber.Ctx, desc string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: BadResponse,
			Desc: desc,
		},
	})
}

func InternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: ServiceUnavailable,
			Desc: InternalError,
		},
	})
}

func NotFoundError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: NotFound,
			Desc: message,
		},
	})
}

func CorrectResponse(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Status: OKResponse,
		Data:   data,
	})
}
