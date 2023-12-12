package cmd

import (
	"avr-cli-backend/config"
	"avr-cli-backend/util"
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os/exec"
)

// commands to call indexer block, and staker
func RefreshValidatorRegistry(cfg *config.AleoValidatorRegistryCliConfig) []*cli.Command {
	// add a command to run API node
	var registerValidatorCmds []*cli.Command

	registerValidatorsCmd := &cli.Command{
		Name:        "refresh-registry",
		Description: "This command will register all validators in the registry. Note that this command will only work if the validators are not already registered. If the validators are already registered, use the update-registry command instead.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "every",
				Usage:   "The interval to refresh the registry in minutes",
				Aliases: []string{"e"},
				Value:   15,
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

			every := context.Duration("every")

			// go cron
			s := gocron.NewScheduler()

			s.Every(uint64(every)).Minutes().Do(func() error {
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
					var name = ""
					var websiteUrl = ""
					var logoUrl = ""
					var description = ""

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
					fmt.Println("Name field bit length: ", nameFieldBitLen)
					if nameFieldBitLen > 256 {
						fmt.Println("Name field is too long. It must be less than 251 bits.")
						return nil
					}

					websiteUrlFieldBitLen := util.CalculateLength(websiteUrlField)
					fmt.Println("Website url field bit length: ", websiteUrlFieldBitLen)
					if websiteUrlFieldBitLen > 256 {
						fmt.Println("Website url field is too long. It must be less than 251 characters.")
						return nil
					}

					logoUrlFieldBitLen := util.CalculateLength(logoUrlField)
					if logoUrlFieldBitLen > 256 {
						fmt.Println("Logo url field is too long. It must be less than 251 bits.")
						return nil
					}

					descriptionFieldBitLen := util.CalculateLength(descriptionField)
					if descriptionFieldBitLen > 256 {
						fmt.Println("Description field is too long. It must be less than 251 bits.")
						return nil
					}

					// put quotes around the entire json body
					var inputString = "" + "\"" + "{\"validator_address\":\"" + validator + "\",\"name\":" + nameField + ", \"website_url\":" + websiteUrlField + ", \"logo_url\":" + logoUrlField + ", \"description\":" + descriptionField + "}" + "\""
					var inputStringF = fmt.Sprintf("snarkos developer execute %s register_validator %s --private-key %s --query %s --broadcast %s --priority-fee %s", cfg.Common.ProgramID, inputString, privateKey, query, broadcast, priorityFee)
					fmt.Println("Executing command: ", inputStringF)
					cmd := exec.Command("bash", "-c", inputStringF)

					fmt.Println("Executing command: ", cmd)
					cmd.Stdout = context.App.Writer
					cmd.Stderr = context.App.ErrWriter

					// execute the command
					err := cmd.Run()

					// check if there was an error
					if err != nil {
						fmt.Println("Error executing command: ", err)
						return nil
					}

					// output the result
					fmt.Println("Successfully registered validator")

				}
				return nil
			})
			s.Start()

			for {
				// do nothing
			}
			return nil

		},
	}
	registerValidatorCmds = append(registerValidatorCmds, registerValidatorsCmd)

	return registerValidatorCmds
}
