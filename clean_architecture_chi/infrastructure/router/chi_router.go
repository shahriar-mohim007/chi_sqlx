//Infrastructure (Router and Database Connection):

//This layer connects external frameworks, like chi for routing and sqlx for the database.

// infrastructure/router/chi_router.go
package router

import (
	"clean-architecture/api/handler"

	"clean-architecture/middleware"

	"github.com/go-chi/chi"
)

func NewRouter(taskHandler *handler.TaskHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/tasks/{id}", taskHandler.GetTask)
	r.Post("/tasks", taskHandler.CreateTask)
	return r
}
