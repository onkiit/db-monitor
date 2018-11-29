package registry

import (
	"log"

	"github.com/onkiit/db-monitor/api"
)

var m = make(map[string]api.Store)

func RegisterModel(name string, store api.Store) {
	m[name] = store
}

func Models(name string) api.Store {
	log.Println("called", name)
	return m[name]
}
