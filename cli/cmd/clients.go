package cmd

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
)

func initializeDestinationClients(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, cqDir string) ([]*clients.DestinationClient, error) {
	var err error
	destClients := make([]*clients.DestinationClient, len(sourceSpec.Destinations))
	for i, destinationSpec := range destinationsSpecs {
		destClients[i], err = clients.NewDestinationClient(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version,
			clients.WithDestinationLogger(log.Logger),
			clients.WithDestinationDirectory(cqDir),
		)
		if err != nil {
			return destClients, fmt.Errorf("failed to create destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		if err := destClients[i].Initialize(ctx, destinationSpec); err != nil {
			return destClients, fmt.Errorf("failed to initialize destination plugin client for %s: %w", destinationSpec.Name, err)
		}
	}
	return destClients, nil
}
