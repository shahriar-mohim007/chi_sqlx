package domain

type Task struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type TaskRepository interface {
	GetTaskByID(id int64) (*Task, error)
	CreateTask(task *Task) error
}
