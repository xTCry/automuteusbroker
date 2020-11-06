package main

import (
	"log"
	"os"

	"github.com/denverquane/automuteusbroker/broker"
	"github.com/joho/godotenv"
)

const DefaultPort = "8123"

func main() {
	godotenv.Load("config.txt")

	redisAddr := os.Getenv("REDIS_ADDRESS")
	if redisAddr == "" {
		log.Fatal("No REDIS_ADDRESS specified. Exiting.")
	}

	port := os.Getenv("BROKER_PORT")
	if port == "" {
		log.Println("No BROKER_PORT provided. Defaulting to " + DefaultPort)
		port = DefaultPort
	}
	redisUser := os.Getenv("REDIS_USER")
	redisPass := os.Getenv("REDIS_PASS")
	if redisUser != "" {
		log.Println("Using REDIS_USER=" + redisUser)
	} else {
		log.Println("No REDIS_USER specified.")
	}

	if redisPass != "" {
		log.Println("Using REDIS_PASS=<redacted>")
	} else {
		log.Println("No REDIS_PASS specified.")
	}

	msgBroker := broker.NewBroker(redisAddr, redisUser, redisPass)

	msgBroker.Start(port)
}
