package account

import (
	"context"
	"go_accounting/internal/shared/pagination"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AccountHandler struct {
	usecase   AccountUsecase
	validator *validator.Validate
}

func NewAccountHandler(router fiber.Router, usecase AccountUsecase) {
	handler := &AccountHandler{
		usecase:   usecase,
		validator: validator.New(),
	}
	//Route
	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)
	router.Get("/:id", handler.GetByID)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}

func (h *AccountHandler) Create(ctx *fiber.Ctx) error {
	var input CreateAccountInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	if err := h.validator.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	account, err := h.usecase.Create(context.Background(), input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(account)
}

func (h *AccountHandler) GetAll(ctx *fiber.Ctx) error {
	pq := pagination.FromQuery(ctx)

	data, total, err := h.usecase.GetAll(context.Background(), pq.Page, pq.Limit)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := pagination.BuildPagedResponse(data, pq.Page, pq.Limit, total)
	return ctx.Status(http.StatusOK).JSON(resp)
}

func (h *AccountHandler) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	account, err := h.usecase.GetByID(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(account)
}

func (h *AccountHandler) Update(ctx *fiber.Ctx) error {
	var input UpdateAccountInput
	id := ctx.Params("id")

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := h.validator.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	account, err := h.usecase.Update(context.Background(), id, input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(account)
}

func (h *AccountHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	if err := h.usecase.Delete(context.Background(), id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "account deleted"})
}
