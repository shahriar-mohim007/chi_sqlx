//Use Case (Application Logic):

// Implements the application's specific business rules and interacts with repository interfaces.

// internal/usecase/task_usecase.go
package usecase

import "clean-architecture/internal/domain"

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (uc *TaskUsecase) GetTask(id int64) (*domain.Task, error) {
	return uc.repo.GetTaskByID(id)
}

func (uc *TaskUsecase) CreateTask(task *domain.Task) error {
	return uc.repo.CreateTask(task)
}
