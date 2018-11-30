package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/onkiit/db-monitor/controller/postgres"
	_ "github.com/onkiit/db-monitor/controller/redis"
	"github.com/onkiit/db-monitor/lib/db/psql"
	"github.com/onkiit/db-monitor/lib/db/redis"

	"github.com/onkiit/db-monitor/lib/db/mongo"
	"github.com/onkiit/db-monitor/registry"
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

	if err := mongo.Open("localhost:27017"); err != nil {
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

	mongo.Close()
}

func run() {
	r := mux.NewRouter()

	//enable cors
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"})
	origins := handlers.AllowedOrigins([]string{"http://127.0.0.1:8080", "http://localhost:8080"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//register router
	for _, router := range registry.Router() {
		router.RegisterRoute(r)
	}

	serve := &http.Server{
		Addr:    "127.0.0.1:8180",
		Handler: handlers.CORS(origins, headers, methods)(r),
	}

	log.Println("Server starting at port 8180")
	log.Fatal(serve.ListenAndServe())
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
