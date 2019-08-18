package apiserver

import (
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.ConfirureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting API server ...")

	return nil
}

// ConfirureLogger ...
func (s *APIServer) ConfirureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
