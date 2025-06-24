package main

import (
	"go_accounting/config"
	"go_accounting/internal/shared/token"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.NewConfig()

	db := InitDB(cfg)
	app := fiber.New()

	// middleware global
	app.Use(logger.New())
	app.Use(recover.New())

	// jwt middleware
	tokenGen := token.NewJWTGenerator(cfg.SecretKey, 12*time.Hour)
	auth := InitAuthMiddleware(tokenGen)
	authAdmin := InitAdminMiddleware(tokenGen)

	InitRoutes(app, db, tokenGen, auth, authAdmin)

	log.Fatal(app.Listen(":" + cfg.Port))
}
