package cmd

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
)

type destinationClients []*clients.DestinationClient

func newDestinationClients(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, cqDir string) (destinationClients, error) {
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
		opts := []clients.DestinationClientOption{
			clients.WithDestinationLogger(log.Logger),
			clients.WithDestinationDirectory(cqDir),
		}
		if disableSentry {
			opts = append(opts, clients.WithDestinationNoSentry())
		}
		destClients[i], err = clients.NewDestinationClient(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version, opts...)
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
