package config

import (
	"log"
	"os"
)

type Config struct {
	Auth0Domain       string
	Auth0Audience     string
	Auth0ClientID     string
	Auth0ClientSecret string
	DatabaseURL       string
}

func LoadConfig() *Config {
	cfg := &Config{
		Auth0Domain:       getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience:     getEnv("AUTH0_AUDIENCE", ""),
		Auth0ClientID:     getEnv("AUTH0_CLIENT_ID", ""),
		Auth0ClientSecret: getEnv("AUTH0_CLIENT_SECRET", ""),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
	}

	// 必須の環境変数が設定されているか確認
	if cfg.Auth0Domain == "" || cfg.Auth0Audience == "" || cfg.DatabaseURL == "" {
		log.Fatal("必須の環境変数が設定されていません")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
