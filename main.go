//go:generate go run frontend/generate.go

package main

import (
	"github.com/rs/zerolog/log"

	"github.com/clintjedwards/scheduler/cmd"
	"github.com/clintjedwards/scheduler/config"
)

func main() {
	conf, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not load env config")
	}

	setupLogging(conf.LogLevel, conf.Debug)

	cmd.RootCmd.Execute()
}
