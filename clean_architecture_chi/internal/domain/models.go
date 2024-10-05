//Entities (Domain):

// Pure Go structs that represent core business objects (e.g., Task).
// Domain interfaces are defined here, decoupling your code from external frameworks like the database or web libraries.

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
