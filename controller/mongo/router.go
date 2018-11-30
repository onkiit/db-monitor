package mongo

import (
	"github.com/gorilla/mux"
	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/model/mongo"
	"github.com/onkiit/db-monitor/registry"
)

func (c *Controller) RegisterRoute(r *mux.Router) {
	m := r.PathPrefix("/mongo").Subrouter()
	m.HandleFunc("/version", c.GetVersion)
	m.HandleFunc("/client", c.GetActiveClient)
}

func New() api.Router {
	return &Controller{mongo.New()}
}

func init() {
	registry.RegisterRouter(New())
}
