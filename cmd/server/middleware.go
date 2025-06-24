package main

import (
	"go_accounting/internal/middleware"
	"go_accounting/internal/shared/token"

	"github.com/gofiber/fiber/v2"
)

func InitAuthMiddleware(tokenParser token.TokenParser) fiber.Handler {
	return middleware.JWTMiddleware(tokenParser)
}

func InitAdminMiddleware(tokenParser token.TokenParser) fiber.Handler {
	return middleware.JWTMiddleware(tokenParser, "admin")
}
