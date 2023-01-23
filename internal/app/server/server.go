package server

import (
	"OzonTestTask/internal/app/data"
	"OzonTestTask/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
	Data   data.IData
}

type request struct {
	URL string `json:"URL"`
}

func NewServer(data data.IData) *Server {
	s := &Server{
		router: mux.NewRouter(),
		Data:   data,
	}

	s.ConfigureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/createShortUrl", s.handleShortLinkCreate()).Methods("POST")
	s.router.HandleFunc("/{shortURL}", s.handleGetOriginURL()).Methods("GET")
}

func (s *Server) handleShortLinkCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		link := &model.Link{
			OriginUrl: req.URL,
			ShortUrl: "",
		}

		if err := s.Data.Link().Create(link); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusCreated, link.ShortUrl)
	}
}

func (s *Server) handleGetOriginURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		link := &model.Link{
			OriginUrl: "",
			ShortUrl: vars["shortURL"],
		}

		if err := s.Data.Link().FindByShortURL(link); err != nil {
			s.error(w, r, http.StatusNoContent, err)
		}

		s.respond(w, r, http.StatusFound, link.OriginUrl)
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

func (s *Server) GetServerData() *data.IData {
	return &s.Data
}


