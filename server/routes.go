package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/tejashwikalptaru/tutorial/handler"
	"net/http"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Get("/welcome", handler.Greet)
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
