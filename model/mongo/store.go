package mongo

import (
	"context"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/onkiit/db-monitor/api"
	mgo "github.com/onkiit/db-monitor/lib/db/mongo"
)

type mongo struct{}

func (m mongo) GetVersion(ctx context.Context) (*api.DBVersion, error) {
	session := mgo.Session()

	info, err := session.BuildInfo()
	if err != nil {
		return nil, err
	}

	res := &api.DBVersion{
		Version: fmt.Sprintf("Mongo version %s", info.Version),
	}

	return res, nil
}

func (m mongo) GetActiveClient(ctx context.Context) (*api.DBActiveClient, error) {
	session := mgo.Session()

	var b bson.M
	if err := session.DB("test").Run("serverStatus", &b); err != nil {
		return nil, err
	}

	client := b["globalLock"].(bson.M)["activeClients"].(bson.M)["total"]

	res := &api.DBActiveClient{
		ActiveClient: client.(int),
	}
	return res, nil
}

func (m mongo) GetHealth(ctx context.Context) (*api.DBHealth, error) {
	session := mgo.Session()
	r := bson.M{}
	if err := session.DB("test").Run("dbstats", &r); err != nil {
		return nil, err
	}

	res := &api.DBHealth{
		MongoHealth: api.MongoHealth{
			DBName:              r["db"].(string),
			AvailableCollection: r["collections"].(int),
			StorageSize:         r["storageSize"].(float64),
			Indexes:             r["indexes"].(int),
			DataSize:            r["dataSize"].(float64),
		},
	}
	return res, nil
}

func New() api.Store {
	return mongo{}
}
