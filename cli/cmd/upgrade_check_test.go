package cmd

import (
	"context"
	"errors"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
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

	from, to, err := loadUpgradeSourceTables(ctx, sourceSpec, "v4.7.0",
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
}

func TestLoadUpgradeSourceTablesRejectsFixedPathRegistries(t *testing.T) {
	for _, registry := range []specs.Registry{specs.RegistryLocal, specs.RegistryGRPC, specs.RegistryDocker} {
		t.Run(registry.String(), func(t *testing.T) {
			sourceSpec := specs.Source{Metadata: specs.Metadata{Name: "test", Path: "/plugins/test", Version: "v1.0.0", Registry: registry}}

			_, _, err := loadUpgradeSourceTables(context.Background(), sourceSpec, "v2.0.0")

			require.ErrorContains(t, err, "upgrade check supports only sources from the cloudquery or github registry")
		})
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
