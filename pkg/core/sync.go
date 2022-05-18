package core

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

type SyncState int

const (
	Installed SyncState = iota + 1
	Upgraded
	Downgraded
	NoChange
)

func (s SyncState) String() string {
	switch s {
	case Installed:
		return "Installed"
	case Upgraded:
		return "Upgraded"
	case Downgraded:
		return "Downgraded"
	case NoChange:
		return "NoChange"
	}
	return "Unknown"
}

type SyncResult struct {
	State      SyncState
	OldVersion string
	NewVersion string
}

const dropTableSQL = "DROP TABLE IF EXISTS %s CASCADE"

func Sync(ctx context.Context, storage database.Storage, pm *plugin.Manager, provider registry.Provider) (*SyncResult, diag.Diagnostics) {
	log.Info().Stringer("provider", provider).Msg("syncing provider schema")

	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if diags.HasDiags() {
		return nil, diags
	}

	// TODO: in future more components will want state, so make state more generic and passable via database.Storage
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "sync"), storage.DSN())
	if err != nil {
		return nil, diag.FromError(err, diag.DATABASE)
	}
	defer db.Close()

	provider.Version = s.Version // override any "latest"

	want := state.ProviderFromRegistry(provider)
	if want.ParsedVersion == nil {
		return nil, diag.FromError(fmt.Errorf("failing provider with invalid version %q", provider.Version), diag.INTERNAL)
	}

	sta, err := state.NewMigratedClient(ctx, storage.DSN())
	if err != nil {
		return nil, diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	cur, err := sta.GetProvider(ctx, provider)
	if err != nil {
		return nil, diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	res := &SyncResult{
		NewVersion: provider.Version,
	}
	if cur != nil {
		res.OldVersion = cur.Version
	}

	switch {
	case cur == nil || cur.ParsedVersion == nil: // New install (or older provider)
		log.Debug().Stringer("provider", provider).Str("version", provider.Version).Msg("installing provider schema")
		res.State = Installed
	case cur.ParsedVersion.Equal(want.ParsedVersion): // Same version
		res.State = NoChange
	case cur.ParsedVersion.LessThan(want.ParsedVersion): // Upgrade
		log.Debug().Stringer("provider", provider).Str("version", provider.Version).Msg("upgrading provider schema")
		res.State = Upgraded
	case cur.ParsedVersion.GreaterThan(want.ParsedVersion): // Downgrade
		log.Debug().Stringer("provider", provider).Str("version", provider.Version).Msg("downgrading provider schema")
		res.State = Downgraded
	default:
		return nil, diag.FromError(fmt.Errorf("sync: unhandled case"), diag.INTERNAL)
	}

	if res.State != NoChange {
		if err := syncTables(ctx, sta, cur, want, s.ResourceTables); err != nil {
			return nil, diag.FromError(err, diag.INTERNAL)
		}
	}

	log.Debug().Stringer("provider", provider).Stringer("state", res.State).Msg("provider sync complete")
	return res, diags
}

func Drop(ctx context.Context, storage database.Storage, pm *plugin.Manager, provider registry.Provider) diag.Diagnostics {
	log.Warn().Stringer("provider", provider).Msg("dropping provider schema")
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if diags.HasDiags() {
		return diags
	}

	sta, err := state.NewMigratedClient(ctx, storage.DSN())
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	tx, err := sta.ProviderSync(ctx)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	if err := dropProvider(ctx, tx, provider, resourceTableNames(s.ResourceTables)); err != nil {
		return diag.FromError(err, diag.DATABASE, diag.WithSummary("failed to drop provider"))
	}

	if err := tx.UninstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	if err := tx.Commit(ctx); err != nil {
		return diag.FromError(err, diag.DATABASE)
	}

	return nil
}

func syncTables(ctx context.Context, sta *state.Client, cur, want *state.Provider, resourceTables map[string]*schema.Table) error {
	want.Tables = resourceTableNames(resourceTables)
	want.Signatures = resourceSignatures(resourceTables)

	var curTables map[string][]string

	if cur != nil && len(cur.Tables) > 0 { // Old provider with known, valid data
		curTables = cur.Tables
	} else {
		curTables = want.Tables // Fallback to installed provider tables
	}

	tx, err := sta.ProviderSync(ctx)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	// TODO: compare signatures, keep unchanged tables, drop extra tables

	if err := dropProvider(ctx, tx, want.Registry(), curTables); err != nil {
		return diag.FromError(fmt.Errorf("drop provider failed: %w", err), diag.INTERNAL)
	}

	if err := installProvider(ctx, tx, resourceTables); err != nil {
		return diag.FromError(fmt.Errorf("install provider failed: %w", err), diag.INTERNAL)
	}

	if err := tx.UninstallProvider(ctx, want.Registry()); err != nil {
		return diag.FromError(fmt.Errorf("uninstall provider failed: %w", err), diag.INTERNAL)
	}

	if err := tx.InstallProvider(ctx, want); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	if err := tx.Commit(ctx); err != nil {
		return diag.FromError(err, diag.DATABASE)
	}

	return nil
}

func dropProvider(ctx context.Context, db execution.QueryExecer, provider registry.Provider, tableNames map[string][]string) error {
	q := fmt.Sprintf(dropTableSQL, strconv.Quote(fmt.Sprintf("%s_%s_schema_migrations", provider.Source, provider.Name)))
	if err := db.Exec(ctx, q); err != nil {
		return err
	}
	for name, tables := range tableNames {
		log.Debug().Str("resource", name).Str("provider", provider.Name).Msg("deleting table and all relations")
		for _, t := range tables {
			if err := db.Exec(ctx, fmt.Sprintf(dropTableSQL, strconv.Quote(t))); err != nil {
				return err
			}
		}
	}

	return nil
}

func installProvider(ctx context.Context, db execution.QueryExecer, resourceTables map[string]*schema.Table) error {
	for _, t := range sort.StringSlice(funk.Keys(resourceTables).([]string)) {
		up, err := migration.CreateTableDefinitions(ctx, schema.PostgresDialect{}, resourceTables[t], nil)
		if err != nil {
			return diag.FromError(err, diag.INTERNAL, diag.WithSummary("failed to get create table definition"), diag.WithResourceName(t))
		}
		for _, sql := range up {
			if err := db.Exec(ctx, sql); err != nil {
				return diag.FromError(err, diag.INTERNAL, diag.WithSummary("error creating table"), diag.WithResourceName(t))
			}
		}
	}

	return nil
}

func resourceTableNames(resourceTables map[string]*schema.Table) map[string][]string {
	ret := make(map[string][]string, len(resourceTables))
	for k, t := range resourceTables {
		ret[k] = t.TableNames()
	}
	return ret
}

func resourceSignatures(resourceTables map[string]*schema.Table) map[string]string {
	ret := make(map[string]string, len(resourceTables))
	for k, t := range resourceTables {
		ret[k] = t.Signature()
	}
	return ret
}
