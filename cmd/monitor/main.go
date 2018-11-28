package main

import (
	"log"
	"net/http"

	"github.com/onkiit/db-monitor/lib/db/psql"
	"github.com/onkiit/db-monitor/lib/db/redis"
)

func openConnection() {
	if err := psql.Open("postgresql://postgres:postgres@localhost/Coba"); err != nil {
		log.Println(err)
		panic(err)
	}

	if err := redis.Connect("localhost:6379"); err != nil {
		log.Println(err)
		panic(err)
	}
}

func closeConnection() {
	if err := psql.Close(); err != nil {
		log.Println(err)
		panic(err)
	}

	if err := redis.Close(); err != nil {
		log.Println(err)
		panic(err)
	}
}

func run() {
	log.Println("Server starting at port 8081")
	http.ListenAndServe(":8081", logger())
}

func main() {
	openConnection()
	run()
	closeConnection()
}

func logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
	})
}
