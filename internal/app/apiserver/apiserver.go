package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer for API requests
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New API server config
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start API server
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	s.logger.Info("Starting api server on ", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Configure APIserver logger level from APIServer.config
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	// API routes
	baseRoute := s.router.PathPrefix("/api/v1").Subrouter()
	baseRoute.Handle("/pin", s.WithLogging(s.handlePin())).Methods("GET")

	// Redoc documentation
	s.router.Handle("/", http.RedirectHandler("/docs", 301))
	s.router.Handle("/docs", s.WithLogging(s.Redoc())).Methods("GET")
	s.router.Handle("/api/swagger.yml", s.WithLogging(http.FileServer(http.Dir("./"))))
}
