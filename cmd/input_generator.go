package cmd

import (
	"avr-cli-backend/config"
	"avr-cli-backend/util"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
)

type InputGenerator struct {
	Validator string `json:"validator_address"`
	Name      string `json:"name"`
	Website   string `json:"website_url"`
	Logo      string `json:"logo_url"`
	Desc      string `json:"description"`
}

// commands to call indexer block, and staker
func InputGeneratorCmd(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var inputGeneratorCmds []*cli.Command

	inputGeneratorCmd := &cli.Command{
		Name:  "input-generator",
		Usage: "Generate input for the validator registry",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "validator",
				Usage:    "The validator to update the registry of",
				Aliases:  []string{"v"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "name",
				Usage:    "The name of the validator",
				Aliases:  []string{"n"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "website-url",
				Usage:    "The website url of the validator",
				Aliases:  []string{"w"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "logo-url",
				Usage:    "The logo url of the validator",
				Aliases:  []string{"l"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "description",
				Usage:    "The description of the validator",
				Aliases:  []string{"d"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "private-key",
				Usage:    "The private key of the validator",
				Aliases:  []string{"p"},
				Value:    "",
				Required: false,
			},
		},
		Action: func(context *cli.Context) error {

			// prepare the command
			var validator = context.String("validator")
			var name = context.String("name")
			var websiteUrl = context.String("website-url")
			var logoUrl = context.String("logo-url")
			var description = context.String("description")

			// convert the name, website url, logo url, and description to number
			// string to rune
			var nameNumber = util.Utf8StringToBigInt(name)
			var websiteUrlNumber = util.Utf8StringToBigInt(websiteUrl)
			var logoUrlNumber = util.Utf8StringToBigInt(logoUrl)
			var descriptionNumber = util.Utf8StringToBigInt(description)

			// add the word "field" to the name, website url, logo url, and description
			nameField := fmt.Sprintf("%dfield", nameNumber)
			websiteUrlField := fmt.Sprintf("%dfield", websiteUrlNumber)
			logoUrlField := fmt.Sprintf("%dfield", logoUrlNumber)
			descriptionField := fmt.Sprintf("%dfield", descriptionNumber)

			// generate the input
			inputGen := InputGenerator{
				Validator: validator,
				Name:      nameField,
				Website:   websiteUrlField,
				Logo:      logoUrlField,
				Desc:      descriptionField,
			}

			// generate the input
			inputGenJson, err := json.Marshal(inputGen)
			if err != nil {
				return err
			}
			// fmt the json
			fmt.Println(string(inputGenJson))

			// decode back

			return nil

		},
	}
	inputGeneratorCmds = append(inputGeneratorCmds, inputGeneratorCmd)

	return inputGeneratorCmds
}
