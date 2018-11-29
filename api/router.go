package api

import (
	"github.com/gorilla/mux"
)

type Router interface {
	RegisterRoute(r *mux.Router)
}
