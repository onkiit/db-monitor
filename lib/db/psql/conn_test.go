package psql

import (
	"testing"

	"github.com/onkiit/db-monitor/config"
)

func TestOpen(t *testing.T) {
	err := Open(config.C().Server.Databases.Postgres.URI)
	if err != nil {
		t.Error(err)
	}
}
