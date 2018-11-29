package mongo

import (
	"log"

	"github.com/globalsign/mgo"
)

var session *mgo.Session

func Open(host string) error {
	sess, err := mgo.Dial(host)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := sess.Ping(); err != nil {
		log.Println(err)
		return err
	}

	session = sess
	log.Println("mongo connected")

	return nil
}

func Close() {
	session.Close()
	session = nil
}

func Session() *mgo.Session {
	return session
}
