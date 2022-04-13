package client

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/rs/zerolog/log"

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
}

func GetProviderSchema(ctx context.Context, manager *plugin.Manager, request *GetProviderSchemaOptions) (*ProviderSchema, diag.Diagnostics) {
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
	return &ProviderSchema{
		GetProviderSchemaResponse: schema,
		ProtocolVersion:           providerPlugin.ProtocolVersion(),
	}, nil

}
