package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"path"
	"time"

	"connectrpc.com/grpcreflect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/muhlemmer/kudofine/gen/submission/v1/submissionv1connect"
	submissionv1 "github.com/muhlemmer/kudofine/internal/api/submission/submissionv1_impl"
	"github.com/muhlemmer/kudofine/internal/data"
	"github.com/muhlemmer/kudofine/internal/version"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Config struct {
	AppName    string
	ListenAddr string
	Timeout    time.Duration
}

// Defaults
const (
	DefaultAppName       = "kudofine api"
	DefaultListenAddress = "0.0.0.0:8212"
	DefaultTimeout       = 10 * time.Second
)

var DefaultConfig = Config{
	AppName:    DefaultAppName,
	ListenAddr: DefaultListenAddress,
	Timeout:    DefaultTimeout,
}

type ApiServer struct {
	config Config
	router *chi.Mux
	server *http.Server
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *ApiServer) RegisterConnectService(services ...func() (path string, h http.Handler)) {
	for _, service := range services {
		route, handler := service()
		route = path.Join(route, "{method}")
		s.router.Method(http.MethodPost, route, handler)
	}
}

func NewApiServer(background context.Context, config Config, data *data.Data) error {
	router := chi.NewRouter()

	s := &ApiServer{
		config: config,
		router: router,
		server: &http.Server{
			Addr: config.ListenAddr,
			Handler: h2c.NewHandler(router, &http2.Server{
				IdleTimeout: time.Minute,
			}),
			ReadHeaderTimeout: time.Second,
			IdleTimeout:       time.Minute,
			// ErrorLog:          slog.NewLogLogger(),
		},
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(httplog.RequestLogger(s.logger()))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(config.Timeout))

	s.RegisterConnectService(submissionv1.NewSubmissionService(background, data))
	s.RegisterConnectService(reflectService()...)

	chi.Walk(router, func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		slog.InfoContext(background, "registered route", "method", method, "route", route)
		return nil
	})

	return s.listenAndServe(background)
}

func (s *ApiServer) listenAndServe(background context.Context) error {
	serverErr := make(chan error, 1)
	go func(ctx context.Context, serverErr chan<- error) {
		err := s.server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			slog.InfoContext(ctx, err.Error())
			err = nil
		}
		serverErr <- err
		close(serverErr)
	}(background, serverErr)

	slog.InfoContext(background, "server listening", "addr", s.config.ListenAddr)

	// When the background context is done, shut down the server
	go func(background context.Context) {
		<-background.Done()
		slog.InfoContext(background, "shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), s.config.Timeout)
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			slog.ErrorContext(background, "error shutting down server", "err", err)
		}
	}(background)

	return <-serverErr
}

func (s *ApiServer) logger() *httplog.Logger {
	return httplog.NewLogger(s.config.AppName, httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelInfo,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": version.Version(),
			"env":     version.Environment(),
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})
}

func reflectService() []func() (string, http.Handler) {
	reflector := grpcreflect.NewStaticReflector(
		submissionv1connect.SubmissionServiceName,
	)
	return []func() (string, http.Handler){
		func() (string, http.Handler) {
			return grpcreflect.NewHandlerV1(reflector)
		},
		func() (string, http.Handler) {
			return grpcreflect.NewHandlerV1Alpha(reflector)
		},
	}
}
