package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

const (
	configShort = "Manage CloudQuery CLI configuration values"
	configLong  = `Manage CloudQuery CLI configuration values.

Use the get and set commands to get or set a configuration value.
`
	configExample = `
# Set the current team
cloudquery config set team my-team

# Get the current team
cloudquery config get team
`

	configPath = "cloudquery/config.json"
)

var configKeys = []string{
	"team",
}

func newCmdConfig() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "config",
		Short:   configShort,
		Long:    configLong,
		Example: configExample,
		Hidden:  true,
	}
	getCmd := newCmdGet()
	setCmd := newCmdSet()
	rootCmd.AddCommand(getCmd, setCmd)
	return rootCmd
}

func newCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <key>",
		Short: "Get a config value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getConfig(cmd, args[0])
		},
	}
	return cmd
}

func newCmdSet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a config value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return setConfig(args[0], args[1])
		},
	}
	return cmd
}

func getConfig(cmd *cobra.Command, key string) error {
	if !slices.Contains(configKeys, key) {
		return fmt.Errorf("invalid config key %v (options are: %v)", key, strings.Join(configKeys, ", "))
	}
	val, err := readConfigFileValue(key)
	if err != nil {
		return err
	}
	cmd.Println(val)
	return nil
}

func setConfig(key, val string) error {
	if !slices.Contains(configKeys, key) {
		return fmt.Errorf("invalid config key %v (options are: %v)", key, strings.Join(configKeys, ", "))
	}
	configFilePath, err := xdg.ConfigFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}
	var config map[string]any
	b, err := os.ReadFile(configFilePath)
	switch {
	case err == nil:
		err = json.Unmarshal(b, &config)
		if err != nil {
			return fmt.Errorf("failed to parse config file: %w", err)
		}
	case os.IsNotExist(err):
		config = make(map[string]any)
	default:
		return fmt.Errorf("failed to read config file: %w", err)
	}
	config[key] = val
	b, err = json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(configFilePath, b, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}

func readConfigFileValue(key string) (string, error) {
	configFilePath, err := xdg.ConfigFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to get config file path: %w", err)
	}
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}
	var config map[string]any
	err = json.Unmarshal(b, &config)
	if err != nil {
		return "", fmt.Errorf("failed to parse config file: %w", err)
	}
	return config[key].(string), nil
}
