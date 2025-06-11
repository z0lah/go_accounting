package user

import (
	"context"
	"net/http"

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
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var input RegisterInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
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
			"error": "Invalid input",
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
	users, err := h.usecase.GetAll(context.Background())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(users)
}
