package redis

import (
	"context"

	"github.com/onkiit/dbinfo"

	"github.com/onkiit/db-monitor/lib/db/redis"
	rd "github.com/onkiit/dbinfo/redis"
)

type rediss struct{}

func (r rediss) GetVersion(ctx context.Context) (*dbinfo.DBVersion, error) {
	redisCon := redis.Dial().Get()
	defer redisCon.Close()

	con := &dbinfo.Conn{
		Con: redisCon,
	}
	store := rd.New(con)
	res, err := store.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r rediss) GetActiveClient(ctx context.Context) (*dbinfo.DBActiveClient, error) {
	redisCon := redis.Dial().Get()
	con := &dbinfo.Conn{
		Con: redisCon,
	}

	store := rd.New(con)
	res, err := store.GetActiveClient(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r rediss) GetHealth(ctx context.Context) (*dbinfo.DBHealth, error) {
	redisCon := redis.Dial().Get()
	defer redisCon.Close()

	con := &dbinfo.Conn{
		Con: redisCon,
	}
	store := rd.New(con)
	res, err := store.GetHealth(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func New() dbinfo.Store {
	return rediss{}
}
