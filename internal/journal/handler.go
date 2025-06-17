package journal

import (
	"context"
	"go_accounting/internal/shared/pagination"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type JournalHandler struct {
	usecase   JournalUsecase
	validator *validator.Validate
}

func NewJournalHandler(router fiber.Router, usecase JournalUsecase) {
	handler := &JournalHandler{
		usecase:   usecase,
		validator: validator.New(),
	}
	//Route
	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)
	router.Get("/:id", handler.GetByID)
}

func (h *JournalHandler) Create(ctx *fiber.Ctx) error {
	var input CreateJournalInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.validator.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := h.usecase.Create(context.Background(), input)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(result)
}

func (h *JournalHandler) GetAll(ctx *fiber.Ctx) error {
	pq := pagination.FromQuery(ctx)

	result, total, err := h.usecase.GetAll(context.Background(), pq.Page, pq.Limit)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp := pagination.BuildPagedResponse(result, pq.Page, pq.Limit, total)

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (h *JournalHandler) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	result, err := h.usecase.GetByID(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
