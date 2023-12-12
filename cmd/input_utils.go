package cmd

import (
	"avr-cli-backend/config"
	"avr-cli-backend/util"
	"fmt"
	"github.com/urfave/cli/v2"
)

type InputUtils struct {
	Validator string `json:"validator_address"`
	Name      string `json:"name"`
	Website   string `json:"website_url"`
	Logo      string `json:"logo_url"`
	Desc      string `json:"description"`
}

// commands to call indexer block, and staker
func InputUtilsCmd(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var inputUtilsCmds []*cli.Command

	inputUtilsCmd := &cli.Command{
		Name:  "input-field",
		Usage: "Generate numerical value of an input",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "value",
				Usage:    "The value of the input",
				Aliases:  []string{"v"},
				Value:    "",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {

			value := context.String("value")

			stringValue := util.BigIntStringToUtf8String(value)

			inputValue := fmt.Sprintf("Value: %v", stringValue)
			fmt.Println(inputValue)
			return nil

		},
	}
	outputUtilsCmd := &cli.Command{
		Name:  "output-field",
		Usage: "Generate string value from an input field",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "value",
				Usage:    "The value of the input",
				Aliases:  []string{"v"},
				Value:    "",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {

			value := context.String("value")

			stringValue := util.Utf8StringToBigInt(value)

			outputValue := fmt.Sprintf("Value: %vfield", stringValue)
			fmt.Println(outputValue)
			return nil

		},
	}
	inputUtilsCmds = append(inputUtilsCmds, inputUtilsCmd)
	inputUtilsCmds = append(inputUtilsCmds, outputUtilsCmd)

	return inputUtilsCmds
}
