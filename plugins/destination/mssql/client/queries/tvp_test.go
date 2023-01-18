package queries

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/stretchr/testify/require"
)

func TestTVPDrop(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `BEGIN TRY
DROP PROCEDURE [cq].[cq_proc_table_name];
DROP TYPE [cq].[cq_tbl_table_name];
END TRY
BEGIN CATCH
END CATCH;`
	)

	query := TVPDrop(schemaName, &schema.Table{Name: "table_name"})

	require.Equal(t, expected, query)
}

func TestTVPType(t *testing.T) {
	const (
		schemaName = "cq"
		expected   = `CREATE TYPE [cq].[cq_tbl_table_name] AS TABLE (
  [_cq_id] uniqueidentifier UNIQUE NOT NULL,
  [_cq_parent_id] uniqueidentifier,
  [_cq_source_name] varchar(8000),
  [_cq_sync_time] datetimeoffset,
  [extra_col_pk1] float NOT NULL,
  [extra_col_pk2] bit NOT NULL,
  [extra_col_not_pk1] bigint,
  [extra_col_not_pk2] varbinary(max)
);`
	)

	query := TVPType(schemaName, &schema.Table{
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
	})

	require.Equal(t, expected, query)
}

func TestTVPProc(t *testing.T) {
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

INSERT [cq].[table_name]
 SELECT * FROM @TVP AS [src]
 WHERE NOT EXISTS (
  SELECT 1 FROM [cq].[table_name] AS [tgt]
  WHERE (
  [tgt].[extra_col_pk1] = [src].[extra_col_pk1]
  AND
  [tgt].[extra_col_pk2] = [src].[extra_col_pk2]
  )
);
END;`
	)

	query := TVPProc(schemaName, &schema.Table{
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
	})

	require.Equal(t, expected, query)
}
