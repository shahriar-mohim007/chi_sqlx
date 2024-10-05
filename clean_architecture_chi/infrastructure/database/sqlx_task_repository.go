//Repository (Data Access):

// The repository implementation that uses sqlx for database interactions. This layer translates Go structs to SQL.

package database

import (
	"clean-architecture/internal/domain"

	"github.com/jmoiron/sqlx"
)

type SqlxTaskRepository struct {
	db *sqlx.DB
}

func NewSqlxTaskRepository(db *sqlx.DB) domain.TaskRepository {
	return &SqlxTaskRepository{db}
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
