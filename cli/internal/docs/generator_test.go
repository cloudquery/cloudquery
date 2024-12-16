//go:build !windows

package docs

// We skip this test on Windows because it fails on newline CR and LF differences,
// not considered worth the effort to fix it right now.

import (
	"os"
	"path"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
)

var testTables = []*schema.Table{
	{
		Name:        "test_table",
		Title:       "Test Table",
		Description: "Description for test table",
		Columns: []schema.Column{
			{
				Name: "int_col",
				Type: arrow.PrimitiveTypes.Int64,
			},
			{
				Name:       "id_col",
				Type:       arrow.PrimitiveTypes.Int64,
				PrimaryKey: true,
			},
			{
				Name:       "id_col2",
				Type:       arrow.PrimitiveTypes.Int64,
				PrimaryKey: true,
			},
			{
				Name: "json_col",
				Type: types.ExtensionTypes.JSON,
			},
			{
				Name: "list_col",
				Type: arrow.ListOf(arrow.PrimitiveTypes.Int64),
			},
			{
				Name: "map_col",
				Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.PrimitiveTypes.Int64),
			},
			{
				Name: "struct_col",
				Type: arrow.StructOf(arrow.Field{Name: "string_field", Type: arrow.BinaryTypes.String}, arrow.Field{Name: "int_field", Type: arrow.PrimitiveTypes.Int64}),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "relation_table",
				Title:       "Relation Table",
				Description: "Description for relational table",
				Columns: []schema.Column{
					{
						Name: "string_col",
						Type: arrow.BinaryTypes.String,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "relation_relation_table_b",
						Title:       "Relation Relation Table B",
						Description: "Description for relational table's relation",
						Columns: []schema.Column{
							{
								Name: "string_col",
								Type: arrow.BinaryTypes.String,
							},
						},
					},
					{
						Name:        "relation_relation_table_a",
						Title:       "Relation Relation Table A",
						Description: "Description for relational table's relation",
						Columns: []schema.Column{
							{
								Name: "string_col",
								Type: arrow.BinaryTypes.String,
							},
						},
					},
				},
			},
			{
				Name:        "relation_table2",
				Title:       "Relation Table2",
				Description: "Description for second relational table",
				Columns: []schema.Column{
					{
						Name: "string_col",
						Type: arrow.BinaryTypes.String,
					},
				},
			},
		},
	},
	{
		Name:          "incremental_table",
		Title:         "Incremental Table",
		Description:   "Description for incremental table",
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name: "int_col",
				Type: arrow.PrimitiveTypes.Int64,
			},
			{
				Name:           "id_col",
				Type:           arrow.PrimitiveTypes.Int64,
				PrimaryKey:     true,
				IncrementalKey: true,
			},
			{
				Name:           "id_col2",
				Type:           arrow.PrimitiveTypes.Int64,
				IncrementalKey: true,
			},
		},
	},
	{
		Name:        "test_table_with_primary_key_component",
		Title:       "Test Table With Primary Key Component",
		Description: "table using a primary key component",
		Columns: []schema.Column{
			{
				Name:       schema.CqIDColumn.Name,
				Type:       schema.CqIDColumn.Type,
				PrimaryKey: true,
			},
			{
				Name: "int_col",
				Type: arrow.PrimitiveTypes.Int64,
			},
			{
				Name:                "id_col",
				Type:                arrow.PrimitiveTypes.Int64,
				PrimaryKeyComponent: true,
			},
			{
				Name:                "id_col2",
				Type:                arrow.PrimitiveTypes.Int64,
				PrimaryKeyComponent: true,
			},
		},
	},
	{
		Name:        "paid_table",
		Title:       "Paid Table",
		Description: "Description for paid table",
		IsPaid:      true,
		Columns:     []schema.Column{},
	},
}

func TestGeneratePluginDocs(t *testing.T) {
	g := NewGenerator("test", testTables)
	cup := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))

	t.Run("Markdown", func(t *testing.T) {
		tmpdir := t.TempDir()

		err := g.Generate(tmpdir, FormatMarkdown)
		if err != nil {
			t.Fatalf("unexpected error calling GeneratePluginDocs: %v", err)
		}

		expectFiles := []string{"test_table.md", "relation_table.md", "relation_relation_table_a.md", "relation_relation_table_b.md", "incremental_table.md", "paid_table.md", "README.md", "test_table_with_primary_key_component.md"}
		for _, exp := range expectFiles {
			t.Run(exp, func(t *testing.T) {
				output := path.Join(tmpdir, exp)
				got, err := os.ReadFile(output)
				require.NoError(t, err)
				cup.SnapshotT(t, got)
			})
		}
	})

	t.Run("JSON", func(t *testing.T) {
		tmpdir := t.TempDir()

		err := g.Generate(tmpdir, FormatJSON)
		if err != nil {
			t.Fatalf("unexpected error calling GeneratePluginDocs: %v", err)
		}

		expectFiles := []string{"__tables.json"}
		for _, exp := range expectFiles {
			t.Run(exp, func(t *testing.T) {
				output := path.Join(tmpdir, exp)
				got, err := os.ReadFile(output)
				require.NoError(t, err)
				cup.SnapshotT(t, got)
			})
		}
	})
}
