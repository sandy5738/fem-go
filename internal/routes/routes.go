package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sandy5738/fem-go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
