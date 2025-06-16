package main

import (
	"fmt"
	"go_accounting/config"
	accountModule "go_accounting/internal/account"
	tokenModule "go_accounting/internal/shared/token"
	userModule "go_accounting/internal/user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// init config
	cfg := config.NewConfig()

	// connect db
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// migrate
	db.AutoMigrate(&userModule.User{}, &accountModule.Account{})

	//init fiber app
	app := fiber.New()

	//init dependencies
	userRepository := userModule.NewUserRepository(db)
	accountRepository := accountModule.NewAccountRepository(db)
	tokenGen := tokenModule.NewJWTGenerator(cfg.SecretKey, 12*time.Hour)
	userUsecase := userModule.NewUserUsecase(userRepository, tokenGen)
	accountUsecase := accountModule.NewAccountUsecase(accountRepository)

	//group route
	userGroup := app.Group("/users")
	accountGroup := app.Group("/accounts")

	userModule.NewUserHandler(userGroup, userUsecase)
	accountModule.NewAccountHandler(accountGroup, accountUsecase)

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// TODO: Setup routes, middleware, etc.

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}
