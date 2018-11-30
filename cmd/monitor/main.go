package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/onkiit/db-monitor/config"

	"github.com/spf13/viper"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/onkiit/db-monitor/controller/mongo"
	_ "github.com/onkiit/db-monitor/controller/postgres"
	_ "github.com/onkiit/db-monitor/controller/redis"
	"github.com/onkiit/db-monitor/lib/db/mongo"
	"github.com/onkiit/db-monitor/lib/db/psql"
	"github.com/onkiit/db-monitor/lib/db/redis"
	"github.com/onkiit/db-monitor/registry"
)

func openConnection() {
	if err := psql.Open(config.C().Server.Databases.Postgres.URI); err != nil {
		log.Println(err)
		panic(err)
	}

	if err := redis.Connect(config.C().Server.Databases.Redis.URI); err != nil {
		log.Println(err)
		panic(err)
	}

	if err := mongo.Open(config.C().Server.Databases.Mongo.URI); err != nil {
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
	log.Println("Database connection closed.")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../config")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		panic(err)
	}

	c := new(config.Config)
	if err := viper.Unmarshal(&c); err != nil {
		log.Println(err)
		panic(err)
	}

	config.C(c)
}

func run() {
	r := mux.NewRouter()

	//enable cors
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"})
	origins := handlers.AllowedOrigins(config.C().Client.AllowedCors)
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//register router
	for _, router := range registry.Router() {
		router.RegisterRoute(r)
	}

	serve := &http.Server{
		Addr:    config.C().Server.Host,
		Handler: handlers.CORS(origins, headers, methods)(r),
	}

	go func() {
		if err := serve.ListenAndServe(); err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Server starting at port 8180")

	stop := make(chan os.Signal)
	//set notification to stop if Interrupt signal (syscall,SIGINT) received
	signal.Notify(stop, syscall.SIGINT, syscall.SIGSTOP)
	<-stop

	closeConnection()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := serve.Shutdown(ctx); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Server closed")
}

func main() {
	initConfig()
	openConnection()
	run()
}
