package util

import (
	"strings"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

const (
	Commit     = "development"
	Date       = "unknown"
	ConfigHelp = "Path to configuration file. can be generated with 'init {provider}' command (env: CQ_CONFIG_PATH)"
)

var (
	InstanceId = uuid.New()
)

// GetConfigFile returns the config filename
// if it ends with ".*", .hcl and .yml extensions are tried in order to find the existing file, if available
func GetConfigFile(path string) string {
	configPath := viper.GetString(path)
	if !strings.HasSuffix(configPath, ".*") {
		return configPath
	}

	fs := file.NewOsFs()
	noSuffix := strings.TrimSuffix(configPath, ".*")
	for _, tryExt := range []string{".hcl", ".yml"} {
		tryFn := noSuffix + tryExt
		if _, err := fs.Stat(tryFn); err == nil {
			return tryFn
		}
	}

	return noSuffix + ".hcl"
}
