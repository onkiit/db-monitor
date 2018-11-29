package version

import (
	"github.com/gorilla/mux"
	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/model/postgres"
	"github.com/onkiit/db-monitor/registry"
)

func (c *Controller) RegisterRoute(r *mux.Router) {
	v := r.PathPrefix("/postgres").Subrouter()
	v.HandleFunc("/version", c.GetPostgresVersion)
	v.HandleFunc("/client", c.GetActiveClient)
}

func New() api.Router {
	return &Controller{postgres.New()}
}

func init() {
	registry.RegisterRouter(New())
}
