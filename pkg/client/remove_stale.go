package client

import (
	"context"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/execution"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cq-provider-sdk/database"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/rs/zerolog/log"
)

type RemoveStaleDataOptions struct {
	Providers  []string
	LastUpdate time.Duration
}

type RemoveStaleDataResult struct{}

func RemoveStaleData(ctx context.Context, storage Storage, plugin *plugin.Manager, request *RemoveStaleDataOptions) (*RemoveStaleDataResult, diag.Diagnostics) {
	if len(request.Providers) == 0 {
		return nil, diag.Diagnostics{diag.NewBaseError(nil, diag.INTERNAL, diag.WithSeverity(diag.WARNING), diag.WithSummary("no providers were given"))}
	}

	db, err := database.New(ctx, logging.NewZHcLog(&log.Logger, "database"), storage.DSN())
	if err != nil {
		return nil, diag.Diagnostics{diag.NewBaseError(err, diag.INTERNAL, diag.WithSeverity(diag.ERROR))}
	}

	var diags diag.Diagnostics
	for _, p := range request.Providers {
		log.Debug().Str("provider", p).TimeDiff("since", time.Now().Add(-request.LastUpdate), time.Now()).Msg("cleaning stale data for provider")
		diags = diags.Add(removeProviderStaleData(ctx, db, plugin, p, request.LastUpdate))
	}
	return &RemoveStaleDataResult{}, diags
}

func removeProviderStaleData(ctx context.Context, storage execution.Storage, plugin *plugin.Manager, provider string, lastUpdate time.Duration) error {
	schema, err := GetProviderSchema(ctx, plugin, &GetProviderSchemaOptions{Provider: provider})
	if err != nil {
		return err
	}
	var diags diag.Diagnostics
	for _, t := range schema.ResourceTables {
		if err := storage.RemoveStaleData(ctx, t, time.Now().Add(-lastUpdate), nil); err != nil {
			diags = diags.Add(diag.NewBaseError(err, diag.DATABASE, diag.WithSeverity(diag.WARNING),
				diag.WithSummary("failed to remove stale data from %s", t.Name)))
		}
	}
	return diags
}
