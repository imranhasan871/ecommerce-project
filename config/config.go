package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variable :", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
