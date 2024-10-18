package app

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func Run() {

	var err error
	if err = prepareTemplates(); err != nil {
		log.Fatal().Msgf("could not parse html templates, reason -> %s", err.Error())
	}

	httpPort := os.Getenv("_GO_HTTP")
	if httpPort == "" {
		httpPort = DefaultHttpPort
	}
	bindHttpHandlers()
	log.Info().Msgf("creating HTTP server on '%s'", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatal().Msgf("creating HTTP server on port '%s' FAILED, reason -> %s", httpPort, err.Error())
	}
}
