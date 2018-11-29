package registry

import (
	"github.com/onkiit/db-monitor/api"
)

var r []api.Router

func RegisterRouter(route api.Router) {
	r = append(r, route)
}

func Router() []api.Router {
	return r
}
