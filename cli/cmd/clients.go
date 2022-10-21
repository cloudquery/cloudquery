package cmd

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
)

func initializeDestinationClients(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, cqDir string) (destClients []*clients.DestinationClient, closeClients func(), err error) {
	destClients = make([]*clients.DestinationClient, len(sourceSpec.Destinations))
	closeClients = func() {
		for _, destClient := range destClients {
			if destClient != nil {
				if err := destClient.Terminate(); err != nil {
					log.Error().Err(err).Msg("Failed to terminate destination client")
					fmt.Println("failed to terminate destination client: ", err)
				}
			}
		}
	}
	defer func() {
		if err != nil {
			// if an error occurs in this function, make sure we (try to) close any clients that were opened
			// before returning
			closeClients()
		}
	}()
	for i, destinationSpec := range destinationsSpecs {
		destClients[i], err = clients.NewDestinationClient(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version,
			clients.WithDestinationLogger(log.Logger),
			clients.WithDestinationDirectory(cqDir),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		if err = destClients[i].Initialize(ctx, destinationSpec); err != nil {
			return nil, nil, fmt.Errorf("failed to initialize destination plugin client for %s: %w", destinationSpec.Name, err)
		}
	}
	return destClients, closeClients, nil
}
