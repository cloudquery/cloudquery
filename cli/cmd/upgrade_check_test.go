package cmd

import (
	"bytes"
	"context"
	"errors"
	"flag"
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
	tablesWithoutConnection schema.Tables
	tablesWithConnection    schema.Tables
	errWithConnection       error

	noConnection bool
	initCalls    []bool
	syncCalled   bool
}

func (f *fakeSourceClient) Init(_ context.Context, req *pluginPb.Init_Request, _ ...grpc.CallOption) (*pluginPb.Init_Response, error) {
	f.noConnection = req.NoConnection
	f.initCalls = append(f.initCalls, req.NoConnection)
	return &pluginPb.Init_Response{}, nil
}

func (f *fakeSourceClient) GetTables(_ context.Context, req *pluginPb.GetTables_Request, _ ...grpc.CallOption) (*pluginPb.GetTables_Response, error) {
	if len(req.Tables) != 1 || req.Tables[0] != "*" {
		return nil, errors.New("expected all tables to be requested")
	}
	tables := f.tablesWithoutConnection
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
		Tables:              []string{"test_removed", "test_kept", "test_added", "test_parent"},
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
				{TableName: "test_kept", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true, CoverageIncompleteReason: destinationNoAssessmentReason},
				{TableName: "test_added", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true, CoverageIncompleteReason: destinationNoAssessmentReason},
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
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs(append([]string{"upgrade", "check", testConfig, "--source", "test", "--to", "v4.7.0"}, testCommandArgs(t)...))

	require.NoError(t, cmd.Execute())

	report := out.String()
	require.Contains(t, report, "test v4.5.1 → v4.7.0 | test (cloudquery/test@v2.5.1)\nwrite_mode: overwrite-delete-stale | pk_mode: default\nUNKNOWN — ")
	require.Contains(t, report, "test_some_table: "+destinationNoAssessmentReason)
}

var updateUpgradeCheckGolden = flag.Bool("update-upgrade-check-golden", false, "rewrite the expected upgrade check reports")

func TestUpgradeCheckRFCCases(t *testing.T) {
	setColorOutput(t, false)
	cqDir := t.TempDir()

	cases := []struct {
		name        string
		toVersion   string
		wantVerdict string
	}{
		{name: "datadog", toVersion: "v6.0.0", wantVerdict: "REVIEW REQUIRED — 6 tables need a manual migration"},
		{name: "okta-overwrite-delete-stale", toVersion: "v7.0.0", wantVerdict: "REVIEW REQUIRED — 1 table needs a manual migration, 1 new table"},
		{name: "okta-append", toVersion: "v7.0.0", wantVerdict: "AUTOMATICALLY MIGRATABLE — 1 changed table, 1 new table"},
		{name: "gcp-selected-tables", toVersion: "v23.0.0", wantVerdict: "SELECTED TABLES REMOVED — 2 removed tables"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			configPath := filepath.Join("testdata", "upgrade-check", tc.name+".yml")
			specReader, err := specs.NewSpecReader([]string{configPath})
			require.NoError(t, err)
			defer CloseLogFile()

			cmd := NewCmdRoot()
			var out bytes.Buffer
			cmd.SetOut(&out)
			cmd.SetArgs([]string{
				"upgrade", "check", configPath, "--source", specReader.Sources[0].Name, "--to", tc.toVersion,
				"--cq-dir", cqDir, "--log-file-name", filepath.Join(t.TempDir(), "cloudquery.log"),
			})
			require.NoError(t, cmd.Execute())

			report := strings.ReplaceAll(out.String(), specReader.Destinations[0].Version, "<version>")
			require.Contains(t, report, "\n"+tc.wantVerdict+"\n")
			requireUpgradeCheckGolden(t, tc.name+".txt", report)
		})
	}
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
