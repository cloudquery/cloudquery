package sync

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/internal/plugin"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	fetchShort   = "Sync resources from configured source plugins to destination"
	fetchExample = `# Sync configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery sync ./directory`
)

func NewCmdFetch() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [directory]",
		Short:   fetchShort,
		Long:    fetchShort,
		Example: fetchExample,
		Args:    cobra.RangeArgs(0, 1),
		RunE:    sync,
	}
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	directory := "."
	if len(args) == 1 {
		directory = args[0]
	}
	fmt.Println("Loading specs from directory: ", directory)
	specReader, err := specs.NewSpecReader(directory)
	if err != nil {
		return fmt.Errorf("failed to load specs from directory %s: %w", directory, err)
	}

	pm := plugin.NewPluginManager()
	for _, sourceSpec := range specReader.GetSources() {
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
	fmt.Println("Starting fetch for: ", sourceSpec.Name, "->", sourceSpec.Destinations)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync(ctx, sourceSpec, resources); err != nil {
			return errors.Wrap(err, "failed to fetch resources")
		}
		return nil
	})

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Fetching"),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
	)

	g.Go(func() error {
		for resource := range resources {
			// fmt.Println("fetched")
			bar.Add(1)
			for i, destination := range sourceSpec.Destinations {
				if err := destPlugins[i].GetClient().Write(ctx, resource); err != nil {
					return fmt.Errorf("failed to write resource for %s->%s: %w", sourceSpec.Name, destination, err)
				}
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		bar.Finish()
		return errors.Wrap(err, "failed to fetch resources")
	}
	bar.Finish()

	fmt.Println("Fetch completed successfully")
	return nil
}
