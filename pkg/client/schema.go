package client

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
)

type GetProviderSchemaOptions struct {
	Provider string
}

type ProviderSchema struct {
	*cqproto.GetProviderSchemaResponse
	ProtocolVersion int
}

func GetProviderSchema(ctx context.Context, plugin *plugin.Manager, request *GetProviderSchemaOptions) (*ProviderSchema, error) {
	providerPlugin, err := plugin.CreatePlugin(request.Provider, "", nil)
	if err != nil {
		log.Error().Str("provider", request.Provider).Err(err).Msg("failed to create provider plugin")
		return nil, err // TODO: should be a diag
	}
	defer plugin.ClosePlugin(providerPlugin)

	schema, err := providerPlugin.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
	if err != nil {
		return nil, err // TODO: make a diag
	}
	log.Debug().Str("provider", request.Provider).Msg("retrieved provider schema successfully")
	return &ProviderSchema{
		GetProviderSchemaResponse: schema,
		ProtocolVersion:           providerPlugin.ProtocolVersion(),
	}, nil

}
