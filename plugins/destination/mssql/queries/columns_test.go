package queries

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestAddColumn(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `ALTER TABLE [cq].[table_name] ADD [my_col] bigint NOT NULL;`
	)

	query := AddColumn(schemaName, &schema.Table{Name: "table_name"}, &schema.Column{
		Name:    "my_col",
		Type:    arrow.PrimitiveTypes.Int64,
		NotNull: true,
	})

	require.Equal(t, expected, query)
}
