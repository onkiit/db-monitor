package redis

import (
	"github.com/gorilla/mux"
	"github.com/onkiit/db-monitor/api"
	rd "github.com/onkiit/db-monitor/model/redis"
	"github.com/onkiit/db-monitor/registry"
)

func (c *Controller) RegisterRoute(r *mux.Router) {
	s := r.PathPrefix("/redis").Subrouter()
	s.HandleFunc("/version", c.GetVersion)
	s.HandleFunc("/client", c.GetActiveClient)
	s.HandleFunc("/health", c.GetHealth)
}

func New() api.Router {
	return &Controller{rd.New()}
}

func init() {
	registry.RegisterRouter(New())
}
