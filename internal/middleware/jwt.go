package middleware

import (
	"go_accounting/internal/shared/token"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(tokenParser token.TokenParser, role ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		_, claims, err := tokenParser.Parse(tokenStr)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// simpan context
		ctx.Locals("user_id", claims["user_id"].(string))
		ctx.Locals("email", claims["email"].(string))
		ctx.Locals("role", claims["role"].(string))

		if len(role) > 0 {
			userRole := claims["role"]
			allowed := false
			for _, r := range role {
				if r == userRole {
					allowed = true
					break
				}
			}

			if !allowed {
				return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
					"error": "forbidden: insufficient privileges",
				})
			}
		}

		return ctx.Next()
	}
}
