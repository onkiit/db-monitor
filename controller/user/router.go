package user

import (
	"github.com/gorilla/mux"
	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/model/user"
	"github.com/onkiit/db-monitor/registry"
)

func (c *Controller) RegisterRoute(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter()

	u.HandleFunc("/register", c.Register).Methods("POST")
	u.HandleFunc("/login", c.Login).Methods("POST")
	u.HandleFunc("/checkuser", c.CheckAvailable).Methods("GET")

}

func New() api.Router {
	return &Controller{user.New()}
}

func init() {
	registry.RegisterRouter("user", New())
}
