package dashboard

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	usecase DashboardUsecase
}

func NewDashboardHandler(router fiber.Router, usecase DashboardUsecase) {
	handler := &DashboardHandler{usecase: usecase}

	router.Get("/", handler.GetSummary)
}

func (h *DashboardHandler) GetSummary(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")

	var start, end time.Time
	var err error

	if startStr == "" || endStr == "" {
		now := time.Now()
		start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		end = start.AddDate(0, 1, -1)
		log.Println(start, end)
	} else {
		start, err = time.Parse("2006-01-02", startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_date format (YYYY-MM-DD)",
			})
		}
		end, err = time.Parse("2006-01-02", endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_date format (YYYY-MM-DD)",
			})
		}
		if end.Before(start) {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "end_date must be after start_date",
			})
		}
	}

	summary, err := h.usecase.GetSummary(context.Background(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(summary)
}
