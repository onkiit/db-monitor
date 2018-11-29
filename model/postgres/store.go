package postgres

import (
	"log"

	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/lib/db/psql"
	"github.com/onkiit/db-monitor/registry"
)

type postgres struct{}

func (p postgres) GetVersion() (*api.DBVersion, error) {
	con := psql.DB()

	var v api.DBVersion
	err := con.QueryRow("SELECT VERSION()").Scan(&v.Version)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (p postgres) GetActiveClient() (*api.DBActiveClient, error) {
	log.Println("From postgresql")
	return nil, nil
}

func (p postgres) GetHealth() (*api.DBHealth, error) {
	return nil, nil
}

func New() api.Store {
	return postgres{}
}

func init() {
	registry.RegisterModel("postgres", New())
}
