package version

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

func (r rediss) GetHealth() (*api.DBHealth, error) {
	return nil, nil
}

func New() api.Store {
	return rediss{}
}
