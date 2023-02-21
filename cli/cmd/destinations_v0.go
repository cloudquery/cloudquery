package cmd

import (
	"context"
	"fmt"

	destination "github.com/cloudquery/plugin-sdk/clients/destination/v0"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
)

type destinationClients []*destination.Client

func newDestinationClientsV0(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, cqDir string) (destinationClients, error) {
	var err error
	destClients := make(destinationClients, len(sourceSpec.Destinations))
	defer func() {
		if err != nil {
			// if an error occurs in this function, make sure we (try to) close any clients that were opened
			// before returning
			destClients.Close()
		}
	}()

	for i, destinationSpec := range destinationsSpecs {
		opts := []destination.ClientOption{
			destination.WithLogger(log.Logger),
			destination.WithDirectory(cqDir),
		}
		if disableSentry {
			opts = append(opts, destination.WithNoSentry())
		}
		destClients[i], err = destination.NewClient(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version, opts...)
		if err != nil {
			return nil, fmt.Errorf("failed to create destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		if err = destClients[i].Initialize(ctx, destinationSpec); err != nil {
			return nil, fmt.Errorf("failed to initialize destination plugin client for %s: %w", destinationSpec.Name, err)
		}
	}
	return destClients, nil
}

func (c destinationClients) Close() {
	for _, destClient := range c {
		if destClient != nil {
			if err := destClient.Terminate(); err != nil {
				log.Error().Err(err).Msg("Failed to terminate destination client")
				fmt.Println("failed to terminate destination client: ", err)
			}
		}
	}
}
