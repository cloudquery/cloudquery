package cmd

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestGetSourceV2DestV3DestinationsTransformers(t *testing.T) {
	destinationSpecs := []specs.Destination{
		{WriteMode: specs.WriteModeAppend, PKMode: specs.PKModeDefaultKeys},
		{WriteMode: specs.WriteModeOverwrite, PKMode: specs.PKModeCQID},
		{WriteMode: specs.WriteModeOverwrite, PKMode: specs.PKModeDefaultKeys},
	}
	destinationsVersions := [][]int{
		{3},
		{3},
		{2}, // v2 only destination, should get nil transformer
	}

	transformers := getSourceV2DestV3DestinationsTransformers(destinationSpecs, destinationsVersions)
	require.Len(t, transformers, 3)
	require.NotNil(t, transformers[0])
	require.NotNil(t, transformers[1])
	require.Nil(t, transformers[2])

	sourceTable := &schema.Table{
		Name: "test_table",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true, Unique: true},
			{Name: "value", Type: arrow.BinaryTypes.String},
		},
	}

	// Append mode should remove primary keys and unique constraints
	schemaBytes, err := pluginPb.SchemaToBytes(sourceTable.ToArrowSchema())
	require.NoError(t, err)

	transformedSchema := transformers[0].TransformSchema(sourceTable.ToArrowSchema())
	b, err := pluginPb.SchemaToBytes(transformedSchema)
	require.NoError(t, err)

	table := tableFromSchemaBytes(t, b)
	require.Empty(t, table.PrimaryKeys())
	require.False(t, table.Columns.Get("id").Unique)

	// CQID mode should remove original PKs and set _cq_id as PK
	transformedCQIDSchema := transformers[1].TransformSchema(sourceTable.ToArrowSchema())
	bCQID, err := pluginPb.SchemaToBytes(transformedCQIDSchema)
	require.NoError(t, err)

	tableCQID := tableFromSchemaBytes(t, bCQID)
	require.Equal(t, []string{schema.CqIDColumn.Name}, tableCQID.PrimaryKeys())
	require.True(t, tableCQID.Columns.Get("id").Unique)

	// Test transformSourceV2DestV3Schemas helper
	originalSchemas := [][]byte{schemaBytes}
	transformedSchemas, err := transformSourceV2DestV3Schemas(originalSchemas, transformers[0])
	require.NoError(t, err)
	require.Len(t, transformedSchemas, 1)

	nilSchemas, err := transformSourceV2DestV3Schemas(originalSchemas, nil)
	require.NoError(t, err)
	require.Equal(t, originalSchemas, nilSchemas)
}
