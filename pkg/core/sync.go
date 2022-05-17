package core

import (
	"context"
	"fmt"
	"sort"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

type SyncState int

const (
	Upgraded SyncState = iota + 1
	Downgraded
	NoChange
)

func (s SyncState) String() string {
	switch s {
	case Upgraded:
		return "Upgraded"
	case Downgraded:
		return "Downgraded"
	case NoChange:
		return "NoChange"
	}
	return "Unknown"
}

type SyncOptions struct {
	Provider registry.Provider
	// DownloadLatest specifies whether the latest provider should be downloaded before sync is called.
	DownloadLatest bool
}

type SyncResult struct {
	State      SyncState
	OldVersion string
	NewVersion string
}

func Sync(ctx context.Context, storage database.Storage, pm *plugin.Manager, opts *SyncOptions) (*SyncResult, diag.Diagnostics) {
	log.Info().Stringer("provider", opts.Provider).Msg("syncing provider schema")
	if opts.DownloadLatest {
		if _, diags := Download(ctx, pm, &DownloadOptions{
			[]registry.Provider{{Name: opts.Provider.Name, Version: registry.LatestVersion, Source: opts.Provider.Source}}, false}); diags.HasErrors() {
			return nil, diags
		}
	}

	rp := registry.Provider{Name: opts.Provider.Name, Version: registry.LatestVersion, Source: opts.Provider.Source}

	// always use latest available provider available
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: rp})
	if diags.HasDiags() {
		return nil, diags
	}

	// TODO: in future more components will want state, so make state more generic and passable via database.Storage
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "sync"), storage.DSN())
	if err != nil {
		return nil, diag.FromError(err, diag.DATABASE)
	}
	defer db.Close()

	rp.Version = s.Version

	want := state.ProviderFromRegistry(rp)
	if want.ParsedVersion == nil {
		return nil, diag.FromError(fmt.Errorf("failing provider with invalid version %q", rp.Version), diag.INTERNAL)
	}

	sta, err := state.NewMigratedClient(ctx, storage.DSN(), logging.NewZHcLog(&log.Logger, "sync"))
	if err != nil {
		return nil, diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	cur, err := sta.GetProvider(ctx, rp)
	if err != nil {
		return nil, diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	res := &SyncResult{
		NewVersion: opts.Provider.Version,
	}
	if cur != nil {
		res.OldVersion = cur.Version
	}

	// TODO run inside TX

	switch {
	case cur == nil || cur.ParsedVersion == nil: // New install (or older provider)
		log.Debug().Stringer("provider", opts.Provider).Str("version", opts.Provider.Version).Msg("installing provider schema")
		res.State = Upgraded
	case cur.ParsedVersion.Equal(want.ParsedVersion): // Same version
		res.State = NoChange
	case cur.ParsedVersion.LessThan(want.ParsedVersion): // Upgrade
		log.Debug().Stringer("provider", opts.Provider).Str("version", opts.Provider.Version).Msg("upgrading provider schema")
		res.State = Upgraded
	case cur.ParsedVersion.GreaterThan(want.ParsedVersion): // Downgrade
		log.Debug().Stringer("provider", opts.Provider).Str("version", opts.Provider.Version).Msg("downgrading provider schema")
		res.State = Downgraded
	default:
		return nil, diag.FromError(fmt.Errorf("sync: unhandled case"), diag.INTERNAL)
	}

	if res.State == Upgraded || res.State == Downgraded {
		if err := dropProvider(ctx, storage.DSN(), rp, s.ResourceTables); err != nil {
			return nil, diag.FromError(fmt.Errorf("drop provider failed: %w", err), diag.INTERNAL)
		}
		if err := installProvider(ctx, storage.DSN(), db, want, s.ResourceTables); err != nil {
			return nil, diag.FromError(fmt.Errorf("install provider failed: %w", err), diag.INTERNAL)
		}
	}

	log.Debug().Stringer("provider", opts.Provider).Stringer("state", res.State).Msg("provider sync complete")
	return res, nil
}

func Drop(ctx context.Context, storage database.Storage, pm *plugin.Manager, provider registry.Provider) diag.Diagnostics {
	log.Warn().Stringer("provider", provider).Msg("dropping provider schema")
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if diags.HasDiags() {
		return diags
	}
	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	err := dropProvider(ctx, storage.DSN(), provider, s.ResourceTables)
	if err != nil {
		return diag.FromError(err, diag.DATABASE, diag.WithSummary("failed to drop provider"))
	}
	return nil
}

func dropProvider(ctx context.Context, dsn string, provider registry.Provider, resourceTables map[string]*schema.Table) error {
	dType, _, err := sdkdb.ParseDialectDSN(dsn)
	if err != nil {
		return err
	}

	m, err := migrator.New(logging.NewZHcLog(&log.Logger, "migrator"), dType, nil, dsn, fmt.Sprintf("%s_%s", provider.Source, provider.Name))
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close migrator connection")
		}
	}()

	if err := m.DropProvider(ctx, resourceTables); err != nil {
		return err
	}

	sta, err := state.NewMigratedClient(ctx, dsn, logging.NewZHcLog(&log.Logger, "sync"))
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	if err := sta.UninstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	return nil
}

func installProvider(ctx context.Context, dsn string, db execution.QueryExecer, provider *state.Provider, resourceTables map[string]*schema.Table) error {
	logger := logging.NewZHcLog(&log.Logger, "sync-install")

	tc := migration.NewTableCreator(logger, schema.PostgresDialect{})

	for _, t := range sort.StringSlice(funk.Keys(resourceTables).([]string)) {
		up, _, err := tc.CreateTableDefinitions(ctx, resourceTables[t], nil)
		if err != nil {
			return diag.FromError(fmt.Errorf("CreateTableDefinitions failed for %s: %w", t, err), diag.INTERNAL)
		}
		for _, sql := range up {
			if err := db.Exec(ctx, sql); err != nil {
				return diag.FromError(fmt.Errorf("error creating table %s: %w", t, err), diag.INTERNAL)
			}
		}
	}

	sta, err := state.NewMigratedClient(ctx, dsn, logging.NewZHcLog(&log.Logger, "sync"))
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	if err := sta.InstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	return nil
}
