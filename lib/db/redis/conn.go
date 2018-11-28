package redis

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn

func Connect(host string) error {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}

	c := pool.Get()
	_, err := c.Do("PING")
	if err != nil {
		return errors.New("Unable to connect with redis.")
	}

	conn = c

	return nil
}

func Close() error {
	if err := conn.Close(); err != nil {
		return err
	}
	conn = nil
	return nil
}

func Dial() redis.Conn {
	return conn
}
