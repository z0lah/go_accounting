package ledger

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LedgerHandler struct {
	ledgerUsecase LedgerUseCase
}

func NewLedgerHandler(router fiber.Router, usecase LedgerUseCase) {
	handler := &LedgerHandler{ledgerUsecase: usecase}

	router.Get("/", handler.GetAll)
}

func (h *LedgerHandler) GetAll(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start_date")
	endStr := ctx.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startStr != "" {
		startDate, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_date format (YYYY-MM-DD)",
			})
		}
	}
	if endStr != "" {
		endDate, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_date format (YYYY-MM-DD)",
			})
		}
	}

	data, err := h.ledgerUsecase.GetAll(context.Background(), startDate, endDate)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(data)
}
