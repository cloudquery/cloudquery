package utils

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var InstanceId = uuid.New()

// getConfigFile returns the config filename
func GetConfigFile() string {
	return viper.GetString("configPath")
}
