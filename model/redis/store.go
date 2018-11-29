package version

import (
	"bufio"
	"bytes"
	"fmt"
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

func (r rediss) redisDoString(command string, args ...interface{}) (string, error) {
	con := redis.Dial()
	info, err := redigo.String(con.Do(command, args...))
	if err != nil {
		return "", err
	}
	return info, nil
}

func getValue(str string) string {
	split := strings.Split(str, ":")
	return split[1]
}

func (r rediss) GetVersion() (*api.DBVersion, error) {
	info, err := r.redisDoString("INFO", "SERVER")
	if err != nil {
		return nil, err
	}

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
	return nil, nil
}

func (r rediss) GetHealth() (*api.DBHealth, error) {
	return nil, nil
}

func New() api.Store {
	return rediss{}
}
