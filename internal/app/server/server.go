package server

import (
	"OzonTestTask/internal/app/data"
	"OzonTestTask/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	data   data.Data
}

type request struct {
	URL string `json:"URL"`
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
		rq := &request{}
		if err := json.NewDecoder(r.Body).Decode(rq); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		l := &model.Link{
			OriginUrl: rq.URL,
			ShortUrl: "",
		}

		if err := s.data.Link().Create(l); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, l.ShortUrl)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string {"error" :err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}


