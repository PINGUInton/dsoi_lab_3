package main

import (
	"log"

	handler "ticket/handler"
	repo "ticket/repository"
	server "ticket/server"
	service "ticket/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found: %v", err)
	}

	db, err := repo.NewPostgresDB(&repo.Config{
		Host:     "postgres",
		Port:     "5432",
		Username: "program",
		Password: "test",
		DBName:   "tickets",
		SSLMode:  "disable",
	})

	if err != nil {
		return
	}

	repos := repo.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8070", handlers.InitRouters()); err != nil {
		return
	}
}
