package main

import (
	accountModule "go_accounting/internal/account"
	balancesheetModule "go_accounting/internal/balancesheet"
	dashboardModule "go_accounting/internal/dashboard"
	journalModule "go_accounting/internal/journal"
	ledgerModule "go_accounting/internal/ledger"
	reportModule "go_accounting/internal/report"
	"go_accounting/internal/shared/token"
	userModule "go_accounting/internal/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB, tokenGen token.TokenGenerator, auth fiber.Handler, authAdmin fiber.Handler) {
	// init repositories
	userRepository := userModule.NewUserRepository(db)
	accountRepository := accountModule.NewAccountRepository(db)
	journalRepository := journalModule.NewJournalRepository(db)
	ledgerRepository := ledgerModule.NewLedgerRepository(db)
	reportRepository := reportModule.NewProfitLossRepository(db)
	balancesheetRepository := balancesheetModule.NewBalanceSheetRepository(db)
	dashboardRepository := dashboardModule.NewDashboardRepository(db)

	// init usecases
	userUsecase := userModule.NewUserUsecase(userRepository, tokenGen)
	accountUsecase := accountModule.NewAccountUsecase(accountRepository)
	journalUsecase := journalModule.NewJournalUsecase(journalRepository, accountRepository)
	ledgerUsecase := ledgerModule.NewLedgerUsecase(ledgerRepository, accountRepository)
	reportUsecase := reportModule.NewProfitLossUsecase(reportRepository)
	balancesheetUsecase := balancesheetModule.NewBalanceSheetUsecase(balancesheetRepository)
	dashboardUsecase := dashboardModule.NewDashboardUsecase(dashboardRepository)

	// route groups
	userGroup := app.Group("/auth")                  // login, register
	userAdminGroup := app.Group("/users", authAdmin) // admin only
	accountGroup := app.Group("/accounts", auth)
	journalGroup := app.Group("/journals", auth)
	ledgerGroup := app.Group("/ledgers", auth)

	// register handlers
	userModule.NewUserHandler(userGroup, userUsecase)
	userModule.NewUserAdminHandler(userAdminGroup, userUsecase)
	accountModule.NewAccountHandler(accountGroup, accountUsecase)
	journalModule.NewJournalHandler(journalGroup, journalUsecase)
	ledgerModule.NewLedgerHandler(ledgerGroup, ledgerUsecase)
	reportModule.NewProfitLossHandler(app.Group("/reports", auth), reportUsecase)
	balancesheetModule.NewBalanceSheetHandler(app.Group("/balance-sheet", auth), balancesheetUsecase)
	dashboardModule.NewDashboardHandler(app.Group("/dashboard", auth), dashboardUsecase)

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
