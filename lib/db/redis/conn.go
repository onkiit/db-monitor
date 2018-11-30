package redis

import (
	"errors"
	"log"

	"github.com/gomodule/redigo/redis"
)

var conn *redis.Pool

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

	conn = pool
	log.Println("redis connected")

	return nil
}

func Close() error {
	if err := conn.Close(); err != nil {
		return err
	}
	conn = nil
	return nil
}

func Dial() *redis.Pool {
	return conn
}
