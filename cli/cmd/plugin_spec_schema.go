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
	pluginSpecSchemaLong  = `Export a plugin's spec JSON schema.

Only plugins published to the CloudQuery hub are supported. registry: local,
registry: grpc, and registry: docker plugins are not exportable because they
have no stable (path, version) identity to anchor the generated filename.
For those registries the plugin binary is already accessible locally, so
'cloudquery validate-config' can validate them in-place without --schemas-dir.

Without --schemas-dir the schema is printed to stdout. With --schemas-dir the
schema is written to <dir>/<plugin-name>@<version>.json, which is the
filename format expected by ` + "`cloudquery validate-config --schemas-dir`" + `.
Including the version in the filename ensures validation always runs against
the schema matching the plugin version in the config.`
	pluginSpecSchemaExample = `
# Print schema to stdout
cloudquery plugin spec-schema cloudquery/source/aws@v33.0.0

# Write to ./schemas/aws@v33.0.0.json
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
	cmd.Flags().StringP("schemas-dir", "D", "", "Write schema to <dir>/<plugin-name>@<version>.json. If omitted, the schema is printed to stdout.")
	return cmd
}

func runPluginSpecSchema(cmd *cobra.Command, args []string) error {
	schemasDir, err := cmd.Flags().GetString("schemas-dir")
	if err != nil {
		return err
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

	// Only registry: cloudquery is supported. The hub-ref input format already
	// constrains us to this registry, but the hardcoded value below is the
	// single source of truth — do not add a flag that lets callers select
	// local / grpc / docker without first defining how the exported filename
	// (which is keyed on a stable name+version) should be derived for those.
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

	return writeSchemaOutput(jsonSchema, ref.Name, ref.Version, schemasDir)
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

func writeSchemaOutput(jsonSchema, pluginName, pluginVersion, schemasDir string) error {
	if schemasDir == "" {
		_, err := fmt.Print(jsonSchema)
		return err
	}
	if err := os.MkdirAll(schemasDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(schemasDir, schemaFileName(pluginName, pluginVersion)), []byte(jsonSchema), 0o644)
}

// schemaFileName returns the canonical filename for a plugin's schema under --schemas-dir.
// Version is included whenever non-empty so consumers can pin validation to the right plugin version.
func schemaFileName(pluginName, pluginVersion string) string {
	if pluginVersion == "" {
		return pluginName + ".json"
	}
	return pluginName + "@" + pluginVersion + ".json"
}
