package user

import (
	"context"
	"go_accounting/internal/shared/pagination"
	"go_accounting/internal/shared/validation"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase UserUsecase
}

func NewUserHandler(router fiber.Router, usecase UserUsecase) {
	handler := &UserHandler{usecase: usecase}

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
	router.Get("/", handler.GetAll)
	router.Get("/not-active", handler.GetNotActive)
	router.Patch("/:id/role", handler.UpdateRole)
	router.Patch("/:id/status", handler.UpdateStatus)
}

var validate = validator.New()

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var input RegisterInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON input",
		})
	}
	//validate input
	if err := validate.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//compare password
	if err := validation.ComparePassword(input.Password, input.ConfirmPassword); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err := h.usecase.Register(context.Background(), input)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(user)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var input LoginInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON input",
		})
	}
	if err := validate.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	authResponse, err := h.usecase.Login(context.Background(), input)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(authResponse)
}

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	pq := pagination.FromQuery(ctx)

	data, total, err := h.usecase.GetAll(context.Background(), pq.Page, pq.Limit)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp := pagination.BuildPagedResponse(data, pq.Page, pq.Limit, total)
	return ctx.Status(http.StatusOK).JSON(resp)
}

func (h *UserHandler) GetNotActive(ctx *fiber.Ctx) error {
	users, err := h.usecase.GetNotActive(context.Background())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(users)
}

func (h *UserHandler) UpdateStatus(ctx *fiber.Ctx) error {
	var input UpdateStatusInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON input",
		})
	}
	if err := validate.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.usecase.UpdateStatus(ctx.Context(), ctx.Params("id"), input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}

func (h *UserHandler) UpdateRole(ctx *fiber.Ctx) error {
	var input UpdateRoleInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON input",
		})
	}
	if err := validate.Struct(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.usecase.UpdateRole(ctx.Context(), ctx.Params("id"), input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{})
}
