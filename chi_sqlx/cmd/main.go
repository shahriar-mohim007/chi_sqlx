package cmd

import (
	"chi_sqlx_project/pkg/config"
	"chi_sqlx_project/pkg/infrastructure/database"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	_, err := database.NewSqlxTaskRepository(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize task repository: %v", err)
	}

}
