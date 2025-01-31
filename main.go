package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Counter-kub/handlers"
	"go-Counter-kub/tools"
	"net/http"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Loading Counter...")
	http.HandleFunc("/", handlers.CounterHandler)

	server := &http.Server{
		Addr: ":8081",
	}

	go func() {
		log.Info().Msg("Server is running on http://localhost:8081")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	tools.StopApp(server)

}
