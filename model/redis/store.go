package redis

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/onkiit/db-monitor/api"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/onkiit/db-monitor/lib/db/redis"
)

type rediss struct{}

func getString(info string, prefix string) (string, error) {
	var text string
	reader := bufio.NewReader(bytes.NewBufferString(info))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, prefix) {
			text = line
			break
		}
	}
	return text, nil
}

func getValue(str string) string {
	split := strings.Split(str, ":")
	return split[1]
}

func (r rediss) GetVersion() (*api.DBVersion, error) {
	con := redis.Dial().Get()
	info, err := redigo.String(con.Do("INFO", "SERVER"))
	if err != nil {
		return nil, err
	}

	defer con.Close()

	strVersion, err := getString(info, "redis_version")
	if err != nil {
		return nil, err
	}

	strOs, err := getString(info, "os")
	if err != nil {
		return nil, err
	}

	strGcc, err := getString(info, "gcc_version")
	if err != nil {
		return nil, err
	}

	v := getValue(strVersion)
	os := getValue(strOs)
	gcc := getValue(strGcc)

	version := fmt.Sprintf("Redis version: %s OS %s gcc_version %s \n", v, os, gcc)

	res := &api.DBVersion{
		Version: version,
	}

	return res, nil
}

func (r rediss) GetActiveClient() (*api.DBActiveClient, error) {
	con := redis.Dial().Get()
	info, err := redigo.String(con.Do("INFO", "CLIENTS"))
	if err != nil {
		return nil, err
	}

	defer con.Close()

	str, err := getString(info, "connected_clients")
	if err != nil {
		return nil, err
	}

	client, err := strconv.ParseInt(getValue(str), 10, 64)
	if err != nil {
		return nil, err
	}

	res := &api.DBActiveClient{
		ActiveClient: int(client),
	}
	return res, nil
}

func getUsage(info string) (string, error) {
	str, err := getString(info, "used_memory")
	if err != nil {
		return "", err
	}

	usage := getValue(str)
	return usage, nil
}

func getKeys(info string) ([]string, error) {
	strExpired, err := getString(info, "expired_keys")
	if err != nil {
		return nil, err
	}

	strEvicted, err := getString(info, "evicted_keys")
	if err != nil {
		return nil, err
	}

	exp := getValue(strExpired)
	evi := getValue(strEvicted)
	return []string{exp, evi}, nil
}

func (r rediss) getSlowlogCount(con redigo.Conn) (int, error) {
	count, err := redigo.Int(con.Do("SLOWLOG", "LEN"))
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r rediss) getMemoryStats(con redigo.Conn) ([]interface{}, error) {
	stats, err := redigo.Values(con.Do("MEMORY", "STATS"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return stats, nil
}

func (r rediss) GetHealth() (*api.DBHealth, error) {
	con := redis.Dial().Get()

	defer con.Close()

	size, err := redigo.Int(con.Do("DBSIZE"))
	if err != nil {
		return nil, err
	}

	info, err := redigo.String(con.Do("INFO"))
	if err != nil {
		return nil, err
	}
	usage, err := getUsage(info)
	if err != nil {
		return nil, err
	}

	keys, err := getKeys(info)
	if err != nil {
		return nil, err
	}

	countLog, err := r.getSlowlogCount(con)
	if err != nil {
		return nil, err
	}

	stats, err := r.getMemoryStats(con)
	if err != nil {
		return nil, err
	}

	res := &api.DBHealth{
		RedisHealth: api.RedisHealth{
			AvailableKey: size,
			ExpiredKeys:  keys[0],
			EvictedKeys:  keys[1],
			MemoryUsage:  usage,
			SlowlogCount: countLog,
			MemoryStats: api.MemoryStats{
				PeakAllocated:    stats[1].(int64),
				TotalAllowed:     stats[3].(int64),
				StartupAllocated: stats[5].(int64),
				PeakPercentage:   0,
				Fragmentation:    0,
			},
		},
	}
	return res, nil
}

func New() api.Store {
	return rediss{}
}
