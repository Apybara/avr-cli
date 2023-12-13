package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type AleoValidatorRegistryCli struct {
	Config AleoValidatorRegistryCliConfig
}

type AleoValidatorRegistryCliConfig struct {
	Common struct {
		Commit           string `env:"COMMIT" envDefault:""`
		Version          string `env:"VERSION" envDefault:""`
		ProgramID        string `env:"PROGRAM_ID" envDefault:"av_registry.aleo"`
		ProgramOwnerPkey string `env:"PROGRAM_OWNER_PKEY" envDefault:""`
		AleoNodeUrl      string `env:"ALEO_NODE_URL" envDefault:"https://api.explorer.aleo.org/v1"`
		DB               string `env:"DB" envDefault:""`
	}
}

// The function initializes the configuration for an Aleo staking operator.
func InitConfig() AleoValidatorRegistryCliConfig {
	godotenv.Load() // load from environment OR .env file if it exists
	var cfg AleoValidatorRegistryCliConfig

	if err := env.Parse(&cfg); err != nil {
		log.Fatal("error parsing config: %+v\n", err)
	}

	return cfg
}
