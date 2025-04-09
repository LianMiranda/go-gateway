package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/repository"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/service"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/web/server"
	_ "github.com/lib/pq"
)
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		getEnv("DB_HOST", "localhost"), 
		getEnv("DB_PORT", "5432"), 
		getEnv("DB_USER", "postgres"), 
		getEnv("DB_PASSWORD", "1234"), 
		getEnv("DB_NAME", "gateway"), 
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "3000")
	server := server.NewServer(accountService, port)
	server.ConfigureRoutes()	

	server.Start()
	if err := server.Start(); err != nil{
		log.Fatal("Error starting server: ", err)
	}
}
