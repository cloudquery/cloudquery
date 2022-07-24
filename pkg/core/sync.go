package core

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

type SyncState int

type SyncResult struct {
	State      SyncState
	OldVersion string
	NewVersion string
}

const dropTableSQL = "DROP TABLE IF EXISTS %s CASCADE"

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

func Sync(ctx context.Context, sta *state.Client, pm *plugin.Manager, provider registry.Provider) (*SyncResult, error) {
	log.Info().Stringer("provider", provider).Msg("syncing provider schema")

	s, err := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if err != nil {
		return nil, err
	}

	provider.Version = s.Version // override any "latest"

	want := state.ProviderFromRegistry(provider)
	if want.ParsedVersion == nil && s.Unmanaged {
		provider.Version = "v0.0.0"
		want = state.ProviderFromRegistry(provider)
	}

	if want.ParsedVersion == nil {
		return nil, fmt.Errorf("expected provider version from schema to be a valid semantic version but got %q. If you're trying to debug a provider %s", s.Version, ui.Link("see our docs", "https://docs.cloudquery.io/docs/developers/debugging"))
	}

	cur, err := sta.GetProvider(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("state failed: %w", err)
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
		return nil, fmt.Errorf("sync: unhandled case")
	}

	if res.State != NoChange {
		// TODO
	}

	log.Debug().Stringer("provider", provider).Stringer("state", res.State).Uint64("errors", 0).Msg("provider sync complete")
	return res, nil
}

func Drop(ctx context.Context, sta *state.Client, pm *plugin.Manager, provider registry.Provider) error {
	log.Warn().Stringer("provider", provider).Msg("dropping provider schema")
	s, err := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if err != nil {
		return err
	}

	tx, err := sta.ProviderSync(ctx)
	if err != nil {
		return fmt.Errorf("state failed: %w", err)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	if err := dropProviderTables(ctx, tx, provider, resourceTableNames(s.ResourceTables), nil, nil); err != nil {
		return err
	}

	if err := tx.UninstallProvider(ctx, provider); err != nil {
		return fmt.Errorf("state failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}

func syncTables(ctx context.Context, sta *state.Client, cur, want *state.Provider, resourceTables map[string]*schema.Table) error {
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
		return fmt.Errorf("state failed: %w", err)
	}
	defer tx.Rollback(ctx) // nolint:errcheck

	// If a single SQL fails, all following commands also fail with "current transaction is aborted"
	if err := dropProviderTables(ctx, tx, want.Registry(), curTables, curSignatures, want.Signatures); err != nil {
		return err
	}
	if err := createProviderTables(ctx, tx, resourceTables); err != nil {
		return err
	}

	if err := tx.UninstallProvider(ctx, want.Registry()); err != nil {
		return fmt.Errorf("uninstall failed: %w", err)
	}

	if err := tx.InstallProvider(ctx, want); err != nil {
		return fmt.Errorf("state failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}

func dropProviderTables(ctx context.Context, db execution.QueryExecer, provider registry.Provider, tableNames map[string][]string, curSignatures, wantSignatures map[string]string) error {
	{
		migTable := fmt.Sprintf("%s_%s_schema_migrations", provider.Source, provider.Name)
		q := fmt.Sprintf(dropTableSQL, strconv.Quote(migTable))
		if err := db.Exec(ctx, q); err != nil {
			return fmt.Errorf("failed to drop %s: %w", migTable, err)
		}
	}

	force := viper.GetBool("force-drop")

	for name, tables := range tableNames {
		if curSignatures != nil && wantSignatures != nil && curSignatures[name] != "" && wantSignatures[name] == curSignatures[name] {
			log.Debug().Str("resource", name).Str("provider", provider.Name).Msg("keeping tables and all data")
			continue
		}

		log.Debug().Str("resource", name).Str("provider", provider.Name).Msg("deleting tables and all relations")
		for _, t := range tables {
			if !force {
				if err := safeToDropTable(ctx, db, name, t); err != nil {
					return err
				}
			}
			if err := db.Exec(ctx, fmt.Sprintf(dropTableSQL, strconv.Quote(t))); err != nil {
				return err
			}
		}
	}

	return nil
}

func createProviderTables(ctx context.Context, db execution.QueryExecer, resourceTables map[string]*schema.Table) error {
	// We didn't filter which tables we already have in the DB (and didn't drop due to signature matches) and trust that all CREATE TABLE statements will have IF NOT EXISTS
	for _, t := range sort.StringSlice(funk.Keys(resourceTables).([]string)) {
		up, err := migration.CreateTableDefinitions(ctx, schema.PostgresDialect{}, resourceTables[t], nil)
		if err != nil {
			log.Err(err)
			continue
		}
		for _, sql := range up {
			if err := db.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to create table %s: %w", t, err)
			}
		}
	}

	return nil
}

func safeToDropTable(ctx context.Context, db execution.QueryExecer, resName, table string) error {
	rows, err := db.Query(ctx, `
SELECT DISTINCT CONCAT(dependent_ns.nspname, '.', dependent_view.relname) as dependent_view
FROM pg_depend
JOIN pg_rewrite ON pg_depend.objid = pg_rewrite.oid
JOIN pg_class as dependent_view ON pg_rewrite.ev_class = dependent_view.oid
JOIN pg_class as source_table ON pg_depend.refobjid = source_table.oid
JOIN pg_attribute ON pg_depend.refobjid = pg_attribute.attrelid AND pg_depend.refobjsubid = pg_attribute.attnum
JOIN pg_namespace dependent_ns ON dependent_ns.oid = dependent_view.relnamespace
JOIN pg_namespace source_ns ON source_ns.oid = source_table.relnamespace
WHERE source_ns.nspname = current_schema() AND source_table.relname = $1 AND pg_attribute.attnum > 0 ORDER BY 1
`, table)
	if err != nil {
		return fmt.Errorf("failed to query dependent views: %w", err)
	}

	var dependentViews []string
	if err := pgxscan.ScanAll(&dependentViews, rows); err != nil {
		return err
	}

	if len(dependentViews) > 0 {
		return fmt.Errorf("One or more views depend on table %q, which is about to be dropped. Run with --force-drop to override this message, which will drop ALL dependent views.")
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
		ret[k] = t.Signature(schema.PostgresDialect{})
	}
	return ret
}
