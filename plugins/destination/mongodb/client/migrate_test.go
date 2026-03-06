package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetIndexTemplates_NoPrimaryKeys(t *testing.T) {
	table := &schema.Table{
		Name: "test_table",
		Columns: schema.ColumnList{
			{Name: "col1"},
		},
	}
	c := &Client{}
	indexes := c.getIndexTemplates(table)
	require.Empty(t, indexes)
}

func TestGetIndexTemplates_WithPrimaryKeys(t *testing.T) {
	table := &schema.Table{
		Name: "test_table",
		Columns: schema.ColumnList{
			{Name: "id", PrimaryKey: true},
			{Name: "tenant_id", PrimaryKey: true},
			{Name: "data"},
		},
	}

	c := &Client{}
	indexes := c.getIndexTemplates(table)

	require.Len(t, indexes, 1)
	tmpl := indexes[0]
	require.Equal(t, "cq_pk", tmpl.name)

	// Verify the keys are the PK columns in order
	keys := tmpl.model.Keys.(bson.D)
	require.Len(t, keys, 2)
	require.Equal(t, "id", keys[0].Key)
	require.Equal(t, 1, keys[0].Value)
	require.Equal(t, "tenant_id", keys[1].Key)
	require.Equal(t, 1, keys[1].Value)
}

func TestGetIndexTemplates_SinglePrimaryKey(t *testing.T) {
	table := &schema.Table{
		Name: "test_table",
		Columns: schema.ColumnList{
			{Name: "_cq_id", PrimaryKey: true},
			{Name: "name"},
		},
	}

	c := &Client{}
	indexes := c.getIndexTemplates(table)

	require.Len(t, indexes, 1)
	require.Equal(t, "cq_pk", indexes[0].name)

	keys := indexes[0].model.Keys.(bson.D)
	require.Len(t, keys, 1)
	require.Equal(t, "_cq_id", keys[0].Key)
}
