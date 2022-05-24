package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/rs/zerolog/log"
)

type GetProviderSchemaOptions struct {
	Provider registry.Provider
}

type ProviderSchema struct {
	*cqproto.GetProviderSchemaResponse
	ProtocolVersion int
	Unmanaged       bool
}

func GetProviderSchema(ctx context.Context, manager *plugin.Manager, request *GetProviderSchemaOptions) (*ProviderSchema, diag.Diagnostics) {
	log.Info().Stringer("provider", request.Provider).Msg("requesting provider schema")
	providerPlugin, err := manager.CreatePlugin(&plugin.CreationOptions{Provider: request.Provider})
	if err != nil {
		log.Error().Stringer("provider", request.Provider).Err(err).Msg("failed to create provider plugin")
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	defer manager.ClosePlugin(providerPlugin)

	schema, err := providerPlugin.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	log.Debug().Stringer("provider", request.Provider).Msg("retrieved provider schema successfully")
	// set version if schema didn't return it.
	if schema.Version == "" {
		schema.Version = request.Provider.Version
	} else if !strings.HasPrefix(schema.Version, "v") {
		schema.Version = fmt.Sprintf("v%s", schema.Version)
	}

	return &ProviderSchema{
		GetProviderSchemaResponse: schema,
		ProtocolVersion:           providerPlugin.ProtocolVersion(),
		Unmanaged:                 providerPlugin.Version() == plugin.Unmanaged,
	}, nil

}
