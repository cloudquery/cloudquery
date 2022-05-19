package core

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
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

func Sync(ctx context.Context, sta *state.Client, pm *plugin.Manager, provider registry.Provider) (*SyncResult, diag.Diagnostics) {
	log.Info().Stringer("provider", provider).Msg("syncing provider schema")

	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if diags.HasDiags() {
		return nil, diags
	}

	provider.Version = s.Version // override any "latest"

	want := state.ProviderFromRegistry(provider)
	if want.ParsedVersion == nil {
		return nil, diag.FromError(fmt.Errorf("failing provider with invalid version %q", provider.Version), diag.INTERNAL)
	}

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
		diags = diags.Add(syncTables(ctx, sta, cur, want, s.ResourceTables))
	}

	log.Debug().Stringer("provider", provider).Stringer("state", res.State).Uint64("errors", diags.Errors()).Msg("provider sync complete")
	return res, diags
}

func Drop(ctx context.Context, sta *state.Client, pm *plugin.Manager, provider registry.Provider) diag.Diagnostics {
	log.Warn().Stringer("provider", provider).Msg("dropping provider schema")
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if diags.HasDiags() {
		return diags
	}

	tx, err := sta.ProviderSync(ctx)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	if diags := dropProviderTables(ctx, tx, provider, resourceTableNames(s.ResourceTables), nil, nil); diags.HasErrors() {
		return diags
	}

	if err := tx.UninstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	if err := tx.Commit(ctx); err != nil {
		return diag.FromError(err, diag.DATABASE)
	}

	return nil
}

func syncTables(ctx context.Context, sta *state.Client, cur, want *state.Provider, resourceTables map[string]*schema.Table) diag.Diagnostics {
	want.Tables = resourceTableNames(resourceTables)
	want.Signatures = resourceSignatures(resourceTables)

	var (
		curTables     map[string][]string
		curSignatures map[string]string
	)

	if cur != nil && len(cur.Tables) > 0 { // Old provider with known, valid data
		curTables = cur.Tables
		curSignatures = cur.Signatures
	} else {
		curTables = want.Tables // Fallback to installed provider tables
	}

	tx, err := sta.ProviderSync(ctx)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	// If a single SQL fails, all following commands also fail with "current transaction is aborted"
	if diags := dropProviderTables(ctx, tx, want.Registry(), curTables, curSignatures, want.Signatures); diags.HasErrors() {
		return diags
	}
	if diags := createProviderTables(ctx, tx, resourceTables); diags.HasErrors() {
		return diags
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

func dropProviderTables(ctx context.Context, db execution.QueryExecer, provider registry.Provider, tableNames map[string][]string, curSignatures, wantSignatures map[string]string) diag.Diagnostics {
	{
		migTable := fmt.Sprintf("%s_%s_schema_migrations", provider.Source, provider.Name)
		q := fmt.Sprintf(dropTableSQL, strconv.Quote(migTable))
		if err := db.Exec(ctx, q); err != nil {
			return diag.FromError(err, diag.DATABASE, diag.WithSummary("drop table failed"), diag.WithResourceName(migTable))
		}
	}

	for name, tables := range tableNames {
		if curSignatures != nil && wantSignatures != nil && curSignatures[name] != "" && wantSignatures[name] == curSignatures[name] {
			log.Debug().Str("resource", name).Str("provider", provider.Name).Msg("keeping tables and all data")
			continue
		}

		log.Debug().Str("resource", name).Str("provider", provider.Name).Msg("deleting tables and all relations")
		for _, t := range tables {
			if err := db.Exec(ctx, fmt.Sprintf(dropTableSQL, strconv.Quote(t))); err != nil {
				return diag.FromError(err, diag.DATABASE, diag.WithSummary("drop table failed"), diag.WithResourceName(t))
			}
		}
	}

	return nil
}

func createProviderTables(ctx context.Context, db execution.QueryExecer, resourceTables map[string]*schema.Table) diag.Diagnostics {
	var diags diag.Diagnostics

	// We didn't filter which tables we already have in the DB (and didn't drop due to signature matches) and trust that all CREATE TABLE statements will have IF NOT EXISTS
	for _, t := range sort.StringSlice(funk.Keys(resourceTables).([]string)) {
		up, err := migration.CreateTableDefinitions(ctx, schema.PostgresDialect{}, resourceTables[t], nil)
		if err != nil {
			diags = diags.Add(diag.FromError(err, diag.INTERNAL, diag.WithSummary("failed to get create table definition"), diag.WithResourceName(t)))
			continue
		}
		for _, sql := range up {
			if err := db.Exec(ctx, sql); err != nil {
				return diags.Add(diag.FromError(err, diag.DATABASE, diag.WithSummary("error creating table"), diag.WithResourceName(t)))
			}
		}
	}

	return diags
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
		ret[k] = t.Signature(schema.PostgresDialect{})
	}
	return ret
}
