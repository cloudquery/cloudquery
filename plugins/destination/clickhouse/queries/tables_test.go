package queries

import (
	"encoding/json"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
)

func TestCreateTable(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableWithPartitionBy(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), []spec.PartitionStrategy{{Tables: []string{"*"}, PartitionBy: "toYYYYMM(`_cq_sync_time`)"}}, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_partition_by.sql")
}

func TestCreateTableWithPartitionByErrorsIfTwoMatching(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), []spec.PartitionStrategy{
		{Tables: []string{"*"}, PartitionBy: "toYYYYMM(`_cq_sync_time`)"},
		{Tables: []string{"*"}, PartitionBy: "toYYYYMM(`_cq_sync_time`)"},
	}, nil, nil)
	require.Error(t, err)
	require.Empty(t, query)
}

func TestCreateTableWithPartitionBySkipsIfNoMatch(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), []spec.PartitionStrategy{{Tables: []string{"*"}, SkipTables: []string{"table_name"}, PartitionBy: "toYYYYMM(`_cq_sync_time`)"}}, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableWithPartitionBySkipsIfIncremental(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
		IsIncremental: true,
	}, "", spec.DefaultEngine(), []spec.PartitionStrategy{{Tables: []string{"*"}, SkipIncrementalTables: true, PartitionBy: "toYYYYMM(`_cq_sync_time`)"}}, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableWithOrderBy(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, []spec.OrderByStrategy{{Tables: []string{"table_name"}, OrderBy: []string{"`_cq_sync_time`", "`_cq_id`"}}}, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_order_by.sql")
}

func TestCreateTableWithOrderByErrorsIfTwoMatching(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, []spec.OrderByStrategy{
		{Tables: []string{"table_name"}, OrderBy: []string{"`_cq_sync_time`", "`_cq_id`"}},
		{Tables: []string{"table_name"}, OrderBy: []string{"`_cq_sync_time`", "`_cq_id`"}},
	}, nil)
	require.Error(t, err)
	require.Empty(t, query)
}

func TestCreateTableWithOrderBySkipsIfNoMatch(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, []spec.OrderByStrategy{{Tables: []string{"table_name"}, SkipTables: []string{"table_name"}, OrderBy: []string{"`_cq_sync_time`", "`_cq_id`"}}}, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table.sql")
}

func TestCreateTableNoOrderBy(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.Column{Name: "extra_col", Type: arrow.PrimitiveTypes.Float64},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_empty_order_by.sql")
}

func TestCreateTableOnCluster(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "my_cluster", spec.DefaultEngine(), nil, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_cluster.sql")
}

func TestCreateTableWithEngine(t *testing.T) {
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", &spec.Engine{
		Name:       "ReplicatedMergeTree",
		Parameters: []any{"a", "b", 1, int32(2), int64(3), float32(1.2), float64(3.4), json.Number("327"), false, true},
	}, nil, nil, nil)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_engine.sql")
}

func TestCreateTableWithTTL_NotNullCqSyncTime(t *testing.T) {
	notNullCqSyncTime := schema.CqSyncTimeColumn
	notNullCqSyncTime.NotNull = true
	ttlStrategy := []spec.TTLStrategy{
		{
			Tables:     []string{"table_*"},
			SkipTables: nil,
			TTL:        "INTERVAL 1 DAY + INTERVAL 5415 SECOND",
		},
	}
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			notNullCqSyncTime,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, nil, ttlStrategy)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_ttl_not_null_cq_time.sql")
}

func TestCreateTableWithTTL_NullableCqSyncTime(t *testing.T) {
	ttlStrategy := []spec.TTLStrategy{
		{
			Tables:     []string{"table_*"},
			SkipTables: nil,
			TTL:        "INTERVAL 1 DAY + INTERVAL 5415 SECOND",
		},
	}
	query, err := CreateTable(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:    "extra_col",
				Type:    arrow.PrimitiveTypes.Float64,
				NotNull: true,
			},
			schema.Column{Name: "extra_inet_col", Type: types.NewInetType()},
			schema.Column{Name: "extra_inet_arr_col", Type: arrow.ListOf(types.NewInetType())},
		},
	}, "", spec.DefaultEngine(), nil, nil, ttlStrategy)
	require.NoError(t, err)
	ensureContents(t, query, "create_table_ttl.sql")
}

func TestDropTable(t *testing.T) {
	query := DropTable(&schema.Table{Name: "table_name"}, "")

	ensureContents(t, query, "drop_table.sql")
}

func TestDropTableOnCLuster(t *testing.T) {
	query := DropTable(&schema.Table{Name: "table_name"}, "my_cluster")

	ensureContents(t, query, "drop_table_cluster.sql")
}
