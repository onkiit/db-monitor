package mongo

import (
	"context"

	mgo "github.com/onkiit/db-monitor/lib/db/mongo"
	"github.com/onkiit/dbinfo"
	mg "github.com/onkiit/dbinfo/mongo"
)

type mongo struct{}

func (m mongo) GetVersion(ctx context.Context) (*dbinfo.DBVersion, error) {
	session := mgo.Session()

	con := &dbinfo.Conn{
		Session: session,
	}
	store := mg.New(con)
	res, err := store.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m mongo) GetActiveClient(ctx context.Context) (*dbinfo.DBActiveClient, error) {
	session := mgo.Session()

	con := &dbinfo.Conn{
		Session: session,
	}
	store := mg.New(con)
	res, err := store.GetActiveClient(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m mongo) GetHealth(ctx context.Context) (*dbinfo.DBHealth, error) {
	session := mgo.Session()

	con := &dbinfo.Conn{
		Session: session,
	}
	store := mg.New(con)
	res, err := store.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func New() dbinfo.Store {
	return mongo{}
}
