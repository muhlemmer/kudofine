package handler

import (
	"log/slog"
	"net/http"

	"github.com/muhlemmer/kudofine/internal/config"
	"github.com/muhlemmer/kudofine/internal/model"
	"github.com/muhlemmer/kudofine/internal/ui"
)

type Server struct {
	config config.Config
	views  ui.Views
	mux    *http.ServeMux
}

func NewServer(cfg config.Config) *Server {
	s := &Server{
		config: cfg,
		views:  ui.NewViews(cfg),
		mux:    http.NewServeMux(),
	}
	s.register(http.MethodGet, "hello", s.helloHandler())
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) register(method, path string, h http.HandlerFunc) {
	slog.Info("registering handler", "method", method, "host", s.config.ExternalDomain, "path", path)
	s.mux.HandleFunc(endpoint(http.MethodGet, s.config.ExternalDomain, path), h)
}

func (s Server) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.views.Hello(r.Context(), w, model.HelloResponse{
			Name: "World",
		})
	}
}
