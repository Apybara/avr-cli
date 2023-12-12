package cmd

import (
	"avr-cli-backend/config"
	"github.com/urfave/cli/v2"
)

// commands to call indexer block, and staker
func LookupApiCmd(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var lookupApiCmds []*cli.Command

	lookupApiCmd := &cli.Command{
		Name:        "lookup-api",
		Usage:       "Lookup API",
		Description: "This command will lookup API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Usage:   "The port of the API",
				Aliases: []string{"p"},
				Value:   "8080",
			},
			&cli.StringFlag{
				Name:    "host",
				Usage:   "The host of the API",
				Aliases: []string{"h"},
				Value:   "localhost",
			},
		},
		Action: func(context *cli.Context) error {
			return nil

		},
	}
	lookupApiCmds = append(lookupApiCmds, lookupApiCmd)

	return lookupApiCmds
}
