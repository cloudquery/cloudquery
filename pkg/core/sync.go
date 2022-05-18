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

	curTables := s.ResourceTables

	if cur != nil && cur.ParsedVersion != nil { // Old provider with known, valid version
		// Make sure we have it
		// _, diags := Download(ctx, pm, &DownloadOptions{Providers: []registry.Provider{cur.Registry()}, NoVerify: viper.GetBool("no-verify")})

		oldSchema, dd := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: cur.Registry()})
		// diags = diags.Add(dd)
		if dd.HasErrors() {
			log.Warn().Stringer("provider", cur.Registry()).Msg("failed to acquire current version")
		} else {
			curTables = oldSchema.ResourceTables
		}
	}

	// TODO run inside TX

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
		if err := dropProvider(ctx, storage.DSN(), db, provider, curTables); err != nil {
			return nil, diag.FromError(fmt.Errorf("drop provider failed: %w", err), diag.INTERNAL)
		}
		if err := installProvider(ctx, storage.DSN(), db, want, s.ResourceTables); err != nil {
			return nil, diag.FromError(fmt.Errorf("install provider failed: %w", err), diag.INTERNAL)
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
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "drop-provider"), storage.DSN())
	if err != nil {
		return diag.FromError(err, diag.DATABASE)
	}
	defer db.Close()

	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	if err := dropProvider(ctx, storage.DSN(), db, provider, s.ResourceTables); err != nil {
		return diag.FromError(err, diag.DATABASE, diag.WithSummary("failed to drop provider"))
	}
	return nil
}

func dropProvider(ctx context.Context, dsn string, db execution.QueryExecer, provider registry.Provider, resourceTables map[string]*schema.Table) error {
	q := fmt.Sprintf(dropTableSQL, strconv.Quote(fmt.Sprintf("%s_%s_schema_migrations", provider.Source, provider.Name)))
	if err := db.Exec(ctx, q); err != nil {
		return err
	}
	for name, table := range resourceTables {
		log.Debug().Str("table", name).Str("provider", provider.Name).Msg("deleting table and all relations")
		if err := dropTables(ctx, db, table); err != nil {
			return err
		}
	}

	sta, err := state.NewMigratedClient(ctx, dsn)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	if err := sta.UninstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	return nil
}

func dropTables(ctx context.Context, db execution.QueryExecer, table *schema.Table) error {
	if err := db.Exec(ctx, fmt.Sprintf(dropTableSQL, strconv.Quote(table.Name))); err != nil {
		return err
	}
	for _, rel := range table.Relations {
		if err := dropTables(ctx, db, rel); err != nil {
			return err
		}
	}
	return nil
}

func installProvider(ctx context.Context, dsn string, db execution.QueryExecer, provider *state.Provider, resourceTables map[string]*schema.Table) error {
	logger := logging.NewZHcLog(&log.Logger, "sync-install")

	tc := migration.NewTableCreator(logger, schema.PostgresDialect{})

	for _, t := range sort.StringSlice(funk.Keys(resourceTables).([]string)) {
		up, _, err := tc.CreateTableDefinitions(ctx, resourceTables[t], nil)
		if err != nil {
			return diag.FromError(err, diag.INTERNAL, diag.WithSummary("failed to get create table definition"), diag.WithResourceName(t))
		}
		for _, sql := range up {
			if err := db.Exec(ctx, sql); err != nil {
				return diag.FromError(err, diag.INTERNAL, diag.WithSummary("error creating table"), diag.WithResourceName(t))
			}
		}
	}

	sta, err := state.NewMigratedClient(ctx, dsn)
	if err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}
	defer sta.Close()

	if err := sta.InstallProvider(ctx, provider); err != nil {
		return diag.FromError(fmt.Errorf("state failed: %w", err), diag.INTERNAL)
	}

	return nil
}
