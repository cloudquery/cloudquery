package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	pluginSpecSchemaShort = "Export a plugin's spec JSON schema."
	pluginSpecSchemaLong  = `Export a plugin's spec JSON schema to a local file.
The exported file can later be passed to ` + "`cloudquery validate-config --schemas-dir`" + ` to validate
configurations fully offline, without spawning the plugin binary or contacting the CloudQuery registry.`
	pluginSpecSchemaExample = `
# Print schema to stdout
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0

# Write schema to a specific file
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0 -o aws.json

# Write schema to <dir>/<name>.json (suitable for --schemas-dir consumption)
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0 -D ./schemas`
)

func newCmdPluginSpecSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spec-schema <team_name>/<plugin_kind>/<plugin_name>@<version>",
		Short:   pluginSpecSchemaShort,
		Long:    pluginSpecSchemaLong,
		Example: pluginSpecSchemaExample,
		Args:    cobra.ExactArgs(1),
		RunE:    runPluginSpecSchema,
	}
	cmd.Flags().StringP("output", "o", "", "Write schema to this file. Mutually exclusive with --schemas-dir.")
	cmd.Flags().StringP("schemas-dir", "D", "", "Write schema to <dir>/<plugin-name>.json. Mutually exclusive with --output.")
	return cmd
}

func runPluginSpecSchema(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	schemasDir, err := cmd.Flags().GetString("schemas-dir")
	if err != nil {
		return err
	}
	if output != "" && schemasDir != "" {
		return errors.New("--output and --schemas-dir are mutually exclusive")
	}
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	ref, err := hub.ParseHubPluginRef(args[0])
	if err != nil {
		return err
	}

	pluginType, err := pluginTypeFromKind(ref.Kind)
	if err != nil {
		return err
	}

	ctx := cmd.Context()

	pluginCfg := managedplugin.Config{
		Name:     ref.Name,
		Version:  ref.Version,
		Path:     fmt.Sprintf("%s/%s", ref.TeamName, ref.Name),
		Registry: managedplugin.RegistryCloudQuery,
	}

	// CloudQuery-registry plugins always need an auth token.
	tc := cqapiauth.NewTokenClient()
	authToken, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name: %w", err)
	}

	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(authToken.Value),
		managedplugin.WithTeamName(teamName),
	}
	if logConsole {
		opts = append(opts, managedplugin.WithNoProgress())
	}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}
	if disableSentry {
		opts = append(opts, managedplugin.WithNoSentry())
	}

	clients, err := managedplugin.NewClients(ctx, pluginType, []managedplugin.Config{pluginCfg}, opts...)
	if err != nil {
		return enrichClientError(clients, []bool{false}, err)
	}
	defer func() {
		if err := clients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	if len(clients) == 0 {
		return errors.New("plugin client not initialized")
	}

	pluginClient := plugin.NewPluginClient(clients[0].Conn)
	jsonSchema, err := getSpecSchemaFromPlugin(ctx, pluginClient)
	if err != nil {
		return fmt.Errorf("failed to fetch spec schema: %w", err)
	}
	if len(jsonSchema) == 0 {
		return fmt.Errorf("plugin %s did not return a spec schema", ref.String())
	}

	return writeSchemaOutput(jsonSchema, ref.Name, output, schemasDir)
}

func pluginTypeFromKind(kind string) (managedplugin.PluginType, error) {
	switch kind {
	case "source":
		return managedplugin.PluginSource, nil
	case "destination":
		return managedplugin.PluginDestination, nil
	default:
		return 0, fmt.Errorf("unsupported plugin kind %q (expected source or destination)", kind)
	}
}

func writeSchemaOutput(jsonSchema, pluginName, output, schemasDir string) error {
	switch {
	case output != "":
		return os.WriteFile(output, []byte(jsonSchema), 0o644)
	case schemasDir != "":
		if err := os.MkdirAll(schemasDir, 0o755); err != nil {
			return err
		}
		path := filepath.Join(schemasDir, pluginName+".json")
		return os.WriteFile(path, []byte(jsonSchema), 0o644)
	default:
		_, err := fmt.Print(jsonSchema)
		return err
	}
}
