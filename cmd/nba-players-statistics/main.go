package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"nba-players-statistics/config"
	"nba-players-statistics/internal"
	"nba-players-statistics/server"
)

var logger = log.New(os.Stdout, "nba-player-statistics ", log.LstdFlags|log.Lshortfile)

func main() {
	ctx := context.Background()
	logger.Print("Initializing...")
	defer logger.Print("Bye Bye :)")

	dbConfig := config.NewDbConfig()
	dao, err := internal.NewPsqlStatisticsDao(ctx, dbConfig, logger)
	if err != nil {
		return
	}
	defer dao.Disconnect()
	service := internal.NewSimpleStatisticsService(dao, logger)
	controller := internal.NewStatisticsController(service, logger)
	timeout := 30 * time.Second
	port := "80" // os.Getenv("SERVER_PORT")

	server.Init(logger)
	server.AddHandler("/health", controller.Health, timeout, server.GET)
	server.AddHandler("/log", controller.Log, timeout, server.POST)
	server.AddHandler("/get-statistics", controller.GetStats, timeout, server.GET)
	server.Listen(port)

	logger.Print("--- Initialization Completed Successfully ---")

	gracefullyQuit()
}

func gracefullyQuit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
}
