package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/stretchr/testify/require"
)

func TestCreateTable(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `CREATE TABLE [cq].[table_name] (
  [_cq_id] uniqueidentifier UNIQUE NOT NULL,
  [_cq_parent_id] uniqueidentifier,
  [_cq_source_name] nvarchar(4000),
  [_cq_sync_time] datetime2,
  [extra_col] float NOT NULL,
  CONSTRAINT [table_name_cqpk] PRIMARY KEY (
  [extra_col]
  )
);`
	)

	query := CreateTable(schemaName,
		schema.CQSchemaToArrow(&schema.Table{
			Name: "table_name",
			Columns: schema.ColumnList{
				schema.CqIDColumn,
				schema.CqParentIDColumn,
				schema.CqSourceNameColumn,
				schema.CqSyncTimeColumn,
				schema.Column{
					Name:            "extra_col",
					Type:            schema.TypeFloat,
					CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		}),
		true,
	)

	require.Equal(t, expected, query)
}

func TestCreateTableNoPK(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `CREATE TABLE [cq].[table_name] (
  [_cq_id] uniqueidentifier UNIQUE NOT NULL,
  [_cq_parent_id] uniqueidentifier,
  [_cq_source_name] nvarchar(4000),
  [_cq_sync_time] datetime2,
  [extra_col] float NOT NULL
);`
	)

	query := CreateTable(schemaName,
		schema.CQSchemaToArrow(&schema.Table{
			Name: "table_name",
			Columns: schema.ColumnList{
				schema.CqIDColumn,
				schema.CqParentIDColumn,
				schema.CqSourceNameColumn,
				schema.CqSyncTimeColumn,
				schema.Column{
					Name:            "extra_col",
					Type:            schema.TypeFloat,
					CreationOptions: schema.ColumnCreationOptions{NotNull: true},
				},
			},
		}),
		true,
	)

	require.Equal(t, expected, query)
}

func TestCreateTableCustomPKConstraintName(t *testing.T) {
	t.Skip("https://github.com/cloudquery/plugin-sdk/pull/833")
	const (
		schemaName = "cq"
		expected   = `CREATE TABLE [cq].[table_name] (
  [_cq_id] uniqueidentifier UNIQUE NOT NULL,
  [_cq_parent_id] uniqueidentifier,
  [_cq_source_name] nvarchar(4000),
  [_cq_sync_time] datetime2,
  [extra_col] float NOT NULL,
  CONSTRAINT [custom_pk_constraint_name] PRIMARY KEY (
  [extra_col]
  )
);`
	)

	query := CreateTable(schemaName,
		schema.CQSchemaToArrow(&schema.Table{
			Name:             "table_name",
			PkConstraintName: "custom_pk_constraint_name",
			Columns: schema.ColumnList{
				schema.CqIDColumn,
				schema.CqParentIDColumn,
				schema.CqSourceNameColumn,
				schema.CqSyncTimeColumn,
				schema.Column{
					Name:            "extra_col",
					Type:            schema.TypeFloat,
					CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true, NotNull: true},
				},
			},
		}),
		true,
	)

	require.Equal(t, expected, query)
}
