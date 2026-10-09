package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	discoveryPb "github.com/cloudquery/plugin-pb-go/pb/discovery/v1"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeSourceClient struct {
	pluginPb.PluginClient
	tablesWithoutConnection  schema.Tables
	tablesWithConnection     schema.Tables
	initErrWithoutConnection error
	initErrWithConnection    error
	errWithoutConnection     error
	errWithConnection        error

	noConnection bool
	initCalls    []bool
	syncCalled   bool
}

func (f *fakeSourceClient) Init(_ context.Context, req *pluginPb.Init_Request, _ ...grpc.CallOption) (*pluginPb.Init_Response, error) {
	f.noConnection = req.NoConnection
	f.initCalls = append(f.initCalls, req.NoConnection)
	if req.NoConnection && f.initErrWithoutConnection != nil {
		return nil, f.initErrWithoutConnection
	}
	if !req.NoConnection && f.initErrWithConnection != nil {
		return nil, f.initErrWithConnection
	}
	return &pluginPb.Init_Response{}, nil
}

func (f *fakeSourceClient) GetTables(_ context.Context, req *pluginPb.GetTables_Request, _ ...grpc.CallOption) (*pluginPb.GetTables_Response, error) {
	if len(req.Tables) != 1 || req.Tables[0] != "*" {
		return nil, errors.New("expected all tables to be requested")
	}
	tables := f.tablesWithoutConnection
	if f.noConnection && f.errWithoutConnection != nil {
		return nil, f.errWithoutConnection
	}
	if !f.noConnection {
		if f.errWithConnection != nil {
			return nil, f.errWithConnection
		}
		tables = f.tablesWithConnection
	}
	b, err := pluginPb.SchemasToBytes(tables.ToArrowSchemas())
	if err != nil {
		return nil, err
	}
	return &pluginPb.GetTables_Response{Tables: b}, nil
}

func (f *fakeSourceClient) Sync(context.Context, *pluginPb.Sync_Request, ...grpc.CallOption) (grpc.ServerStreamingClient[pluginPb.Sync_Response], error) {
	f.syncCalled = true
	return nil, errors.New("sync must not be called")
}

func testTable(name string) *schema.Table {
	return &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "id", Type: arrow.BinaryTypes.String}}}
}

func TestListSourceTables(t *testing.T) {
	cases := []struct {
		name          string
		source        *fakeSourceClient
		wantTables    []string
		wantConnected bool
		wantUnknown   string
		wantInitCalls []bool
	}{
		{
			name:          "static source lists tables without a connection",
			source:        &fakeSourceClient{tablesWithoutConnection: schema.Tables{testTable("static_table")}},
			wantTables:    []string{"static_table"},
			wantInitCalls: []bool{true},
		},
		{
			name:          "dynamic source is re-initialized with a connection",
			source:        &fakeSourceClient{tablesWithConnection: schema.Tables{testTable("dynamic_table")}},
			wantTables:    []string{"dynamic_table"},
			wantConnected: true,
			wantInitCalls: []bool{true, false},
		},
		{
			name:          "source with no tables after connecting is unknown",
			source:        &fakeSourceClient{},
			wantConnected: true,
			wantUnknown:   "source returned no tables with or without a connection",
			wantInitCalls: []bool{true, false},
		},
		{
			name:          "source that errors after connecting is unknown with the reason",
			source:        &fakeSourceClient{errWithConnection: errors.New("tables only discovered during sync")},
			wantConnected: true,
			wantUnknown:   "failed to get tables: tables only discovered during sync",
			wantInitCalls: []bool{true, false},
		},
		{
			name:          "init error without a connection is unknown and does not connect",
			source:        &fakeSourceClient{initErrWithoutConnection: errors.New("invalid spec")},
			wantUnknown:   "failed to init source: invalid spec",
			wantInitCalls: []bool{true},
		},
		{
			name:          "get tables error without a connection is unknown and does not connect",
			source:        &fakeSourceClient{errWithoutConnection: errors.New("listing failed")},
			wantUnknown:   "failed to get tables: listing failed",
			wantInitCalls: []bool{true},
		},
		{
			name:          "init error with a connection is unknown with the reason",
			source:        &fakeSourceClient{initErrWithConnection: errors.New("connection refused")},
			wantConnected: true,
			wantUnknown:   "failed to init source: connection refused",
			wantInitCalls: []bool{true, false},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, registerOnce())

			got := listSourceTables(context.Background(), tc.source, map[string]any{"key": "value"})

			require.Equal(t, tc.wantTables, tableNames(got.Tables))
			require.Equal(t, tc.wantConnected, got.Connected)
			require.Equal(t, tc.wantUnknown, got.UnknownReason)
			require.Equal(t, tc.wantInitCalls, tc.source.initCalls)
			require.False(t, tc.source.syncCalled)
		})
	}
}

func TestLoadUpgradeSourceTables(t *testing.T) {
	sourceSpec := specs.Source{
		Metadata: specs.Metadata{
			Name:     "test",
			Path:     "cloudquery/test",
			Version:  "v4.5.1",
			Registry: specs.RegistryCloudQuery,
		},
	}

	ctx := context.Background()
	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, []*specs.Source{&sourceSpec}, nil, nil)
	require.NoError(t, err)
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	require.NoError(t, err)

	from, to, err := loadUpgradeSourceTables(ctx, sourceSpec, upgradeTarget{Version: "v4.7.0"}, managedplugin.NewClient,
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(authToken.Value),
		managedplugin.WithTeamName(teamName),
		managedplugin.WithDirectory(t.TempDir()),
	)
	require.NoError(t, err)

	for _, got := range []upgradeSourceTables{from, to} {
		require.Empty(t, got.UnknownReason)
		require.False(t, got.Connected)
		require.Contains(t, tableNames(got.Tables.FlattenTables()), "test_some_table")
	}
	require.Equal(t, "v4.5.1", from.Version)
	require.Equal(t, "v4.7.0", to.Version)
	require.Equal(t, upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/test"}, from.Origin)
	require.Equal(t, from.Origin, to.Origin)
}

type fakeSourceServer struct {
	pluginPb.UnimplementedPluginServer
	discoveryPb.UnimplementedDiscoveryServer
	name    string
	version string
	tables  schema.Tables
}

func (*fakeSourceServer) GetVersions(context.Context, *discoveryPb.GetVersions_Request) (*discoveryPb.GetVersions_Response, error) {
	return &discoveryPb.GetVersions_Response{Versions: []int32{3}}, nil
}

func (f *fakeSourceServer) GetName(context.Context, *pluginPb.GetName_Request) (*pluginPb.GetName_Response, error) {
	if f.name == "" {
		return nil, status.Error(codes.Unimplemented, "no name")
	}
	return &pluginPb.GetName_Response{Name: f.name}, nil
}

func (f *fakeSourceServer) GetVersion(context.Context, *pluginPb.GetVersion_Request) (*pluginPb.GetVersion_Response, error) {
	if f.version == "" {
		return nil, status.Error(codes.Unimplemented, "no version")
	}
	return &pluginPb.GetVersion_Response{Version: f.version}, nil
}

func (*fakeSourceServer) Init(context.Context, *pluginPb.Init_Request) (*pluginPb.Init_Response, error) {
	return &pluginPb.Init_Response{}, nil
}

func (f *fakeSourceServer) GetTables(context.Context, *pluginPb.GetTables_Request) (*pluginPb.GetTables_Response, error) {
	b, err := pluginPb.SchemasToBytes(f.tables.ToArrowSchemas())
	if err != nil {
		return nil, err
	}
	return &pluginPb.GetTables_Response{Tables: b}, nil
}

func startFakeSourceServer(t *testing.T, source *fakeSourceServer) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	server := grpc.NewServer()
	pluginPb.RegisterPluginServer(server, source)
	discoveryPb.RegisterDiscoveryServer(server, source)
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(server.Stop)
	return listener.Addr().String()
}

type recordingPluginStarter struct {
	fromAddress string
	toAddress   string
	toVersion   string
	toErr       error
	configs     []managedplugin.Config
}

func (r *recordingPluginStarter) newClient(ctx context.Context, typ managedplugin.PluginType, config managedplugin.Config, opts ...managedplugin.Option) (*managedplugin.Client, error) {
	r.configs = append(r.configs, config)
	address := r.fromAddress
	if config.Version == r.toVersion {
		if r.toErr != nil {
			return nil, r.toErr
		}
		address = r.toAddress
	}
	return managedplugin.NewClient(ctx, typ, managedplugin.Config{Name: config.Name, Registry: managedplugin.RegistryGrpc, Path: address}, opts...)
}

func TestLoadUpgradeSourceTablesRegistryDispatch(t *testing.T) {
	fromAddress := startFakeSourceServer(t, &fakeSourceServer{name: "semgrep", version: "v3.0.0", tables: schema.Tables{testTable("semgrep_old")}})
	toAddress := startFakeSourceServer(t, &fakeSourceServer{name: "semgrep", version: "v3.1.0", tables: schema.Tables{testTable("semgrep_new")}})

	cases := []struct {
		name           string
		metadata       specs.Metadata
		toPath         string
		wantFromConfig managedplugin.Config
		wantToConfig   managedplugin.Config
		wantFrom       upgradeSourceOrigin
		wantTo         upgradeSourceOrigin
		wantFromVer    string
	}{
		{
			name:           "cloudquery registry starts the same path at both versions",
			metadata:       specs.Metadata{Name: "semgrep", Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.0.0"},
			wantFromConfig: managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.0.0"},
			wantToConfig:   managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.1.0"},
			wantFrom:       upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"},
			wantTo:         upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"},
			wantFromVer:    "v3.0.0",
		},
		{
			name:           "github registry starts the same path at both versions",
			metadata:       specs.Metadata{Name: "semgrep", Registry: specs.RegistryGitHub, Path: "cloudquery/semgrep", Version: "v3.0.0"},
			wantFromConfig: managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryGithub, Path: "cloudquery/semgrep", Version: "v3.0.0"},
			wantToConfig:   managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryGithub, Path: "cloudquery/semgrep", Version: "v3.1.0"},
			wantFrom:       upgradeSourceOrigin{Registry: specs.RegistryGitHub, Path: "cloudquery/semgrep"},
			wantTo:         upgradeSourceOrigin{Registry: specs.RegistryGitHub, Path: "cloudquery/semgrep"},
			wantFromVer:    "v3.0.0",
		},
		{
			name:           "local registry starts the configured binary and the inferred hub plugin",
			metadata:       specs.Metadata{Name: "semgrep-prod", Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"},
			wantFromConfig: managedplugin.Config{Name: "semgrep-prod", Registry: managedplugin.RegistryLocal, Path: "/opt/plugins/semgrep"},
			wantToConfig:   managedplugin.Config{Name: "semgrep-prod", Registry: managedplugin.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.1.0"},
			wantFrom:       upgradeSourceOrigin{Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"},
			wantTo:         upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"},
			wantFromVer:    "v3.0.0",
		},
		{
			name:           "grpc registry connects to the configured server and the inferred hub plugin",
			metadata:       specs.Metadata{Name: "semgrep", Registry: specs.RegistryGRPC, Path: "localhost:7777"},
			wantFromConfig: managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryGrpc, Path: "localhost:7777"},
			wantToConfig:   managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.1.0"},
			wantFrom:       upgradeSourceOrigin{Registry: specs.RegistryGRPC, Path: "localhost:7777"},
			wantTo:         upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"},
			wantFromVer:    "v3.0.0",
		},
		{
			name:           "docker registry starts the configured image and the hub plugin from --to-path",
			metadata:       specs.Metadata{Name: "semgrep", Registry: specs.RegistryDocker, Path: "ghcr.io/acme/semgrep:v3.0.0", DockerRegistryAuthToken: "secret"},
			toPath:         "acme/semgrep-fork",
			wantFromConfig: managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryDocker, Path: "ghcr.io/acme/semgrep:v3.0.0", DockerAuth: "secret"},
			wantToConfig:   managedplugin.Config{Name: "semgrep", Registry: managedplugin.RegistryCloudQuery, Path: "acme/semgrep-fork", Version: "v3.1.0"},
			wantFrom:       upgradeSourceOrigin{Registry: specs.RegistryDocker, Path: "ghcr.io/acme/semgrep:v3.0.0"},
			wantTo:         upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "acme/semgrep-fork"},
			wantFromVer:    "v3.0.0",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			starter := &recordingPluginStarter{fromAddress: fromAddress, toAddress: toAddress, toVersion: "v3.1.0"}
			sourceSpec := specs.Source{Metadata: tc.metadata, Spec: map[string]any{}}

			from, to, err := loadUpgradeSourceTables(t.Context(), sourceSpec, upgradeTarget{Version: "v3.1.0", Path: tc.toPath}, starter.newClient, managedplugin.WithLogger(log.Logger))

			require.NoError(t, err)
			require.Equal(t, []managedplugin.Config{tc.wantFromConfig, tc.wantToConfig}, starter.configs)
			require.Equal(t, tc.wantFrom, from.Origin)
			require.Equal(t, tc.wantTo, to.Origin)
			require.Equal(t, tc.wantFromVer, from.Version)
			require.Equal(t, "v3.1.0", to.Version)
			require.Equal(t, []string{"semgrep_old"}, tableNames(from.Tables))
			require.Equal(t, []string{"semgrep_new"}, tableNames(to.Tables))
		})
	}
}

func TestLoadUpgradeSourceTablesUnknownReportedVersion(t *testing.T) {
	starter := &recordingPluginStarter{
		fromAddress: startFakeSourceServer(t, &fakeSourceServer{name: "semgrep", tables: schema.Tables{testTable("semgrep_old")}}),
		toAddress:   startFakeSourceServer(t, &fakeSourceServer{name: "semgrep", tables: schema.Tables{testTable("semgrep_new")}}),
		toVersion:   "v3.1.0",
	}
	sourceSpec := specs.Source{Metadata: specs.Metadata{Name: "semgrep", Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"}}

	from, _, err := loadUpgradeSourceTables(t.Context(), sourceSpec, upgradeTarget{Version: "v3.1.0"}, starter.newClient, managedplugin.WithLogger(log.Logger))

	require.NoError(t, err)
	require.Equal(t, upgradeUnknownVersion, from.Version)
}

func TestLoadUpgradeSourceTablesErrors(t *testing.T) {
	namedAddress := startFakeSourceServer(t, &fakeSourceServer{name: "semgrep", tables: schema.Tables{testTable("semgrep_old")}})
	unnamedAddress := startFakeSourceServer(t, &fakeSourceServer{tables: schema.Tables{testTable("semgrep_old")}})
	local := specs.Metadata{Name: "semgrep", Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"}

	cases := []struct {
		name        string
		metadata    specs.Metadata
		toPath      string
		fromAddress string
		toErr       error
		wantErr     string
		wantStarts  int
	}{
		{
			name:     "to-path with the cloudquery registry",
			metadata: specs.Metadata{Name: "semgrep", Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep", Version: "v3.0.0"},
			toPath:   "cloudquery/semgrep",
			wantErr:  "--to-path is supported only for sources from the local, grpc or docker registry, source semgrep uses the cloudquery registry",
		},
		{
			name:     "to-path without a team",
			metadata: local,
			toPath:   "semgrep",
			wantErr:  `invalid --to-path "semgrep". Use <team>/<name>, for example cloudquery/aws`,
		},
		{
			name:     "to-path with too many parts",
			metadata: local,
			toPath:   "cloudquery/semgrep/v3",
			wantErr:  `invalid --to-path "cloudquery/semgrep/v3"`,
		},
		{
			name:        "source does not report its name",
			metadata:    local,
			fromAddress: unnamedAddress,
			wantErr:     "could not infer the CloudQuery Hub path to upgrade source semgrep to, because the source from the local registry did not report its name. Use --to-path <team>/<name> to set the path",
			wantStarts:  1,
		},
		{
			name:        "inferred hub path cannot be started",
			metadata:    local,
			fromAddress: namedAddress,
			toErr:       errors.New("plugin not found"),
			wantErr:     "failed to start source semgrep cloudquery/semgrep@v3.1.0: plugin not found. The CloudQuery Hub path cloudquery/semgrep was inferred from the name the source reports. Use --to-path <team>/<name> to set a different path",
			wantStarts:  2,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			starter := &recordingPluginStarter{fromAddress: tc.fromAddress, toVersion: "v3.1.0", toErr: tc.toErr}

			_, _, err := loadUpgradeSourceTables(t.Context(), specs.Source{Metadata: tc.metadata}, upgradeTarget{Version: "v3.1.0", Path: tc.toPath}, starter.newClient, managedplugin.WithLogger(log.Logger))

			require.ErrorContains(t, err, tc.wantErr)
			require.Len(t, starter.configs, tc.wantStarts)
		})
	}
}

func TestUpgradeAuthSources(t *testing.T) {
	for _, registry := range []specs.Registry{specs.RegistryCloudQuery, specs.RegistryGitHub} {
		sourceSpec := specs.Source{Metadata: specs.Metadata{Name: "semgrep", Registry: registry}}
		require.Equal(t, []*specs.Source{&sourceSpec}, upgradeAuthSources(sourceSpec))
	}
	for _, registry := range []specs.Registry{specs.RegistryLocal, specs.RegistryGRPC, specs.RegistryDocker} {
		sourceSpec := specs.Source{Metadata: specs.Metadata{Name: "semgrep", Registry: registry}}
		sources := upgradeAuthSources(sourceSpec)
		require.Len(t, sources, 2)
		require.Equal(t, specs.RegistryCloudQuery, sources[1].Registry)
	}
}

func tableNames(tables schema.Tables) []string {
	if len(tables) == 0 {
		return nil
	}
	names := make([]string, len(tables))
	for i, table := range tables {
		names[i] = table.Name
	}
	return names
}

func TestSelectUpgradeTables(t *testing.T) {
	parent := testTable("test_parent")
	child := testTable("test_parent_child")
	child.Parent = parent
	parent.Relations = schema.Tables{child}
	from := schema.Tables{testTable("test_removed"), testTable("test_kept"), parent, testTable("test_skipped")}.FlattenTables()
	to := schema.Tables{testTable("test_kept"), testTable("test_added"), testTable("test_skipped")}

	selection, err := selectUpgradeTables(specs.Source{
		Metadata:            specs.Metadata{Name: "test", Version: "v1.0.0"},
		Tables:              []string{"test_removed", "test_kept", "test_added", "test_parent", "test_skipped"},
		SkipTables:          []string{"test_skipped"},
		SkipDependentTables: lo.ToPtr(false),
	}, from, to)
	require.NoError(t, err)

	require.Equal(t, []string{"test_removed", "test_parent", "test_parent_child"}, selection.RemovedTables)
	require.Len(t, selection.Pairs, 2)
	require.Equal(t, "test_kept", selection.Pairs[0].Name)
	require.Equal(t, "test_kept", selection.Pairs[0].From.Name)
	require.Equal(t, "test_kept", selection.Pairs[0].To.Name)
	require.Equal(t, "test_added", selection.Pairs[1].Name)
	require.Nil(t, selection.Pairs[1].From)
	require.Equal(t, "test_added", selection.Pairs[1].To.Name)
}

func TestSelectUpgradeTablesSkipsDependentTables(t *testing.T) {
	parent := testTable("test_parent")
	child := testTable("test_parent_child")
	child.Parent = parent
	parent.Relations = schema.Tables{child}
	tables := schema.Tables{parent}.FlattenTables()

	selection, err := selectUpgradeTables(specs.Source{
		Metadata:            specs.Metadata{Name: "test", Version: "v1.0.0"},
		Tables:              []string{"test_parent"},
		SkipDependentTables: lo.ToPtr(true),
	}, tables, tables)
	require.NoError(t, err)

	require.Empty(t, selection.RemovedTables)
	require.Len(t, selection.Pairs, 1)
	require.Equal(t, "test_parent", selection.Pairs[0].Name)
}

func TestTransformUpgradeTables(t *testing.T) {
	fromTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		schema.CqIDColumn,
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	toTable := &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
		schema.CqIDColumn,
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	pairs := []upgradeTablePair{
		{Name: "okta_policy_rules", From: fromTable, To: toTable},
		{Name: "okta_added", To: testTable("okta_added")},
	}

	cases := []struct {
		name            string
		destinationSpec specs.Destination
		transformers    []pluginPb.PluginClient
		wantTableName   string
		wantFromPKs     []string
		wantToPKs       []string
	}{
		{
			name:            "overwrite-delete-stale keeps source primary keys",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeOverwriteDeleteStale, PKMode: specs.PKModeDefaultKeys},
			wantTableName:   "okta_policy_rules",
			wantFromPKs:     []string{"id"},
			wantToPKs:       []string{"id", "policy_id"},
		},
		{
			name:            "append removes primary keys",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeAppend, PKMode: specs.PKModeDefaultKeys},
			wantTableName:   "okta_policy_rules",
		},
		{
			name:            "transformer plugins are applied",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeOverwriteDeleteStale, PKMode: specs.PKModeDefaultKeys},
			transformers:    []pluginPb.PluginClient{renamingTransformerClient{}},
			wantTableName:   "renamed_okta_policy_rules",
			wantFromPKs:     []string{"id"},
			wantToPKs:       []string{"id", "policy_id"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.destinationSpec.Name = "postgresql"
			sourceSpec := specs.Source{Metadata: specs.Metadata{Name: "okta", Version: "v6.8.2"}}

			got, err := transformUpgradeTables(context.Background(), sourceSpec, tc.destinationSpec, tc.transformers, false, pairs)
			require.NoError(t, err)
			require.Len(t, got, 2)
			require.Equal(t, "okta_policy_rules", got[0].Name)
			require.Equal(t, "okta_added", got[1].Name)
			require.Nil(t, got[1].From)
			require.NotNil(t, got[1].To)

			from := tableFromSchemaBytes(t, got[0].From)
			to := tableFromSchemaBytes(t, got[0].To)
			require.Equal(t, tc.wantTableName, from.Name)
			require.Equal(t, tc.wantTableName, to.Name)
			require.Equal(t, tc.wantFromPKs, emptyToNil(from.PrimaryKeys()))
			require.Equal(t, tc.wantToPKs, emptyToNil(to.PrimaryKeys()))
			require.NotNil(t, to.Columns.Get(schema.CqSourceNameColumn.Name))
			require.NotNil(t, to.Columns.Get(schema.CqSyncTimeColumn.Name))
		})
	}
}

type fakeDestinationClient struct {
	pluginPb.PluginClient
	assessErr error

	initRequests   []*pluginPb.Init_Request
	assessRequests []*pluginPb.AssessTables_Request
	writeCalled    bool
}

func (f *fakeDestinationClient) Init(_ context.Context, req *pluginPb.Init_Request, _ ...grpc.CallOption) (*pluginPb.Init_Response, error) {
	f.initRequests = append(f.initRequests, req)
	return &pluginPb.Init_Response{}, nil
}

func (f *fakeDestinationClient) AssessTables(_ context.Context, req *pluginPb.AssessTables_Request, _ ...grpc.CallOption) (*pluginPb.AssessTables_Response, error) {
	f.assessRequests = append(f.assessRequests, req)
	if f.assessErr != nil {
		return nil, f.assessErr
	}
	return &pluginPb.AssessTables_Response{Tables: []*pluginPb.AssessTables_TableFinding{
		{TableName: "test_kept", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE},
	}}, nil
}

func (f *fakeDestinationClient) Write(context.Context, ...grpc.CallOption) (grpc.ClientStreamingClient[pluginPb.Write_Request, pluginPb.Write_Response], error) {
	f.writeCalled = true
	return nil, errors.New("write must not be called")
}

func TestAssessTables(t *testing.T) {
	tables := []upgradeTableSchemas{
		{Name: "test_kept", From: []byte("old"), To: []byte("new")},
		{Name: "test_added", To: []byte("added")},
	}
	cases := []struct {
		name             string
		destination      *fakeDestinationClient
		migrateMode      specs.MigrateMode
		wantMigrateForce bool
		wantFindings     []*pluginPb.AssessTables_TableFinding
		wantErr          string
	}{
		{
			name:         "returns destination findings",
			destination:  &fakeDestinationClient{},
			wantFindings: []*pluginPb.AssessTables_TableFinding{{TableName: "test_kept", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE}},
		},
		{
			name:             "passes forced migrate mode",
			destination:      &fakeDestinationClient{},
			migrateMode:      specs.MigrateModeForced,
			wantMigrateForce: true,
			wantFindings:     []*pluginPb.AssessTables_TableFinding{{TableName: "test_kept", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE}},
		},
		{
			name:        "destination without the assessment RPC is unknown",
			destination: &fakeDestinationClient{assessErr: status.Error(codes.Unimplemented, "method AssessTables not implemented")},
			wantFindings: []*pluginPb.AssessTables_TableFinding{
				{TableName: "test_kept", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, IncompleteCoverageReason: destinationNoAssessmentReason},
				{TableName: "test_added", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, IncompleteCoverageReason: destinationNoAssessmentReason},
			},
		},
		{
			name:        "other assessment errors fail the check",
			destination: &fakeDestinationClient{assessErr: status.Error(codes.Internal, "boom")},
			wantErr:     "failed to assess tables for destination postgresql",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			destinationSpec := specs.Destination{Metadata: specs.Metadata{Name: "postgresql"}, MigrateMode: tc.migrateMode, Spec: map[string]any{"connection_string": "postgres://unused"}}

			findings, err := assessTables(t.Context(), tc.destination, destinationSpec, tables)
			if tc.wantErr != "" {
				require.ErrorContains(t, err, tc.wantErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantFindings, findings)
			}

			require.Len(t, tc.destination.initRequests, 1)
			require.True(t, tc.destination.initRequests[0].NoConnection)
			require.Len(t, tc.destination.assessRequests, 1)
			require.Equal(t, tc.wantMigrateForce, tc.destination.assessRequests[0].MigrateForce)
			require.Equal(t, []*pluginPb.AssessTables_TablePair{
				{OldTable: []byte("old"), NewTable: []byte("new")},
				{NewTable: []byte("added")},
			}, tc.destination.assessRequests[0].Tables)
			require.False(t, tc.destination.writeCalled)
		})
	}
}

func TestUpgradeCheck(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testConfig := path.Join(path.Dir(filename), "testdata", "transformation.yml")
	cmd := NewCmdRoot()
	var out, errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs(append([]string{"upgrade", "check", testConfig, "--source", "test", "--to", "v4.7.0"}, testCommandArgs(t)...))

	var exitCodeErr *ExitCodeError
	require.ErrorAs(t, cmd.Execute(), &exitCodeErr)
	require.Equal(t, upgradeExitUnknown, exitCodeErr.Code)
	require.Empty(t, errOut.String())

	report := out.String()
	require.Contains(t, report, "test v4.5.1 → v4.7.0 | test (cloudquery/test@v2.5.1)\nwrite_mode: overwrite-delete-stale | pk_mode: default\nUNKNOWN — ")
	require.Contains(t, report, "test_some_table: "+destinationNoAssessmentReason)
}

var updateUpgradeCheckGolden = flag.Bool("update-upgrade-check-golden", false, "rewrite the expected upgrade check reports")

func TestUpgradeCheckRFCCases(t *testing.T) {
	setColorOutput(t, false)
	cqDir := t.TempDir()

	cases := []struct {
		name         string
		toVersion    string
		wantVerdict  string
		wantExitCode int
	}{
		{name: "datadog", toVersion: "v6.0.0", wantVerdict: "REVIEW REQUIRED — 6 tables need a manual migration", wantExitCode: 3},
		{name: "okta-overwrite-delete-stale", toVersion: "v7.0.0", wantVerdict: "REVIEW REQUIRED — 1 table needs a manual migration, 1 new table", wantExitCode: 3},
		{name: "okta-append", toVersion: "v7.0.0", wantVerdict: "AUTOMATICALLY MIGRATABLE — 1 changed table, 1 new table", wantExitCode: 0},
		{name: "gcp-selected-tables", toVersion: "v23.0.0", wantVerdict: "SELECTED TABLES REMOVED — 2 removed tables", wantExitCode: 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			configPath := filepath.Join("testdata", "upgrade-check", tc.name+".yml")
			specReader, err := specs.NewSpecReader([]string{configPath})
			require.NoError(t, err)
			runCheck := func(output string) (string, error) {
				defer CloseLogFile()
				cmd := NewCmdRoot()
				var out bytes.Buffer
				cmd.SetOut(&out)
				cmd.SetArgs([]string{
					"upgrade", "check", configPath, "--source", specReader.Sources[0].Name, "--to", tc.toVersion, "--output", output,
					"--cq-dir", cqDir, "--log-file-name", filepath.Join(t.TempDir(), "cloudquery.log"),
				})
				err := cmd.Execute()
				return strings.ReplaceAll(out.String(), specReader.Destinations[0].Version, "<version>"), err
			}

			text, err := runCheck(upgradeOutputText)
			require.Contains(t, text, "\n"+tc.wantVerdict+"\n")
			requireUpgradeCheckGolden(t, tc.name+".txt", text)
			requireUpgradeCheckExitCode(t, err, tc.wantExitCode)

			jsonReport, err := runCheck(upgradeOutputJSON)
			requireUpgradeCheckGolden(t, tc.name+".json", jsonReport)
			requireUpgradeCheckExitCode(t, err, tc.wantExitCode)
		})
	}
}

func requireUpgradeCheckExitCode(t *testing.T, err error, want int) {
	t.Helper()
	if want == 0 {
		require.NoError(t, err)
		return
	}
	var exitErr *ExitCodeError
	require.ErrorAs(t, err, &exitErr)
	require.Equal(t, want, exitErr.Code)
}

func requireUpgradeCheckGolden(t *testing.T, name, got string) {
	t.Helper()
	goldenPath := filepath.Join("testdata", "upgrade-check", name)
	if *updateUpgradeCheckGolden {
		require.NoError(t, os.WriteFile(goldenPath, []byte(got), 0o644))
	}
	want, err := os.ReadFile(goldenPath)
	require.NoError(t, err)
	require.Equal(t, string(want), got)
}

func TestLoadUpgradeCheckSpecsInjectsPlatformDestination(t *testing.T) {
	tenant := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/external-syncs/supported-source-versions":
			_ = json.NewEncoder(w).Encode(map[string]string{"cloudquery/aws": "v1.0.0"})
		case "/api/external-syncs/whoami":
			_ = json.NewEncoder(w).Encode(map[string]any{"tenant_id": "11111111-1111-1111-1111-111111111111", "plugin_version": "v1.0.1"})
		default:
			http.Error(w, "unexpected "+r.URL.Path, http.StatusNotFound)
		}
	}))
	t.Cleanup(tenant.Close)
	payload, err := json.Marshal(map[string]any{"u": tenant.URL, "tm": "team-x"})
	require.NoError(t, err)
	t.Setenv("CLOUDQUERY_API_KEY", "cqpd_"+base64.RawURLEncoding.EncodeToString(payload)+".sig")

	_, filename, _, _ := runtime.Caller(0)
	testConfig := path.Join(path.Dir(filename), "testdata", "validate-config-platform-source-only.yml")

	loaded, err := loadUpgradeCheckSpecs(t.Context(), []string{testConfig}, "aws")
	require.NoError(t, err)
	require.Equal(t, "aws", loaded.source.Name)
	require.Len(t, loaded.destinations, 1)
	require.Equal(t, "platform", loaded.destinations[0].Name)
}

func TestUnknownTableFindingsUseTransformedTableNames(t *testing.T) {
	pairs := []upgradeTablePair{
		{Name: "test_kept", From: testTable("test_kept"), To: testTable("test_kept")},
		{Name: "test_added", To: testTable("test_added")},
	}
	tables, err := transformUpgradeTables(t.Context(), specs.Source{Metadata: specs.Metadata{Name: "test"}}, specs.Destination{Metadata: specs.Metadata{Name: "postgresql"}}, []pluginPb.PluginClient{renamingTransformerClient{}}, false, pairs)
	require.NoError(t, err)

	findings := unknownTableFindings(tables, destinationNoAssessmentReason)

	require.Equal(t, []string{"renamed_test_kept", "renamed_test_added"}, []string{findings[0].TableName, findings[1].TableName})
}
