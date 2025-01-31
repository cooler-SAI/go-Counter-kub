package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Counter-kub/handlers"
	"net/http"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Loading Counter...")
	http.HandleFunc("/", handlers.CounterHandler)
	log.Info().Msg("Server running on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}

}
