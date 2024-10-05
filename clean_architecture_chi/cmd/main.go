// cmd/main.go
package main

import (
	"clean-architecture/api/handler"
	"clean-architecture/config"
	"clean-architecture/infrastructure/router"
	"clean-architecture/internal/usecase"

	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()

	// Create the database connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connect to the database
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	taskRepo := database.NewSqlxTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	r := router.NewRouter(taskHandler)
	http.ListenAndServe(":8080", r)
}
