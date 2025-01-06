package server

import (
	"net/http"
	"time"

	// "github.com/dlsu-lscs/lscs-central-auth-api/internal/middlewares"
	"github.com/dlsu-lscs/lscs-central-auth-api/internal/handlers"
	"github.com/dlsu-lscs/lscs-central-auth-api/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(httprate.LimitByIP(100, time.Minute))

	registerAuthRoutes(r)
	r.Mount("/", registerAdminRoutes())
	return r
}

/* Auth Routes */
func registerAuthRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test test test"))
	})
	r.Post("/request-key", handlers.RequestAPIKey) // `/request_key?` TODO: change to POST if need condition before able to request, ex. need to be admin email only
}

func registerAdminRoutes() chi.Router {
	/* Protected Routes */
	r := chi.NewRouter()
	r.Use(middlewares.AdminMiddleware)

	r.Get("/members", handlers.GetAllMembersHandler)
	r.Post("/member", handlers.GetMemberInfo)
	r.Post("/check-email", handlers.CheckEmailHandler)
	r.Post("/refresh-token", handlers.RefreshTokenHandler)
	r.Get("/protected-test", handlers.GetAllMembersHandler)
	return r
}
