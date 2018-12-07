package mysql

import (
	"context"

	"github.com/onkiit/db-monitor/lib/db/mysql"
	"github.com/onkiit/dbinfo"
	mq "github.com/onkiit/dbinfo/mysql"
)

type msq struct{}

func (m msq) GetVersion(ctx context.Context) (*dbinfo.DBVersion, error) {
	db := mysql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}
	store := mq.New(con)
	res, err := store.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m msq) GetActiveClient(ctx context.Context) (*dbinfo.DBActiveClient, error) {
	db := mysql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}

	store := mq.New(con)
	res, err := store.GetActiveClient(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m msq) GetHealth(ctx context.Context) (*dbinfo.DBHealth, error) {
	db := mysql.DB()
	con := &dbinfo.Conn{
		DB: db,
	}
	store := mq.New(con)
	res, err := store.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func New() dbinfo.Store {
	return msq{}
}
