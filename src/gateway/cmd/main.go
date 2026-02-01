package main

import (
	"log"

	handler "gateway/handler"
	redis "gateway/rollback"
	worker "gateway/rollback/worker"
	server "gateway/server"
	services "gateway/services"
)

func main() {
	redis.InitRedis()
	worker.StartRetryWorker()

	services := services.NewServices()
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatal("Failed to start server: ", err)
		return
	}
}
