package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pitshifer/valera-acceptor/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// NewServer ...
func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/devices", s.handleDeviceCreate()).Methods("POST")
}

func (s *server) handleDeviceCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
