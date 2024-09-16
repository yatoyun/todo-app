package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yatoyun/todo-app/api/infrastructure/router"
	"log"
	"os"
	"time"
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

func connectDB() *sqlx.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	c := mysql.Config{
		DBName:    dbName,
		User:      dbUser,
		Passwd:    dbPass,
		Addr:      fmt.Sprintf("%s:%s", dbHost, dbPort),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sqlx.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open connection to database: %v", err)
	}
	return db
}

func main() {
	// Connect to the database
	dbConn := connectDB()
	err := dbConn.Ping()
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
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
