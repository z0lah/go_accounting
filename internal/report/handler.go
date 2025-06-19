package report

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProfitLossHandler struct {
	usecase ProfitLossUsecase
}

func NewProfitLossHandler(router fiber.Router, usecase ProfitLossUsecase) {
	handler := &ProfitLossHandler{usecase}

	router.Get("/profit-loss", handler.Get)
}

func (h *ProfitLossHandler) Get(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")

	var start, end time.Time
	var err error

	if startStr != "" {
		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_date format (YYYY-MM-DD)",
			})
		}
	}
	if endStr != "" {
		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_date format (YYYY-MM-DD)",
			})
		}
	}

	result, err := h.usecase.GetReport(ctx.Context(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
