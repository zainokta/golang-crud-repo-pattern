package router

import (
	"net/http"
	"os"
	"path/filepath"
	"user-crud/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//Router is main router
func Router(config *config.DBConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/user", userRouter(config.Db))
	r.Mount("/auth", authRouter(config.Db))
	r.Mount("/upload", uploadRouter())

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "upload"))

	fileServer(r, "/upload", filesDir)

	return r
}
