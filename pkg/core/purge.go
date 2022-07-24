package core

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/rs/zerolog/log"
)

type PurgeProviderDataOptions struct {
	// Providers to purge data from, the provider name should be the plugin name
	Providers []registry.Provider
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
func PurgeProviderData(ctx context.Context, storage database.Storage, manager *plugin.Manager, opts *PurgeProviderDataOptions) (*PurgeProviderDataResult, error) {
	if len(opts.Providers) == 0 {
		return nil, fmt.Errorf("no providers specified")
	}
	log.Info().Interface("providers", opts.Providers).Bool("dry-run", opts.DryRun).Msg("purging stale data for providers")
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "database"), storage.DSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var (
		result = PurgeProviderDataResult{
			TotalAffected:     0,
			AffectedResources: make(map[string]int),
		}
	)

	lastUpdateTime := time.Now().UTC().Add(-opts.LastUpdate)
	for _, p := range opts.Providers {
		log.Debug().Stringer("provider", p).TimeDiff("since", lastUpdateTime, time.Now().UTC()).Msg("cleaning stale data for provider")
		affectedResources, affected, err := removeProviderStaleData(ctx, db, manager, p, lastUpdateTime, opts.DryRun)
		if err != nil {
			return nil, err
		}
		result.TotalAffected += affected
		for k, v := range affectedResources {
			result.AffectedResources[k] = v
		}
	}
	return &result, nil
}

func removeProviderStaleData(ctx context.Context, storage execution.Storage, manager *plugin.Manager, provider registry.Provider, lastUpdateTime time.Time, dryRun bool) (map[string]int, int, error) {
	providerSchema, err := GetProviderSchema(ctx, manager, &GetProviderSchemaOptions{Provider: provider})
	if err != nil {
		return nil, 0, err
	}
	var (
		logger            = log.With().Bool("dry-run", dryRun).Stringer("provider", provider).Logger()
		totalAffected     int
		affectedResources = make(map[string]int)
	)

	for r, t := range providerSchema.ResourceTables {
		logger.Debug().Str("table", t.Name).Msg("purging data from table")
		if dryRun {
			affected, err := dryRunPurge(ctx, storage, t, lastUpdateTime)
			if err != nil {
				return nil, 0, err
			}
			if affected > 0 {
				logger.Info().Str("table", t.Name).Int("affected", affected).Msgf("%d resources will removed from table", affected)
				totalAffected += affected
				affectedResources[r] = affected
			}
			continue
		}
		if err := storage.RemoveStaleData(ctx, t, lastUpdateTime, nil); err != nil {
			return nil, 0, err
		}
	}
	return affectedResources, totalAffected, err
}

func dryRunPurge(ctx context.Context, storage execution.Storage, t *schema.Table, lastUpdateTime time.Time) (int, error) {
	q := goqu.Select(goqu.COUNT(goqu.Star())).From(t.Name).
		WithDialect("postgres").Prepared(true).
		Where(goqu.L(`extract(epoch from (cq_meta->>'last_updated')::timestamp)`).Lt(lastUpdateTime.Unix()))
	sql, args, _ := q.ToSQL()
	result, err := storage.Query(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	var affected int
	if err := pgxscan.ScanOne(&affected, result); err != nil {
		return 0, err
	}
	return affected, nil
}
