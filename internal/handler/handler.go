package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muhlemmer/kudofine/internal/config"
	"github.com/muhlemmer/kudofine/internal/model"
	"github.com/muhlemmer/kudofine/internal/ui"
)

type Server struct {
	config config.Config
	views  ui.Views
	router *chi.Mux
}

func NewServer(cfg config.Config) *Server {
	s := &Server{
		config: cfg,
		views:  ui.NewViews(cfg),
		router: chi.NewRouter(),
	}
	s.register(http.MethodGet, "/hello", newHandlerFunc(s.hello, s.views.Hello))
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) register(method, pattern string, h http.HandlerFunc) {
	slog.Info("registering handler", "method", method, "host", s.config.ExternalDomain, "pattern", pattern)
	s.router.Method(method, pattern, h)
}

func (s Server) hello(*http.Request) (response[model.HelloResponse], error) {
	// Simulate some processing
	return response[model.HelloResponse]{
		Data: model.HelloResponse{
			Name: "World",
		},
		Status: http.StatusOK,
	}, nil
}

type response[T any] struct {
	Data   T
	Status int
}

type contoller[T any] func(r *http.Request) (response[T], error)

type renderer[T any] func(ctx context.Context, w http.ResponseWriter, data T) error

func newHandlerFunc[T any](call contoller[T], render renderer[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := call(r)
		if err != nil {
			panic(err) // TODO: handle error
		}
		if resp.Status != 0 {
			w.WriteHeader(resp.Status)
		}
		err = render(r.Context(), w, resp.Data)
		if err != nil {
			panic(err) // TODO: handle error
		}
	}
}
