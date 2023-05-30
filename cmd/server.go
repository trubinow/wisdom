package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
	"wisdom/internal/handlers"
	"wisdom/internal/servers"
	"wisdom/internal/services"
)

func main() {
	log := logrus.WithFields(logrus.Fields{
		"service": "tcp-server",
	})

	// Convert PORT env variable value to an integer
	port, err := strconv.Atoi(os.Getenv("TCP_SERVER_PORT"))
	if err != nil {
		log.Errorf("error converting PORT variable to integer: %s", err.Error())
		return
	}

	// Convert DIFFICULTY env variable value to an integer
	difficulty, err := strconv.Atoi(os.Getenv("DIFFICULTY"))
	if err != nil {
		log.Errorf("error converting DIFFICULTY variable to integer: %s", err.Error())
		return
	}

	// Convert DIFFICULTY env variable value to an integer
	tcpConnectionDeadline, err := time.ParseDuration(os.Getenv("TCP_CONNECTION_DEADLINE"))
	if err != nil {
		log.Errorf("error converting TCP_CONNECTION_DEADLINE variable to time.Duration: %s", err.Error())
		return
	}

	// Create quotation service
	quotationService := services.NewQuotationService()

	// Load quotations from a text file
	fileName := os.Getenv("QUOTATION_FILE")
	err = quotationService.Load(fileName)
	if err != nil {
		log.Fatalf("failed to load quotations: %s", err.Error())
	}

	// Create tcp connection handler
	handler := handlers.NewWisdomHandler(log, quotationService, difficulty)

	// Run tcp server
	tcpServer := servers.NewTCPServer(handler, tcpConnectionDeadline)
	err = tcpServer.Run(port)
	if err != nil {
		log.Fatalf("failed to run tcp server: %s", err.Error())
	}
}
