package main

import (
	"avr-cli-backend/cmd"
	"avr-cli-backend/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var Commit string
var Version string

func main() {
	// database
	cfg := config.InitConfig()

	cfg.Common.Commit = Commit
	cfg.Common.Version = Version

	// get all the commands
	var commands []*cli.Command

	// commands
	commands = append(commands, cmd.UpdateRegistryCmd(&cfg)...)
	commands = append(commands, cmd.RegisterValidatorRegistryCmd(&cfg)...)
	commands = append(commands, cmd.InputGeneratorCmd(&cfg)...)
	commands = append(commands, cmd.LookupApiCmd(&cfg)...)
	commands = append(commands, cmd.InputUtilsCmd(&cfg)...)
	commands = append(commands, cmd.RefreshValidatorRegistry(&cfg)...)

	app := &cli.App{
		Commands:    commands,
		Name:        "avr-cli",
		Description: "a cli for aleo validator registry",
		Version:     fmt.Sprintf("%s+git.%s\n", cfg.Common.Version, cfg.Common.Commit),
	}

	// run the job

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}

	// Example usage
	//str := "this is the biggest string ever"
	//fmt.Println("Original String:", str)
	//
	//// Convert string to big.Int
	//bigIntResult := utf8StringToBigInt(str)
	//fmt.Println("BigInt:", bigIntResult)
	////460070329670722855200115
	////newBigInt := big.NewInt(460070329670722855200115)
	//
	//intMe := big.Int{}
	//intMe.SetString("2329422148041661608978791199966887706646635891", 10)
	//fmt.Println("intMe:", intMe)
	//
	//fmt.Println("New BigInt:", bigIntResult)
	//// Convert big.Int back to string
	//convertedString := bigIntToUTF8String(intMe)
	//fmt.Println("Converted String:", convertedString)

}
