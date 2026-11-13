package queries

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestCreateTableAndInsertEscapesTrickyIdentifiers(t *testing.T) {
	table := &schema.Table{
		Name: "db`name",
		Columns: schema.ColumnList{
			schema.Column{
				Name:    `col\slash`,
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{
				Name:    "col`tick",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
		},
	}

	createSQL, err := CreateTable(table, "", spec.DefaultEngine(), nil, nil, nil)
	require.NoError(t, err)
	require.Contains(t, createSQL, "CREATE TABLE IF NOT EXISTS `db``name`")
	require.Contains(t, createSQL, "`col\\\\slash`")
	require.Contains(t, createSQL, "`col``tick`")

	insertSQL := Insert(table)
	require.Equal(t,
		"INSERT INTO `db``name` (`col\\\\slash`, `col``tick`)",
		insertSQL,
	)
}
