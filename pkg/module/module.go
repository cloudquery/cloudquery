package module

import (
	"context"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/rs/zerolog/log"
)

type GetModuleOptions struct {
	Provider registry.Provider
	Request  cqproto.GetModuleRequest
}

func GetProviderModule(ctx context.Context, pm *plugin.Manager, opts *GetModuleOptions) (*cqproto.GetModuleResponse, error) {
	providerPlugin, err := pm.CreatePlugin(&plugin.CreationOptions{Provider: opts.Provider})
	if err != nil {
		log.Error().Stringer("provider", opts.Provider).Err(err).Msg("failed to create provider plugin")
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	defer pm.ClosePlugin(providerPlugin)

	inf, err := providerPlugin.Provider().GetModuleInfo(ctx, &opts.Request)
	if err != nil && strings.Contains(err.Error(), `unknown method GetModuleInfo`) {
		return &cqproto.GetModuleResponse{}, nil
	}

	return inf, err
}
