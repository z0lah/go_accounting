package main

import (
	"fmt"
	"go_accounting/config"
	accountModule "go_accounting/internal/account"
	journalModule "go_accounting/internal/journal"
	ledgerModule "go_accounting/internal/ledger"
	reportModule "go_accounting/internal/report"
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
	db = db.Debug()
	db.AutoMigrate(
		&userModule.User{},
		&accountModule.Account{},
		&journalModule.Journal{},
		&journalModule.JournalDetail{},
	)
	// seed dummy data
	// if err := seeder.SeedJournals(db); err != nil {
	// 	log.Fatal("Gagal seed:", err)
	// }

	//init fiber app
	app := fiber.New()

	//init dependencies
	userRepository := userModule.NewUserRepository(db)
	accountRepository := accountModule.NewAccountRepository(db)
	journalRepository := journalModule.NewJournalRepository(db)
	ledgerRepository := ledgerModule.NewLedgerRepository(db)
	reportRepository := reportModule.NewProfitLossRepository(db)

	tokenGen := tokenModule.NewJWTGenerator(cfg.SecretKey, 12*time.Hour)

	userUsecase := userModule.NewUserUsecase(userRepository, tokenGen)
	accountUsecase := accountModule.NewAccountUsecase(accountRepository)
	journalUsecase := journalModule.NewJournalUsecase(journalRepository, accountRepository)
	ledgerUsecase := ledgerModule.NewLedgerUsecase(ledgerRepository, accountRepository)
	reportUsecase := reportModule.NewProfitLossUsecase(reportRepository)

	//group route
	userGroup := app.Group("/users")
	accountGroup := app.Group("/accounts")
	journalGroup := app.Group("/journals")
	ledgerGroup := app.Group("/ledgers")

	userModule.NewUserHandler(userGroup, userUsecase)
	accountModule.NewAccountHandler(accountGroup, accountUsecase)
	journalModule.NewJournalHandler(journalGroup, journalUsecase)
	ledgerModule.NewLedgerHandler(ledgerGroup, ledgerUsecase)
	reportModule.NewProfitLossHandler(app.Group("/reports"), reportUsecase)

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// TODO: Setup routes, middleware, etc.

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}
