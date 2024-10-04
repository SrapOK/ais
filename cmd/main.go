package main

import (
	"flag"
	"kis/internal/config"
	"kis/internal/contracts"
	"kis/internal/repository"
	"kis/internal/server"
	"kis/internal/service"
	"log"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "path to config")

	cfg := config.Load(*configPath)

	db, err := repository.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := contracts.NewHandler(services)

	server := server.NewServer(cfg.HttpServer, handlers.InitRoutes())

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}

	log.Print("server stopped")
}
