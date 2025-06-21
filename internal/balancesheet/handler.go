package balancesheet

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BalanceSheetHandler struct {
	usecase BalanceSheetUsecase
}

func NewBalanceSheetHandler(router fiber.Router, usecase BalanceSheetUsecase) {
	handler := &BalanceSheetHandler{usecase: usecase}

	router.Get("/", handler.GetBalanceSheet)
}

func (h *BalanceSheetHandler) GetBalanceSheet(ctx *fiber.Ctx) error {
	dateStr := ctx.Query("date")
	if dateStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "date is required",
		})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format (YYYY-MM-DD)",
		})
	}

	balanceSheet, err := h.usecase.GetBalanceSheet(ctx.Context(), date)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(balanceSheet)
}
