package client

import (
	"context"
	"sort"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"

	"github.com/cloudquery/cq-provider-sdk/provider/execution"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cq-provider-sdk/database"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/rs/zerolog/log"
)

type PurgeProviderDataOptions struct {
	// Providers to purge data from, the provider name should be the plugin name
	Providers []string
	// LastUpdate defines how long from time.Now() should the resources be removed from the database.
	LastUpdate time.Duration
	// DryRun whether to run the purge in "dry" mode and only report the amount of affected resources that will be purged, if executed.
	DryRun bool
}

type PurgeProviderDataResult struct {
	// Total amount of affected resources, this value is only returned when dry run is set to true
	TotalAffected int
	// AffectedResources is all tables that have one or more resources affected
	AffectedResources map[string]int
}

func (p PurgeProviderDataResult) Resources() []string {
	r := make([]string, 0, len(p.AffectedResources))
	for rn := range p.AffectedResources {
		r = append(r, rn)
	}
	sort.Strings(r)
	return r
}

// PurgeProviderData purges resources that were not updated recently, if dry run is set to true, no resources will be removed.
func PurgeProviderData(ctx context.Context, storage Storage, plugin *plugin.Manager, opts *PurgeProviderDataOptions) (*PurgeProviderDataResult, diag.Diagnostics) {
	if len(opts.Providers) == 0 {
		return nil, diag.Diagnostics{diag.NewBaseError(nil, diag.INTERNAL, diag.WithSeverity(diag.WARNING), diag.WithSummary("no providers were given"))}
	}
	log.Info().Strs("providers", opts.Providers).Bool("dry-run", opts.DryRun).Msg("purging stale data for providers")
	db, err := database.New(ctx, logging.NewZHcLog(&log.Logger, "database"), storage.DSN())
	if err != nil {
		return nil, diag.Diagnostics{diag.NewBaseError(err, diag.INTERNAL)}
	}
	defer db.Close()
	var (
		diags  diag.Diagnostics
		result = PurgeProviderDataResult{
			TotalAffected:     0,
			AffectedResources: make(map[string]int),
		}
	)

	lastUpdateTime := time.Now().UTC().Add(-opts.LastUpdate)
	for _, p := range opts.Providers {
		log.Debug().Str("provider", p).TimeDiff("since", lastUpdateTime, time.Now().UTC()).Msg("cleaning stale data for provider")
		affectedResources, affected, err := removeProviderStaleData(ctx, db, plugin, p, lastUpdateTime, opts.DryRun)
		diags = diags.Add(err)
		result.TotalAffected += affected
		for k, v := range affectedResources {
			result.AffectedResources[k] = v
		}
	}
	return &result, diags
}

func removeProviderStaleData(ctx context.Context, storage execution.Storage, plugin *plugin.Manager, provider string, lastUpdateTime time.Time, dryRun bool) (map[string]int, int, error) {
	providerSchema, err := GetProviderSchema(ctx, plugin, &GetProviderSchemaOptions{Provider: provider})
	if err != nil {
		return nil, 0, err
	}
	var (
		diags             diag.Diagnostics
		logger            = log.With().Bool("dry-run", dryRun).Str("provider", provider).Logger()
		totalAffected     int
		affectedResources = make(map[string]int)
	)

	for r, t := range providerSchema.ResourceTables {
		logger.Debug().Str("table", t.Name).Msg("purging data from table")
		if dryRun {
			affected, err := dryRunPurge(ctx, storage, t, lastUpdateTime)
			if err != nil {
				return nil, 0, diags.Add(err)
			}
			if affected > 0 {
				logger.Info().Str("table", t.Name).Int("affected", affected).Msgf("%d resources will removed from table", affected)
				totalAffected += affected
				affectedResources[r] = affected
			}
			continue
		}
		if err := storage.RemoveStaleData(ctx, t, lastUpdateTime, nil); err != nil {
			diags = diags.Add(diag.NewBaseError(err, diag.DATABASE, diag.WithSeverity(diag.WARNING),
				diag.WithSummary("failed to remove stale data from %s", t.Name),
				diag.WithDetails("table might not exist, is your provider schema version the same as the provider configured?")))
		}
	}
	return affectedResources, totalAffected, diags
}

func dryRunPurge(ctx context.Context, storage execution.Storage, t *schema.Table, lastUpdateTime time.Time) (int, error) {
	q := goqu.Select(goqu.COUNT(goqu.Star())).From(t.Name).
		WithDialect("postgres").Prepared(true).
		Where(goqu.L(`extract(epoch from (cq_meta->>'last_updated')::timestamp)`).Lt(lastUpdateTime.Unix()))
	sql, args, _ := q.ToSQL()
	result, err := storage.Query(ctx, sql, args...)
	if err != nil {
		return 0, diag.NewBaseError(err, diag.DATABASE)
	}
	var affected int
	if err := pgxscan.ScanOne(&affected, result); err != nil {
		return 0, diag.NewBaseError(err, diag.DATABASE)
	}
	return affected, nil
}
