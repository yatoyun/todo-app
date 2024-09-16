package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yatoyun/todo-app/api/infrastructure/router"
	"log"
	"net/url"
	"os"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", url.QueryEscape(dbUser), url.QueryEscape(dbPass), dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "true")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open connection to database: %v", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	// Start the server
	r := router.NewRouter(dbConn)
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
