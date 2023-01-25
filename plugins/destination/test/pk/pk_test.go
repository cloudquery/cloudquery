package pk_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/test/pk"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	type testCase struct {
		table    *schema.Table
		expected string
	}

	for _, tc := range []testCase{
		{
			table: &schema.Table{
				Name: "int",
				Columns: schema.ColumnList{
					{
						Name:            "col1",
						Type:            schema.TypeInt,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				},
			},
			expected: "col1",
		},
		{
			table: &schema.Table{
				Name: "int_str",
				Columns: schema.ColumnList{
					{
						Name:            "col1",
						Type:            schema.TypeInt,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:            "col2",
						Type:            schema.TypeString,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				},
			},
			expected: "col1,col2",
		},
	} {
		t.Run(tc.table.Name, func(t *testing.T) {
			require.Equal(t, tc.expected, pk.Columns(tc.table))
		})
	}
}

func TestColumns(t *testing.T) {
	type testCase struct {
		table    *schema.Table
		resource []any
		expected string
	}

	for _, tc := range []testCase{
		{
			table: &schema.Table{
				Name: "int",
				Columns: schema.ColumnList{
					{
						Name:            "col1",
						Type:            schema.TypeInt,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				},
			},
			resource: []any{2},
			expected: "2",
		},
		{
			table: &schema.Table{
				Name: "int_str",
				Columns: schema.ColumnList{
					{
						Name:            "col1",
						Type:            schema.TypeInt,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:            "col2",
						Type:            schema.TypeString,
						CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				},
			},
			resource: []any{2, "some"},
			expected: "2,some",
		},
	} {
		t.Run(tc.table.Name, func(t *testing.T) {
			require.Equal(t, tc.expected, pk.Convert(tc.table, tc.resource))
		})
	}
}
