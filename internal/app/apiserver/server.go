package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pitshifer/valera-acceptor/internal/app/model"
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
	s.router.HandleFunc("/accept", s.handleAccept()).Methods("POST")
}

func (s *server) handleAccept() http.HandlerFunc {
	type request struct {
		MacAddress string               `json:"mac_address"`
		RSSI       int                  `json:"rssi"`
		Data       model.IndicationData `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var device *model.Device
		device, err := s.store.Device().FindByMacAddress(req.MacAddress)
		if err != nil {
			if err == store.ErrRecordNotFound {
				device = &model.Device{
					MacAddress: req.MacAddress,
				}
				if err := s.store.Device().Create(device); err != nil {
					s.error(w, r, http.StatusUnprocessableEntity, err)
					return
				}
				logrus.Infof("Insert new device into DB. Device mac address: %s", device.MacAddress)
			} else {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}

		s.store.HandleNewIndication(&model.Indication{
			DeviceID:  device.ID,
			CreatedAt: time.Now(),
			Data:      req.Data,
		})

		s.respond(w, r, http.StatusAccepted, device)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
