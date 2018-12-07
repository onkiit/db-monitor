package registry

import (
	"github.com/onkiit/db-monitor/api"
)

var r = make(map[string]api.Router)

func RegisterRouter(name string, route api.Router) {
	r[name] = route
}

func Router(name string) api.Router {
	return r[name]
}
