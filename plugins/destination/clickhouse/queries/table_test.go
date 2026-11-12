package queries

import (
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestGetPartitionKeyAndSortingKeyQuery(t *testing.T) {
	names := []string{
		"table_name",
		"table'name",
		`table"name`,
		"table`name",
		`table\name`,
		"db.table",
		"x' OR name = 'y",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			query, params := GetPartitionKeyAndSortingKeyQuery(name, name)
			require.Equal(t, "SELECT partition_key, sorting_key FROM system.tables WHERE database = @databaseName AND name = @tableName", query)
			require.Equal(t, []any{
				driver.NamedValue{Name: "databaseName", Value: name},
				driver.NamedValue{Name: "tableName", Value: name},
			}, params)
		})
	}
}

func TestIdentifiersAreEscaped(t *testing.T) {
	table := &schema.Table{
		Name: "x`.`y",
		Columns: schema.ColumnList{
			schema.Column{Name: "a'b"},
			schema.Column{Name: `c"d`},
			schema.Column{Name: "e`f"},
			schema.Column{Name: `g\`},
			schema.Column{Name: "h.i"},
		},
	}
	const columns = "`a'b`, `c\"d`, `e\\`f`, `g\\\\`, `h.i`"
	require.Equal(t, "INSERT INTO `x\\`.\\`y` ("+columns+")", Insert(table))
	require.Equal(t, "SELECT "+columns+" FROM `x\\`.\\`y`", Read(table))
	require.Equal(t, "DROP TABLE IF EXISTS `x\\`.\\`y`", DropTable(table, ""))
}
