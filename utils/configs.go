package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func InitializeConfiguration(prefix, configFile, configPath string) {
	viper.SetEnvPrefix(prefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetConfigName(configFile)

	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("Fatal error config file: %s", err)
	}
}
