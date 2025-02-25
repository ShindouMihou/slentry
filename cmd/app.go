package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"slentry/internal/http"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Stack()
	log.Logger = logger.Logger()
	http.Start()
}
