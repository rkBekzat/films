package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rkBekzat/films/internal/handler"
	"github.com/rkBekzat/films/internal/repository"
	"github.com/rkBekzat/films/internal/service"
	"github.com/spf13/viper"
)

//	@title			Film API
//	@version		1.0
//	@description	WEB API

//	@host		localhost:8080
//	@basePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}
	log.Println("created client db")

	repo := repository.NewRepo(db)
	service := service.NewService(repo)
	handle := handler.NewHandler(service)

	servMx := http.NewServeMux()

	handle.RegisterRoute(servMx)
	log.Println("registered routes")

	s := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: servMx,
	}

	log.Println("server running...")
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("Server ended with: ", err.Error())
		return
	}
	log.Println("server ends...")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
