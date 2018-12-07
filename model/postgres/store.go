package postgres

import (
	"context"

	"github.com/onkiit/db-monitor/lib/db/psql"
	"github.com/onkiit/dbinfo"
	pq "github.com/onkiit/dbinfo/postgres"
)

type postgres struct {
}

func (p postgres) GetVersion(ctx context.Context) (*dbinfo.DBVersion, error) {
	db := psql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}
	store := pq.New(con)
	res, err := store.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p postgres) GetActiveClient(ctx context.Context) (*dbinfo.DBActiveClient, error) {
	db := psql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}

	store := pq.New(con)
	res, err := store.GetActiveClient(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p postgres) GetHealth(ctx context.Context) (*dbinfo.DBHealth, error) {
	db := psql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}
	store := pq.New(con)
	res, err := store.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func New() dbinfo.Store {
	return postgres{}
}
