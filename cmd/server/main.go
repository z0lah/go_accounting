package main

import (
	"fmt"
	"go_accounting/config"
	accountModule "go_accounting/internal/account"
	balancesheetModule "go_accounting/internal/balancesheet"
	dashboardModule "go_accounting/internal/dashboard"
	journalModule "go_accounting/internal/journal"
	ledgerModule "go_accounting/internal/ledger"
	"go_accounting/internal/middleware"
	reportModule "go_accounting/internal/report"
	token "go_accounting/internal/shared/token"
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
	// db = db.Debug()
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

	//init middleware
	tokenGen := token.NewJWTGenerator(cfg.SecretKey, 12*time.Hour)

	auth := middleware.JWTMiddleware(tokenGen)
	authAdmin := middleware.JWTMiddleware(tokenGen, "admin")

	//init dependencies
	userRepository := userModule.NewUserRepository(db)
	accountRepository := accountModule.NewAccountRepository(db)
	journalRepository := journalModule.NewJournalRepository(db)
	ledgerRepository := ledgerModule.NewLedgerRepository(db)
	reportRepository := reportModule.NewProfitLossRepository(db)
	balancesheetRepository := balancesheetModule.NewBalanceSheetRepository(db)
	dashboardRepository := dashboardModule.NewDashboardRepository(db)

	userUsecase := userModule.NewUserUsecase(userRepository, tokenGen)
	accountUsecase := accountModule.NewAccountUsecase(accountRepository)
	journalUsecase := journalModule.NewJournalUsecase(journalRepository, accountRepository)
	ledgerUsecase := ledgerModule.NewLedgerUsecase(ledgerRepository, accountRepository)
	reportUsecase := reportModule.NewProfitLossUsecase(reportRepository)
	balancesheetUsecase := balancesheetModule.NewBalanceSheetUsecase(balancesheetRepository)
	dashboardUsecase := dashboardModule.NewDashboardUsecase(dashboardRepository)

	//group route
	userGroup := app.Group("/auth")
	userAdminGroup := app.Group("/users", authAdmin)
	accountGroup := app.Group("/accounts", auth)
	journalGroup := app.Group("/journals", auth)
	ledgerGroup := app.Group("/ledgers", auth)

	userModule.NewUserHandler(userGroup, userUsecase)
	userModule.NewUserAdminHandler(userAdminGroup, userUsecase)
	accountModule.NewAccountHandler(accountGroup, accountUsecase)
	journalModule.NewJournalHandler(journalGroup, journalUsecase)
	ledgerModule.NewLedgerHandler(ledgerGroup, ledgerUsecase)
	reportModule.NewProfitLossHandler(app.Group("/reports", auth), reportUsecase)
	balancesheetModule.NewBalanceSheetHandler(app.Group("/balance-sheet", auth), balancesheetUsecase)
	dashboardModule.NewDashboardHandler(app.Group("/dashboard", auth), dashboardUsecase)

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}
