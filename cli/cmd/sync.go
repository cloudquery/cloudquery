package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/cloudquery/cloudquery/cli/internal/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
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
		Args:    cobra.MaximumNArgs(1),
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
	specReader, err := specs.NewSpecReader(directory)
	if err != nil {
		return fmt.Errorf("failed to load specs from directory %s: %w", directory, err)
	}

	if len(specReader.GetSources()) == 0 {
		return fmt.Errorf("no sources found in directory: %s", directory)
	}

	pm := plugins.NewPluginManager()
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

func syncConnection(ctx context.Context, pm *plugins.PluginManager, specReader *specs.SpecReader, sourceSpec specs.Source) error {
	sourcePlugin, err := pm.NewSourcePlugin(ctx, &sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer sourcePlugin.Close()
	sourceClient := sourcePlugin.GetClient()

	destPlugins := make([]*plugins.DestinationPlugin, len(sourceSpec.Destinations))
	defer func() {
		for _, destPlugin := range destPlugins {
			if destPlugin != nil {
				destPlugin.Close()
			}
		}
	}()
	for i, destination := range sourceSpec.Destinations {
		spec := specReader.GetDestinationByName(destination)
		if spec == nil {
			return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
		}
		plugin, err := pm.NewDestinationPlugin(ctx, *spec)
		if err != nil {
			return fmt.Errorf("failed to create destination plugin client for %s: %w", destination, err)
		}
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

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	startTime := time.Now()
	format := " Syncing (%d resources) %s"
	s.Suffix = fmt.Sprintf(format, 0, time.Duration(0))
	s.Start()
	failedWrites := 0
	totalResources := 0
	g.Go(func() error {
		for resource := range resources {
			totalResources++
			s.Suffix = fmt.Sprintf(format, totalResources, time.Since(startTime).Truncate(time.Second))
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
		s.Stop()
		return fmt.Errorf("failed to fetch resources: %w", err)
	}
	s.Stop()
	fmt.Println("Fetch completed successfully.")
	fmt.Printf("Summary: Resources: %d, Failed Writes: %d, Fetch Errors: %d, Fetch Warnings: %d\n",
		totalResources, failedWrites, sourcePlugin.Errors(), sourcePlugin.Warnings())
	if sourcePlugin.Errors() > 0 || sourcePlugin.Warnings() > 0 || failedWrites > 0 {
		fmt.Println("Please check the logs for more details on errors/warnings.")
	}
	return nil
}
