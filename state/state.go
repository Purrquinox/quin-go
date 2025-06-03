package state

import (
	"context"
	"os"
	"quin/genconfig"
	"quin/types"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/infinitybotlist/eureka/snippets"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	Data      *types.AIData  // Global configuration variable
	Secrets   *types.Secrets // Global secrets variable
	Logger    *zap.Logger
	Context   = context.Background()
	Validator = validator.New()
)

func Setup() {
	// Validator setup
	Validator.RegisterValidation("notblank", validators.NotBlank)
	Validator.RegisterValidation("nospaces", snippets.ValidatorNoSpaces)
	Validator.RegisterValidation("https", snippets.ValidatorIsHttps)
	Validator.RegisterValidation("httporhttps", snippets.ValidatorIsHttpOrHttps)

	// Generate the sample config file
	genconfig.GenConfigTo(types.AIData{}, "data.yaml.sample")
	genconfig.GenConfigTo(types.Secrets{}, "secrets.yaml.sample")

	// Read the configuration.
	cfg, err := os.ReadFile("data.yaml")
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	err = yaml.Unmarshal(cfg, &Data)
	if err != nil {
		panic("Failed to parse config file: " + err.Error())
	}

	// Read the secrets.
	sec, err := os.ReadFile("secrets.yaml")
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	err = yaml.Unmarshal(sec, &Secrets)
	if err != nil {
		panic("Failed to parse config file: " + err.Error())
	}
}
