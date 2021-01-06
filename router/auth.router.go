package router

import (
	"database/sql"
	"user-crud/controllers"
	"user-crud/repositories"

	"github.com/go-chi/chi"
)

func authRouter(db *sql.DB) chi.Router {

	a := chi.NewRouter()

	userRepository := repositories.NewUserRepository(db)
	authController := controllers.NewAuthController(userRepository)
	a.Post("/login", authController.Login)
	a.Post("/register", authController.Register)

	return a
}
