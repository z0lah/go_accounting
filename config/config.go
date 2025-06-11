package config

import "os"

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	Port       string
	SecretKey  string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		Port:       os.Getenv("PORT"),
		SecretKey:  os.Getenv("JWT_SECRET"),
	}
}
