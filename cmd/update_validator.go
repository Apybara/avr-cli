package cmd

import (
	"avr-cli-backend/config"
	"avr-cli-backend/util"
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
)

// commands to call indexer block, and staker
func UpdateRegistryCmd(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var updateValidatorRegistryCmds []*cli.Command

	updateValidatorRegistryCmd := &cli.Command{
		Name:        "register-my-validator",
		Aliases:     []string{"update-registry"},
		Usage:       "Update the registry of a validator",
		Description: "This command will update the registry of a validator. Note that this command will only work if the validator is already registered. If the validator is not registered, use the register-validators command instead.",
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
				Name:     "website_url",
				Usage:    "The website url of the validator",
				Aliases:  []string{"w"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "logo_url",
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
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {

			// prepare the command
			var query = cfg.Common.AleoNodeUrl
			var broadcast = cfg.Common.AleoNodeUrl + "/testnet3/transaction/broadcast"
			var priorityFee = "0"
			var privateKey = context.String("private-key")
			var validator = context.String("validator")
			var name = context.String("name")
			var websiteUrl = context.String("website_url")
			var logoUrl = context.String("logo_url")
			var description = context.String("description")

			var nameNumber = util.Utf8StringToBigInt(name)
			var websiteUrlNumber = util.Utf8StringToBigInt(websiteUrl)
			var logoUrlNumber = util.Utf8StringToBigInt(logoUrl)
			var descriptionNumber = util.Utf8StringToBigInt(description)

			// add the word "field" to the name, website url, logo url, and description
			nameField := fmt.Sprintf("%dfield", nameNumber)
			websiteUrlField := fmt.Sprintf("%dfield", websiteUrlNumber)
			logoUrlField := fmt.Sprintf("%dfield", logoUrlNumber)
			descriptionField := fmt.Sprintf("%dfield", descriptionNumber)

			// field can only be 251 bits, throw an error if it is too long
			nameFieldBitLen := util.CalculateLength(nameField)
			if nameFieldBitLen > 76 {
				fmt.Println("Name field is too long. It must be less than 76 characters.", nameFieldBitLen, nameField)
				return nil
			}

			websiteUrlFieldBitLen := util.CalculateLength(websiteUrlField)
			if websiteUrlFieldBitLen > 76 {
				fmt.Println("Website url field is too long. It must be less than 76 characters.", websiteUrlFieldBitLen, websiteUrlField)
				return nil
			}

			logoUrlFieldBitLen := util.CalculateLength(logoUrlField)
			if logoUrlFieldBitLen > 76 {
				fmt.Println("Logo url field is too long. It must be less than 76 characters.", logoUrlFieldBitLen, logoUrlField)
				return nil
			}

			descriptionFieldBitLen := util.CalculateLength(descriptionField)
			if descriptionFieldBitLen > 76 {
				fmt.Println("Description field is too long. It must be less than 76 characters.", descriptionFieldBitLen, descriptionField)
				return nil
			}

			// put quotes around the entire json body
			var inputString = "" + "\"" + "{\"validator_address\":\"" + validator + "\",\"name\":" + nameField + ", \"website_url\":" + websiteUrlField + ", \"logo_url\":" + logoUrlField + ", \"description\":" + descriptionField + "}" + "\""
			var inputStringF = fmt.Sprintf("snarkos developer execute %s update_validator %s --private-key %s --query %s --broadcast %s --priority-fee %s", cfg.Common.ProgramID, inputString, privateKey, query, broadcast, priorityFee)

			cmd := exec.Command("bash", "-c", inputStringF)

			// command in bash

			fmt.Println("Executing command: ", cmd)

			// handle the output warnings
			cmd.Stdout = context.App.Writer
			cmd.Stderr = context.App.ErrWriter

			cmd.Run()

			return nil

		},
	}
	updateValidatorInfoAsAdminCmd := &cli.Command{
		Name:        "update-registry-as-admin",
		Usage:       "Update the registry of a validator as an admin",
		Description: "This command will update the registry of a validator as an admin. Note that this command will only work if the validator is already registered. If the validator is not registered, use the register-validators command instead.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "aleo-program-id",
				Usage: "The program id of the aleo program. This is the program id of the program that is used to " +
					"register validators. This is the program id of the program that is used to register validators. " +
					"By default, this is the program id of the snarkOS program.",
				Aliases:  []string{"a"},
				Value:    cfg.Common.ProgramID,
				Required: false,
			},
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
				Name:     "website_url",
				Usage:    "The website url of the validator",
				Aliases:  []string{"w"},
				Value:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "logo_url",
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
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {
			// prepare the command
			var query = cfg.Common.AleoNodeUrl
			var broadcast = cfg.Common.AleoNodeUrl + "/testnet3/transaction/broadcast"
			var priorityFee = "0"
			var aleoProgramID = context.String("aleo-program-id")
			var privateKey = context.String("private-key")
			var validator = context.String("validator")
			var name = context.String("name")
			var websiteUrl = context.String("website_url")
			var logoUrl = context.String("logo_url")
			var description = context.String("description")

			var nameNumber = util.Utf8StringToBigInt(name)
			var websiteUrlNumber = util.Utf8StringToBigInt(websiteUrl)
			var logoUrlNumber = util.Utf8StringToBigInt(logoUrl)
			var descriptionNumber = util.Utf8StringToBigInt(description)

			// add the word "field" to the name, website url, logo url, and description
			nameField := fmt.Sprintf("%dfield", nameNumber)
			websiteUrlField := fmt.Sprintf("%dfield", websiteUrlNumber)
			logoUrlField := fmt.Sprintf("%dfield", logoUrlNumber)
			descriptionField := fmt.Sprintf("%dfield", descriptionNumber)

			// field can only be 251 bits, throw an error if it is too long
			nameFieldBitLen := util.CalculateLength(nameField)
			if nameFieldBitLen > 251 {
				fmt.Println("Name field is too long. It must be less than 251 bits.")
				return nil
			}

			websiteUrlFieldBitLen := util.CalculateLength(websiteUrlField)
			if websiteUrlFieldBitLen > 251 {
				fmt.Println("Website url field is too long. It must be less than 251 bits.")
				return nil
			}

			logoUrlFieldBitLen := util.CalculateLength(logoUrlField)
			if logoUrlFieldBitLen > 251 {
				fmt.Println("Logo url field is too long. It must be less than 251 bits.")
				return nil
			}

			descriptionFieldBitLen := util.CalculateLength(descriptionField)
			if descriptionFieldBitLen > 251 {
				fmt.Println("Description field is too long. It must be less than 251 bits.")
				return nil
			}

			// put quotes around the entire json body
			var inputString = "" + "\"" + "{\"validator_address\":\"" + validator + "\",\"name\":" + nameField + ", \"website_url\":" + websiteUrlField + ", \"logo_url\":" + logoUrlField + ", \"description\":" + descriptionField + "}" + "\""
			var inputStringF = fmt.Sprintf("snarkos developer execute %s update_validator_as_admin %s --private-key %s --query %s --broadcast %s --priority-fee %s", aleoProgramID, inputString, privateKey, query, broadcast, priorityFee)

			cmd := exec.Command("bash", "-c", inputStringF)

			fmt.Println("Executing command: ", cmd)
			cmd.Stdout = context.App.Writer
			cmd.Stderr = context.App.ErrWriter
			return nil
		},
	}
	updateValidatorRegistryCmds = append(updateValidatorRegistryCmds, updateValidatorRegistryCmd)
	updateValidatorRegistryCmds = append(updateValidatorRegistryCmds, updateValidatorInfoAsAdminCmd)

	return updateValidatorRegistryCmds
}
