package database

import (
	"chi_sqlx_project/pkg/config"
	"chi_sqlx_project/pkg/internal/domain"
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
)

type SqlxTaskRepository struct {
	db *sqlx.DB
}

var (
	once       sync.Once
	repository *SqlxTaskRepository
)

func NewSqlxTaskRepository(cfg config.Config) (domain.TaskRepository, error) {
	var err error
	databaseUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	once.Do(func() {
		db, dbErr := sqlx.Connect("pgx", databaseUrl)
		if dbErr != nil {
			err = dbErr
			log.Fatalf("Database Connection Error: %v", err)
			return
		}
		if pingErr := db.Ping(); pingErr != nil {
			err = pingErr
			log.Fatalf("Database Connection Error: %v", err)
			return
		}

		repository = &SqlxTaskRepository{db: db}
	})

	return repository, err
}

func (r *SqlxTaskRepository) GetTaskByID(id int64) (*domain.Task, error) {
	task := &domain.Task{}
	err := r.db.Get(task, "SELECT id, title, description FROM tasks WHERE id = $1", id)
	return task, err
}

func (r *SqlxTaskRepository) CreateTask(task *domain.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description) VALUES ($1, $2)", task.Title, task.Description)
	return err
}
