package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/internal/plugin"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	fetchShort   = "Sync resources from configured source plugins to destination"
	fetchExample = `# Sync configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery sync ./directory`
)

func NewCmdSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [directory]",
		Short:   fetchShort,
		Long:    fetchShort,
		Example: fetchExample,
		Args:    cobra.ExactArgs(1),
		RunE:    sync,
	}
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	directory := args[0]
	fmt.Println("Loading specs from directory: ", directory)
	specReader, err := specs.NewSpecReader(directory)
	if err != nil {
		return fmt.Errorf("failed to load specs from directory %s: %w", directory, err)
	}

	if len(specReader.GetSources()) == 0 {
		return fmt.Errorf("no sources found in directory: %s", directory)
	}

	pm := plugin.NewPluginManager()
	for _, sourceSpec := range specReader.GetSources() {
		if len(sourceSpec.Destinations) == 0 {
			return fmt.Errorf("no destinations found for source %s", sourceSpec.Name)
		}
		if err := syncConnection(ctx, pm, specReader, sourceSpec); err != nil {
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
	}

	return nil
}

func syncConnection(ctx context.Context, pm *plugin.PluginManager, specReader *specs.SpecReader, sourceSpec specs.Source) error {
	sourcePlugin, err := pm.NewSourcePlugin(ctx, sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer sourcePlugin.Close()
	sourceClient := sourcePlugin.GetClient()

	destPlugins := make([]*plugin.DestinationPlugin, len(sourceSpec.Destinations))
	for i, destination := range sourceSpec.Destinations {
		spec := specReader.GetDestinatinoByName(destination)
		if spec == nil {
			return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
		}
		plugin, err := pm.NewDestinationPlugin(ctx, *spec)
		if err != nil {
			return fmt.Errorf("failed to create destination plugin client for %s: %w", destination, err)
		}
		defer plugin.Close()
		destPlugins[i] = plugin
		if err := destPlugins[i].GetClient().Initialize(ctx, *spec); err != nil {
			return fmt.Errorf("failed to initialize destination plugin client for %s: %w", destination, err)
		}
		tables, err := sourceClient.GetTables(ctx)
		if err != nil {
			return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
		}

		if err := destPlugins[i].GetClient().Migrate(ctx, tables); err != nil {
			return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destination, err)
		}
	}

	resources := make(chan *schema.Resource)
	g, ctx := errgroup.WithContext(ctx)
	fmt.Println("Starting sync for: ", sourceSpec.Name, "->", sourceSpec.Destinations)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync(ctx, sourceSpec, resources); err != nil {
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
		return nil
	})

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Syncing"),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("resources"),
	)
	failedWrites := 0
	totalResources := 0
	g.Go(func() error {
		for resource := range resources {
			_ = bar.Add(1)
			totalResources++
			for i, destination := range sourceSpec.Destinations {
				if err := destPlugins[i].GetClient().Write(ctx, resource.TableName, resource.Data); err != nil {
					failedWrites++
					log.Error().Err(err).Msgf("failed to write resource for %s->%s", sourceSpec.Name, destination)
				}
			}
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		_ = bar.Finish()
		return fmt.Errorf("failed to fetch resources: %w", err)
	}
	_ = bar.Finish()
	fmt.Println("Fetch completed successfully.")
	fmt.Printf("Summary: Resources: %d FailedWrites: %d\n", totalResources, failedWrites)
	return nil
}
