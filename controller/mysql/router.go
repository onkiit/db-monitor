package mysql

import (
	"github.com/gorilla/mux"
	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/model/mysql"
	"github.com/onkiit/db-monitor/registry"
)

func (c *Controller) RegisterRoute(r *mux.Router) {
	m := r.PathPrefix("/mysql").Subrouter()
	m.HandleFunc("/version", c.GetVersion)
	m.HandleFunc("/client", c.GetActiveClient)
	m.HandleFunc("/health", c.GetHealth)
}

func New() api.Router {
	return &Controller{mysql.New()}
}

func init() {
	registry.RegisterRouter("mysql", New())
}
