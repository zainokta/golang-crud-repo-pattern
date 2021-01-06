package router

import (
	"database/sql"
	"user-crud/controllers"
	"user-crud/middlewares"
	"user-crud/repositories"

	"github.com/go-chi/chi"
)

func userRouter(db *sql.DB) chi.Router {

	u := chi.NewRouter()
	u.Use(middlewares.Authorization)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	u.Get("/", userController.GetUser)
	u.Get("/{id}", userController.GetUserById)
	u.Put("/{id}", userController.Update)
	u.Delete("/{id}", userController.Delete)

	return u
}
