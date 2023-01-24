package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/clients/source/v0"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
)

func syncConnectionV0(ctx context.Context, cqDir string, sourceSpec specs.Source, destinationsSpecs []specs.Destination, uid string, noMigrate bool) error {
	opts := []source.ClientOption{
		source.WithLogger(log.Logger),
		source.WithDirectory(cqDir),
	}
	if disableSentry {
		opts = append(opts, source.WithNoSentry())
	}
	sourceClient, err := source.NewClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version, opts...)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	//nolint:revive
	defer func() {
		if err := sourceClient.Terminate(); err != nil {
			log.Error().Err(err).Msg("Failed to terminate source client")
			fmt.Println("failed to terminate source client: ", err)
		}
	}()

	v, err := sourceClient.GetProtocolVersion(ctx)
	if err != nil {
		return fmt.Errorf("failed to get protocol version for source %s: %w", sourceSpec.Name, err)
	}
	switch v {
	case 1:
		if err := syncConnectionV0_1(ctx, cqDir, sourceClient, sourceSpec, destinationsSpecs, uid, noMigrate); err != nil {
			return err
		}
	case 2:
		if err := syncConnectionV0_2(ctx, cqDir, sourceClient, sourceSpec, destinationsSpecs, uid, noMigrate); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown protocol version %d for source %s", v, sourceSpec.Name)
	}
	return nil
}
