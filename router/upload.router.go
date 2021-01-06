package router

import (
	"log"
	"net/http"
	"strings"
	"user-crud/controllers"
	"user-crud/middlewares"

	"github.com/go-chi/chi"
)

func uploadRouter() chi.Router {

	u := chi.NewRouter()
	u.Use(middlewares.Authorization)

	uploadController := controllers.NewUploadController()
	u.Post("/", uploadController.Upload)

	return u
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Fatalf("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
