package config

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
)

type ConfigOptions struct {
	ConfigFile string `long:"config" description:"Path to the configuration file"`
}

func ImportConfig() error {
	var opts ConfigOptions

	parser := flags.NewParser(&opts, flags.Default|flags.IgnoreUnknown)

	_, err := parser.Parse()
	if err != nil {
		return err
	}

	if opts.ConfigFile != "" {
		viper.SetConfigFile(opts.ConfigFile)
	} else {
		viper.SetConfigFile("config/default_config.json")
	}

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
