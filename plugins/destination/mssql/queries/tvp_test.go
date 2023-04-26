package queries

import (
	"database/sql"
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/stretchr/testify/require"
)

func TestTVPDropProc(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `IF EXISTS (
 SELECT * FROM sys.procedures p
 INNER JOIN sys.schemas s ON p.schema_id = s.schema_id
 WHERE s.[name] = @schemaName AND p.[name] = @procName
)
DROP PROCEDURE [cq].[cq_proc_table_name];`
	)

	query, params := TVPDropProc(schemaName, schema.CQSchemaToArrow(&schema.Table{Name: "table_name"}))

	require.Equal(t, expected, query)
	require.Equal(t, 2, len(params))

	named, ok := params[0].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "schemaName", named.Name)
	require.Equal(t, schemaName, named.Value)

	named, ok = params[1].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "procName", named.Name)
	require.Equal(t, "cq_proc_table_name", named.Value)
}

func TestTVPDropType(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `IF EXISTS (
 SELECT * FROM sys.table_types tt
 INNER JOIN sys.schemas s ON tt.schema_id = s.schema_id
 WHERE s.[name] = @schemaName AND tt.[name] = @typeName
)
DROP TYPE [cq].[cq_tbl_table_name];`
	)

	query, params := TVPDropType(schemaName, schema.CQSchemaToArrow(&schema.Table{Name: "table_name"}))

	require.Equal(t, expected, query)
	require.Equal(t, 2, len(params))

	named, ok := params[0].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "schemaName", named.Name)
	require.Equal(t, schemaName, named.Value)

	named, ok = params[1].(sql.NamedArg)
	require.True(t, ok)
	require.Equal(t, "typeName", named.Name)
	require.Equal(t, "cq_tbl_table_name", named.Value)
}

func TestTVPAddType(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `CREATE TYPE [cq].[cq_tbl_table_name] AS TABLE (
  [_cq_id] uniqueidentifier UNIQUE NOT NULL,
  [_cq_parent_id] uniqueidentifier,
  [_cq_source_name] nvarchar(4000),
  [_cq_sync_time] datetime2,
  [extra_col_pk1] float NOT NULL,
  [extra_col_pk2] bit NOT NULL,
  [extra_col_not_pk1] bigint,
  [extra_col_not_pk2] varbinary(max)
);`
	)

	query := TVPAddType(schemaName, schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col_pk1",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			schema.Column{
				Name:            "extra_col_pk2",
				Type:            schema.TypeBool,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			schema.Column{
				Name: "extra_col_not_pk1",
				Type: schema.TypeInt,
			},
			schema.Column{
				Name: "extra_col_not_pk2",
				Type: schema.TypeByteArray,
			},
		},
	}))

	require.Equal(t, expected, query)
}

func TestTVPAddProc(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `CREATE PROCEDURE [cq].[cq_proc_table_name] @TVP [cq].[cq_tbl_table_name] READONLY
AS
BEGIN
 SET NOCOUNT ON;
 UPDATE [tgt] WITH (UPDLOCK)
 SET
  [tgt].[_cq_id] = [src].[_cq_id],
  [tgt].[_cq_parent_id] = [src].[_cq_parent_id],
  [tgt].[_cq_source_name] = [src].[_cq_source_name],
  [tgt].[_cq_sync_time] = [src].[_cq_sync_time],
  [tgt].[extra_col_not_pk1] = [src].[extra_col_not_pk1],
  [tgt].[extra_col_not_pk2] = [src].[extra_col_not_pk2]
 FROM [cq].[table_name] AS [tgt]
 INNER JOIN @TVP AS [src]
 ON
  [tgt].[extra_col_pk1] = [src].[extra_col_pk1]
  AND
  [tgt].[extra_col_pk2] = [src].[extra_col_pk2]
;

INSERT [cq].[table_name] (
  [_cq_id],
  [_cq_parent_id],
  [_cq_source_name],
  [_cq_sync_time],
  [extra_col_pk1],
  [extra_col_pk2],
  [extra_col_not_pk1],
  [extra_col_not_pk2]
) SELECT
  [src].[_cq_id],
  [src].[_cq_parent_id],
  [src].[_cq_source_name],
  [src].[_cq_sync_time],
  [src].[extra_col_pk1],
  [src].[extra_col_pk2],
  [src].[extra_col_not_pk1],
  [src].[extra_col_not_pk2]
 FROM @TVP AS [src]
 LEFT JOIN [cq].[table_name] AS [tgt] ON (
  [tgt].[extra_col_pk1] = [src].[extra_col_pk1]
  AND
  [tgt].[extra_col_pk2] = [src].[extra_col_pk2]
 ) WHERE (
  [tgt].[extra_col_pk1] IS NULL
  AND
  [tgt].[extra_col_pk2] IS NULL
);
END;`
	)

	query := TVPAddProc(schemaName, schema.CQSchemaToArrow(&schema.Table{
		Name: "table_name",
		Columns: schema.ColumnList{
			schema.CqIDColumn,
			schema.CqParentIDColumn,
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
			schema.Column{
				Name:            "extra_col_pk1",
				Type:            schema.TypeFloat,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			schema.Column{
				Name:            "extra_col_pk2",
				Type:            schema.TypeBool,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			schema.Column{
				Name: "extra_col_not_pk1",
				Type: schema.TypeInt,
			},
			schema.Column{
				Name: "extra_col_not_pk2",
				Type: schema.TypeByteArray,
			},
		},
	}))

	require.Equal(t, expected, query)
}
