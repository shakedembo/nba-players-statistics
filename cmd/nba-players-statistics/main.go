package main

import (
	"context"
	"log"
	"nba-players-statistics/cmd/server"
	"nba-players-statistics/internal"
	"os"
	"os/signal"
	"time"
)

var logger = log.New(os.Stdout, "url-shortener ", log.LstdFlags|log.Lshortfile)

func main() {
	logger.Print("Initializing...")
	defer logger.Print("Bye Bye :)")

	service := internal.NewSimpleStatisticsService()
	controller := internal.NewStatisticsController(service)
	timeout := 30 * time.Second
	port := "80" // os.Getenv("SERVER_PORT")

	server.Init(logger)
	server.AddHandler("/health", controller.Health, timeout, server.GET)
	server.Listen(port)

}

func gracefullyQuit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
}
