package server

import (
	"OzonTestTask/internal/app/data"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	data   data.Data
}

func NewServer(data data.Data) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		data: data,
	}

	s.ConfigureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/saveUrl", s.handleShortLinkCreate()).Methods("POST")
}

func (s *Server) handleShortLinkCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}


