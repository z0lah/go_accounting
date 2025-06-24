package main

import (
	"fmt"
	"go_accounting/config"
	accountModule "go_accounting/internal/account"
	journalModule "go_accounting/internal/journal"
	userModule "go_accounting/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(
		&userModule.User{},
		&accountModule.Account{},
		&journalModule.Journal{},
		&journalModule.JournalDetail{},
	)

	return db
}
