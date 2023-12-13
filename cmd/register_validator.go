package cmd

import (
	"avr-cli-backend/config"
	"avr-cli-backend/util"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os/exec"
	"time"
)

type AleoValidatorResponse struct {
	StartingRound int64             `json:"starting_round"`
	Members       map[string]string `json:"members"`
	TotalStake    int64             `json:"total_stake"`
}

// commands to call indexer block, and staker
func RegisterValidatorRegistryCmd(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var registerValidatorCmds []*cli.Command

	registerValidatorCmd := &cli.Command{
		Name:        "register-validator",
		Usage:       "Register a validator",
		Description: "This command will register a validator. Note that this command will only work if the validator is not already registered. If the validator is already registered, use the update-registry command instead.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "aleo-program-id",
				Usage:    "The program id of the aleo program",
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
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {

			// prepare the command
			var query = cfg.Common.AleoNodeUrl
			var broadcast = cfg.Common.AleoNodeUrl + "/testnet3/transaction/broadcast"
			var priorityFee = "0"
			var privateKey = context.String("private-key")
			var aleoProgramID = context.String("aleo-program-id")
			var validator = context.String("validator")
			var name = context.String("name")
			var websiteUrl = context.String("website-url")
			var logoUrl = context.String("logo-url")
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
			var inputStringF = fmt.Sprintf("snarkos developer execute %s register_validator %s --private-key %s --query %s --broadcast %s --priority-fee %s", aleoProgramID, inputString, privateKey, query, broadcast, priorityFee)
			fmt.Println("Executing command: ", inputStringF)
			cmd := exec.Command("bash", "-c", inputStringF)

			fmt.Println("Executing command: ", cmd)
			cmd.Stdout = context.App.Writer
			cmd.Stderr = context.App.ErrWriter

			// execute the command
			err := cmd.Run()

			// check if there was an error
			if err != nil {
				return err
			}

			// output the result
			fmt.Println("Successfully registered validator")

			return nil

		},
	}
	registerValidatorsCmd := &cli.Command{
		Name:        "register-all-validators",
		Description: "This command will register all validators in the registry. Note that this command will only work if the validators are not already registered. If the validators are already registered, use the update-registry command instead.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "private-key",
				Usage:    "The private key of the validator",
				Aliases:  []string{"p"},
				Value:    "",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {

			url := cfg.Common.AleoNodeUrl + "/testnet3/latest/committee"
			committeeResp, err := http.Get(url)
			if err != nil {
				return err
			}

			defer committeeResp.Body.Close()
			bodyGetPage, _ := io.ReadAll(committeeResp.Body)

			var aleoValidatorResponse AleoValidatorResponse
			json.Unmarshal(bodyGetPage, &aleoValidatorResponse)

			// cache the result
			if aleoValidatorResponse.Members == nil || len(aleoValidatorResponse.Members) == 0 {
				fmt.Println("No validators found")
				return nil
			}
			for i := range aleoValidatorResponse.Members {

				// prepare the command
				var query = cfg.Common.AleoNodeUrl
				var broadcast = cfg.Common.AleoNodeUrl + "/testnet3/transaction/broadcast"
				var priorityFee = "0"
				var privateKey = context.String("private-key")
				var validator = i
				var name = " "
				var websiteUrl = " "
				var logoUrl = " "
				var description = " "

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
				var inputStringF = fmt.Sprintf("snarkos developer execute %s register_validator %s --private-key %s --query %s --broadcast %s --priority-fee %s", cfg.Common.ProgramID, inputString, privateKey, query, broadcast, priorityFee)

				cmd := exec.Command("bash", "-c", inputStringF)

				cmd.Stdout = context.App.Writer
				cmd.Stderr = context.App.ErrWriter

				fmt.Println("Executing command with the following fields:", "validator_address:", validator, "name:", nameField, "website_url:", websiteUrlField, "logo_url:", logoUrlField, "description:", descriptionField)

				// execute the command
				err := cmd.Run()

				// check if there was an error
				if err != nil {
					return err
				}

				// output the result
				fmt.Println("Successfully registered validator")
				time.Sleep(5 * time.Second)
			}
			return nil

		},
	}
	registerValidatorCmds = append(registerValidatorCmds, registerValidatorCmd)
	registerValidatorCmds = append(registerValidatorCmds, registerValidatorsCmd)

	return registerValidatorCmds
}
