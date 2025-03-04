package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type LogConfig struct {
	EnableStacktrace bool `envconfig:"CODEGEN_CLI_LOG_ENABLE_STACKTRACE" default:"false"`
}

type CodegenSvcConfig struct {
	URL string `envconfig:"CODEGEN_CLI_CODEGEN_URL" required:"true"`
}

type Config struct {
	Debug bool `envconfig:"CODEGEN_CLI_DEBUG" default:"false"`

	Log        LogConfig
	CodegenSvc CodegenSvcConfig
}

func New() *Config {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		err = fmt.Errorf("error while parse env config | %w", err)

		log.Fatal(err)
	}

	return cfg
}
