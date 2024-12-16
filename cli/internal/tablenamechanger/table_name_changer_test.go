package tablenamechanger

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestTableNameChanger(t *testing.T) {
	destinationSpecs := []specs.Destination{
		{Metadata: specs.Metadata{Name: "dest1"}},
		{Metadata: specs.Metadata{Name: "dest2"}},
	}

	t.Run("UpdateTableNames", func(t *testing.T) {
		tests := []struct {
			name              string
			destinationName   string
			tables            map[string]bool
			expectedNewTables map[string]bool
		}{
			{
				name:              "NoChanges",
				destinationName:   "dest1",
				tables:            map[string]bool{"table1": true, "table2": true},
				expectedNewTables: map[string]bool{"table1": true, "table2": true},
			},
			{
				name:              "WithChanges",
				destinationName:   "dest1",
				tables:            map[string]bool{"old_table1": true, "table2": true},
				expectedNewTables: map[string]bool{"new_table1": true, "table2": true},
			},
		}

		tnc := New(destinationSpecs)
		err := tnc.LearnTableNameChange("dest1", "old_table1", createSchemaBytes(t, "new_table1"))
		require.NoError(t, err)

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				newTables := tnc.UpdateTableNames(tt.destinationName, tt.tables)
				require.Equal(t, tt.expectedNewTables, newTables)
			})
		}
	})

	t.Run("UpdateTableName", func(t *testing.T) {
		tests := []struct {
			name            string
			destinationName string
			oldTableName    string
			expectedNewName string
		}{
			{
				name:            "NoChange",
				destinationName: "dest1",
				oldTableName:    "table1",
				expectedNewName: "table1",
			},
			{
				name:            "WithChange",
				destinationName: "dest1",
				oldTableName:    "old_table1",
				expectedNewName: "new_table1",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tnc := New(destinationSpecs)
				err := tnc.LearnTableNameChange("dest1", "old_table1", createSchemaBytes(t, "new_table1"))
				require.NoError(t, err)

				newName := tnc.UpdateTableName(tt.destinationName, tt.oldTableName)
				require.Equal(t, tt.expectedNewName, newName)
			})
		}
	})

	t.Run("LearnTableNameChange", func(t *testing.T) {
		tests := []struct {
			name               string
			learnDestName      string
			learnOldTableName  string
			learnNewTableName  string
			updateDestName     string
			updateOldTableName string
			updateNewTableName string
		}{
			{
				name:               "ValidChange",
				learnDestName:      "dest1",
				learnOldTableName:  "old_table",
				learnNewTableName:  "new_table",
				updateDestName:     "dest1",
				updateOldTableName: "old_table",
				updateNewTableName: "new_table",
			},
			{
				name:               "NoChangeOnSameDestination",
				learnDestName:      "dest1",
				learnOldTableName:  "old_table",
				learnNewTableName:  "old_table",
				updateDestName:     "dest1",
				updateOldTableName: "old_table",
				updateNewTableName: "old_table",
			},
			{
				name:               "ValidChangeButDoesntUpdateOnDifferentDestination",
				learnDestName:      "dest1",
				learnOldTableName:  "old_table",
				learnNewTableName:  "new_table",
				updateDestName:     "dest2",
				updateOldTableName: "old_table",
				updateNewTableName: "old_table",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tnc := New(destinationSpecs)

				err := tnc.LearnTableNameChange(tt.learnDestName, tt.learnOldTableName, createSchemaBytes(t, tt.learnNewTableName))
				require.NoError(t, err)
				newName := tnc.UpdateTableName(tt.updateDestName, tt.updateOldTableName)
				require.Equal(t, tt.updateNewTableName, newName)
			})
		}
	})
}

func createSchemaBytes(t *testing.T, tableName string) []byte {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{tableName})
	sc := arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String, Nullable: true},
			{Name: "col2", Type: arrow.BinaryTypes.String, Nullable: true},
		},
		&md,
	)
	bytes, err := plugin.SchemaToBytes(sc)
	require.NoError(t, err)
	return bytes
}
