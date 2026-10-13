package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type renamingTransformerClient struct {
	pluginPb.PluginClient
}

func (renamingTransformerClient) TransformSchema(_ context.Context, req *pluginPb.TransformSchema_Request, _ ...grpc.CallOption) (*pluginPb.TransformSchema_Response, error) {
	sc, err := pluginPb.NewSchemaFromBytes(req.Schema)
	if err != nil {
		return nil, err
	}
	table, err := schema.NewTableFromArrowSchema(sc)
	if err != nil {
		return nil, err
	}
	table.Name = "renamed_" + table.Name
	b, err := pluginPb.SchemaToBytes(table.ToArrowSchema())
	if err != nil {
		return nil, err
	}
	return &pluginPb.TransformSchema_Response{Schema: b}, nil
}

func TestTransformSchemaForDestination(t *testing.T) {
	sourceTable := &schema.Table{Name: "test_table", Columns: schema.ColumnList{
		schema.CqIDColumn,
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true, Unique: true},
	}}

	cases := []struct {
		name            string
		destinationSpec specs.Destination
		syncGroupId     string
		transformers    []pluginPb.PluginClient
		wantTableName   string
		wantPKs         []string
		wantUnique      bool
		wantSyncGroupId bool
	}{
		{
			name:            "default keys keep source primary keys",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeOverwriteDeleteStale, PKMode: specs.PKModeDefaultKeys},
			wantTableName:   "test_table",
			wantPKs:         []string{"id"},
			wantUnique:      true,
		},
		{
			name:            "append removes primary keys and unique constraints",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeAppend, PKMode: specs.PKModeDefaultKeys},
			wantTableName:   "test_table",
		},
		{
			name:            "cq-id pk mode uses _cq_id as the primary key",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeOverwrite, PKMode: specs.PKModeCQID},
			wantTableName:   "test_table",
			wantPKs:         []string{schema.CqIDColumn.Name},
			wantUnique:      true,
		},
		{
			name:            "sync group id adds the sync group id column",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeAppend, PKMode: specs.PKModeDefaultKeys},
			syncGroupId:     "group",
			wantTableName:   "test_table",
			wantSyncGroupId: true,
		},
		{
			name:            "transformer plugins run after the record transformer",
			destinationSpec: specs.Destination{WriteMode: specs.WriteModeAppend, PKMode: specs.PKModeDefaultKeys},
			transformers:    []pluginPb.PluginClient{renamingTransformerClient{}, renamingTransformerClient{}},
			wantTableName:   "renamed_renamed_test_table",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			recordTransformer := newDestinationRecordTransformer(tc.destinationSpec, "test", time.Now(), tc.syncGroupId, false)

			b, err := transformSchemaForDestination(context.Background(), recordTransformer, tc.transformers, sourceTable.ToArrowSchema())
			require.NoError(t, err)

			got := tableFromSchemaBytes(t, b)
			require.Equal(t, tc.wantTableName, got.Name)
			require.Equal(t, tc.wantPKs, emptyToNil(got.PrimaryKeys()))
			require.Equal(t, tc.wantUnique, got.Columns.Get("id").Unique)
			require.NotNil(t, got.Columns.Get(schema.CqSourceNameColumn.Name))
			require.NotNil(t, got.Columns.Get(schema.CqSyncTimeColumn.Name))
			require.Equal(t, tc.wantSyncGroupId, got.Columns.Get("_cq_sync_group_id") != nil)
		})
	}
}

func tableFromSchemaBytes(t *testing.T, b []byte) *schema.Table {
	t.Helper()
	sc, err := pluginPb.NewSchemaFromBytes(b)
	require.NoError(t, err)
	table, err := schema.NewTableFromArrowSchema(sc)
	require.NoError(t, err)
	return table
}

func emptyToNil(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	return s
}
