package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/internal/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	fetchShort   = "Sync resources from configured source plugins to destinations"
	fetchExample = `# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml
`
)

func NewCmdSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [file or directories...]",
		Short:   fetchShort,
		Long:    fetchShort,
		Example: fetchExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    sync,
	}
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	directory := "."
	if len(args) > 0 {
		directory = args[0]
	}
	fmt.Println("Loading specs from directory: ", directory)
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load specs from directory %s: %w", directory, err)
	}

	pm := plugins.NewPluginManager()
	for _, sourceSpec := range specReader.Sources {
		if len(sourceSpec.Destinations) == 0 {
			return fmt.Errorf("no destinations found for source %s", sourceSpec.Name)
		}
		var destinationsSpecs []specs.Destination
		for _, destination := range sourceSpec.Destinations {
			spec := specReader.Destinations[destination]
			if spec == nil {
				return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
			}
			destinationsSpecs = append(destinationsSpecs, *spec)
		}
		if err := syncConnection(ctx, pm, *sourceSpec, destinationsSpecs); err != nil {
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
	}

	return nil
}

func syncConnection(ctx context.Context, pm *plugins.PluginManager, sourceSpec specs.Source, destinationsSpecs []specs.Destination) error {
	sourcePlugin, err := pm.NewSourcePlugin(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer sourcePlugin.Close()
	sourceClient := sourcePlugin.GetClient()

	destPlugins := make([]*plugins.DestinationPlugin, len(sourceSpec.Destinations))
	destSubscriptions := make([]chan []byte, len(sourceSpec.Destinations))
	for i := range destSubscriptions {
		destSubscriptions[i] = make(chan []byte)
	}
	defer func() {
		for _, destPlugin := range destPlugins {
			if destPlugin != nil {
				destPlugin.Close()
			}
		}
	}()
	for i, destinationSpec := range destinationsSpecs {
		plugin, err := pm.NewDestinationPlugin(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version)
		if err != nil {
			return fmt.Errorf("failed to create destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		destPlugins[i] = plugin
		if err := destPlugins[i].GetClient().Initialize(ctx, destinationSpec); err != nil {
			return fmt.Errorf("failed to initialize destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		tables, err := sourceClient.GetTables(ctx)
		if err != nil {
			return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
		}

		if err := destPlugins[i].GetClient().Migrate(ctx, tables); err != nil {
			return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destinationSpec.Name, err)
		}
	}

	resources := make(chan []byte)
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
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
	)
	failedWrites := uint64(0)
	totalResources := 0
	for i, destination := range sourceSpec.Destinations {
		i := i
		destination := destination
		g.Go(func() error {
			var destFailedWrites uint64
			var err error
			if destFailedWrites, err = destPlugins[i].GetClient().Write(ctx, destSubscriptions[i]); err != nil {
				log.Error().Err(err).Msgf("failed to write for %s->%s", sourceSpec.Name, destination)
			}
			failedWrites += destFailedWrites
			return nil
		})
	}

	g.Go(func() error {
		for resource := range resources {
			totalResources++
			_ = bar.Add(1)
			for i := range destSubscriptions {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case destSubscriptions[i] <- resource:
				}
			}
		}
		for i := range destSubscriptions {
			close(destSubscriptions[i])
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		_ = bar.Finish()
		return fmt.Errorf("failed to fetch resources: %w", err)
	}
	_ = bar.Finish()
	fmt.Println("Fetch completed successfully.")
	fmt.Printf("Summary: Resources: %d, Failed Writes: %d, Fetch Errors: %d, Fetch Warnings: %d\n",
		totalResources, failedWrites, sourcePlugin.Errors(), sourcePlugin.Warnings())
	if sourcePlugin.Errors() > 0 || sourcePlugin.Warnings() > 0 || failedWrites > 0 {
		fmt.Println("Please check the logs for more details on errors/warnings.")
	}
	return nil
}
