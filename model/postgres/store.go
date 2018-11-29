package postgres

import (
	"fmt"

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
	con := psql.DB()
	var c api.DBActiveClient
	err := con.QueryRow("SELECT count(0) as active_client FROM pg_stat_activity where state='active' ").Scan(&c.ActiveClient)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &c, nil
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
