package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/onkiit/db-monitor/api"
	"github.com/onkiit/db-monitor/lib/db/psql"
)

type postgres struct {
}

func (p postgres) GetVersion(ctx context.Context) (*api.DBVersion, error) {
	con := psql.DB()

	var v api.DBVersion
	err := con.QueryRowContext(ctx, "SELECT VERSION()").Scan(&v.Version)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (p postgres) GetActiveClient(ctx context.Context) (*api.DBActiveClient, error) {
	con := psql.DB()
	var c api.DBActiveClient
	err := con.QueryRowContext(ctx, "SELECT count(0) as active_client FROM pg_stat_activity where state='active' ").Scan(&c.ActiveClient)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &c, nil
}

func (p postgres) getTables(ctx context.Context) (map[string][]string, error) {
	rows, err := psql.DB().QueryContext(ctx, "select schemaname as schema, relname as table from pg_statio_all_tables where schemaname not in ('pg_catalog', 'pg_toast', 'information_schema')")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tables := make(map[string][]string)
	for rows.Next() {
		var schema, table string
		err := rows.Scan(&schema, &table)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tables[schema] = append(tables[schema], table)
	}

	if len(tables) < 1 {
		return nil, errors.New("Could not find table")
	}

	return tables, nil
}

func (p postgres) getTableSize(ctx context.Context) ([]api.TableInformation, error) {
	tables, err := p.getTables(ctx)
	if err != nil {
		return nil, err
	}

	var res []api.TableInformation
	for k, v := range tables {
		var info api.TableInformation
		var tableSize, indexSize string
		if len(v) < 1 {
			return nil, errors.New("Schema has no table")
		}

		for _, val := range v {
			qTable := fmt.Sprintf("SELECT (pg_total_relation_size('%s.%s')) as tableSize", k, val)
			qIndex := fmt.Sprintf("SELECT (pg_indexes_size('%s.%s')) as indexSize", k, val)
			err := psql.DB().QueryRow(qTable).Scan(&tableSize)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			err = psql.DB().QueryRow(qIndex).Scan(&indexSize)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			info = api.TableInformation{
				SchemaName: k,
				TableName:  val,
				TableSize:  tableSize,
				IndexSize:  indexSize,
			}

			res = append(res, info)
		}
	}
	return res, nil
}

func (p postgres) GetHealth(ctx context.Context) (*api.DBHealth, error) {
	con := psql.DB()

	rows, err := con.QueryContext(ctx, "select datname, pg_database_size(datname) as size from pg_database where datname = (select current_database()) order by pg_database_size(datname) desc;")
	if err != nil {
		return nil, err
	}

	var datname, size string
	for rows.Next() {
		err := rows.Scan(&datname, &size)
		if err != nil {
			return nil, err
		}
	}

	tableInfo, err := p.getTableSize(ctx)
	if err != nil {
		return nil, err
	}

	res := &api.DBHealth{
		PsqlHealth: api.PsqlHealth{
			DBInformation: api.DBInformation{
				DBName: datname,
				DBSize: size,
			},
			TableInformation: tableInfo,
		},
	}

	return res, nil
}

func New() api.Store {
	return postgres{}
}
