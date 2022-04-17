package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
)

type SyncState int

const (
	Upgraded SyncState = iota + 1
	Downgraded
	NoChange
)

var (
	ErrMigrationsNotSupported = errors.New("provider doesn't support migrations")
)

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
	if opts.DownloadLatest {
		if _, diags := Download(ctx, pm, &DownloadOptions{
			[]registry.Provider{{Name: opts.Provider.Name, Version: registry.LatestVersion, Source: opts.Provider.Source}}, false}); diags.HasErrors() {
			return nil, diags
		}
	}
	// always use latest available provider available
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: registry.Provider{Name: opts.Provider.Name, Version: registry.LatestVersion, Source: opts.Provider.Source}})
	if len(diags) > 0 {
		return nil, diags
	}
	if s.Migrations == nil {
		return nil, diag.FromError(ErrMigrationsNotSupported, diag.DATABASE, diag.WithSeverity(diag.IGNORE))
	}
	// create migrator
	m, err := newMigrator(ctx, storage, s.Migrations, opts.Provider)
	if err != nil {
		return nil, diag.FromError(err, diag.DATABASE)
	}
	defer func() {
		if err := m.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close migrator connection")
		}
	}()

	// get version from current schema
	ver, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	if dirty {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	currentVersion, err := version.NewVersion(ver)

	var providerVersion *version.Version
	if opts.Provider.Version != registry.LatestVersion {
		providerVersion, err = version.NewVersion(opts.Provider.Version)
		if err != nil {
			return nil, diag.FromError(err, diag.INTERNAL)
		}
	}

	state := NoChange
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	if opts.Provider.Version == registry.LatestVersion || providerVersion.GreaterThan(currentVersion) {
		if err := m.UpgradeProvider(opts.Provider.Version); err != nil && err != migrate.ErrNoChange {
			return nil, diag.FromError(err, diag.DATABASE)
		}
		state = Upgraded
	} else if providerVersion.LessThan(currentVersion) {
		if err := m.DowngradeProvider(opts.Provider.Version); err != nil && err != migrate.ErrNoChange {
			return nil, diag.FromError(err, diag.DATABASE)
		}
		state = Downgraded
	}
	return &SyncResult{
		State:      state,
		OldVersion: currentVersion.Original(),
		NewVersion: opts.Provider.Version,
	}, nil
}

func Drop(ctx context.Context, storage database.Storage, pm *plugin.Manager, provider registry.Provider) diag.Diagnostics {
	s, diags := GetProviderSchema(ctx, pm, &GetProviderSchemaOptions{Provider: provider})
	if len(diags) > 0 {
		return diags
	}
	m, err := newMigrator(ctx, storage, s.Migrations, provider)
	if err != nil {
		return nil // TODO: return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close migrator connection")
		}
	}()

	log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("dropping provider tables")
	if err := m.DropProvider(ctx, s.ResourceTables); err != nil {
		return nil // TODO: return err
	}
	return nil
}

func newMigrator(ctx context.Context, storage database.Storage, migrations map[string]map[string][]byte, provider registry.Provider) (*migrator.Migrator, error) {
	dsn, err := storage.DialectExecutor().Setup(ctx)
	if err != nil {
		return nil, fmt.Errorf("dialectExecutor.Setup: %w", err)
	}

	dType, _, err := sdkdb.ParseDialectDSN(storage.DSN())
	if err != nil {
		return nil, err
	}

	m, err := migrator.New(logging.NewZHcLog(&log.Logger, "migrator"), dType, migrations, dsn,
		fmt.Sprintf("%s_%s", provider.Source, provider.Name),
		migrator.WithPreHook(storage.DialectExecutor().Prepare),
		migrator.WithPostHook(storage.DialectExecutor().Finalize))
	if err != nil {
		return nil, err
	}
	return m, err
}
