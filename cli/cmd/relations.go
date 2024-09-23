package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	relationsShort   = "Show relations between tables"
	relationsExample = `# Show relations between tables
cloudquery relations ./directory
`
)

func NewCmdRelations() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "relations [files or directories]",
		Short:   relationsShort,
		Long:    relationsShort,
		Example: relationsExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    relations,
	}
	cmd.Flags().Int("min-depth", 1, "Minimum depth to show relations")
	return cmd
}

func relations(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewRelaxedSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sources := specReader.Sources
	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, nil, nil)
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
	minDepth, err := cmd.Flags().GetInt("min-depth")
	if err != nil {
		return err
	}
	pluginConfigs := make([]managedplugin.Config, len(sources))
	sourceRegInferred := make([]bool, len(sources))
	for i, sourceSpec := range sources {
		pluginConfigs[i] = managedplugin.Config{
			Name:       sourceSpec.Name,
			Path:       sourceSpec.Path,
			Version:    sourceSpec.Version,
			Registry:   SpecRegistryToPlugin(sourceSpec.Registry),
			DockerAuth: sourceSpec.DockerRegistryAuthToken,
		}
		sourceRegInferred[i] = sourceSpec.Registry == specs.RegistryUnset
	}

	sourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, pluginConfigs, opts...)
	if err != nil {
		return enrichClientError(sourceClients, sourceRegInferred, err)
	}
	defer func() {
		if err := sourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	for _, source := range specReader.Sources {
		cl := sourceClients.ClientByName(source.Name)
		if err := runRelations(ctx, cl, source, minDepth); err != nil {
			return err
		}
	}

	return nil
}

func runRelations(ctx context.Context, cl *managedplugin.Client, source *specs.Source, minDepth int) error {
	sourcePbClient := pluginPb.NewPluginClient(cl.Conn)
	if err := initPlugin(ctx, sourcePbClient, source.Spec, true, invocationUUID.String()); err != nil {
		return fmt.Errorf("failed to init source: %w", err)
	}
	req := &pluginPb.GetTables_Request{
		Tables: []string{"*"},
	}
	getTablesResp, err := sourcePbClient.GetTables(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get tables: %w", err)
	}
	schemas, err := pluginPb.NewSchemasFromBytes(getTablesResp.Tables)
	if err != nil {
		return fmt.Errorf("failed to parse schemas: %w", err)
	}
	tables, err := schema.NewTablesFromArrowSchemas(schemas)
	if err != nil {
		return fmt.Errorf("failed to convert schemas to tables: %w", err)
	}
	topLevelTables, err := tables.UnflattenTables()
	if err != nil {
		return fmt.Errorf("failed to unflatten tables: %w", err)
	}
	for _, table := range topLevelTables {
		printTableRelationsRecursive(table, "", 1, minDepth)
	}
	return nil
}

func printTableRelationsRecursive(table *schema.Table, path string, depth, minDepth int) {
	if path == "" {
		path = table.Name
	} else {
		path += " -> " + table.Name
	}

	if len(table.Relations) == 0 {
		if depth >= minDepth {
			fmt.Println(path)
		}
	} else {
		for _, relation := range table.Relations {
			printTableRelationsRecursive(relation, path, depth+1, minDepth)
		}
	}
}
