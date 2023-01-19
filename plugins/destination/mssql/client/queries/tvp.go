package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type tvpProcQueryBuilder struct {
	Name    string
	Type    string
	Table   string
	Columns Definitions
	PK      []string
	Values  []string
}

func TVPProcName(schemaName string, tableName string) string {
	const pfx = "cq_proc_"
	return SanitizeID(schemaName, pfx+tableName)
}

func TVPTableType(schemaName string, tableName string) string {
	const pfx = "cq_tbl_"
	return SanitizeID(schemaName, pfx+tableName)
}

func TVPDrop(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Name: TVPProcName(schemaName, table.Name),
		Type: TVPTableType(schemaName, table.Name),
	}

	return execTemplate("tvp_drop.sql.tpl", data)
}

func TVPProc(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Name:   TVPProcName(schemaName, table.Name),
		Type:   TVPTableType(schemaName, table.Name),
		Table:  SanitizeID(schemaName, table.Name),
		PK:     GetPKColumns(table, true),
		Values: GetValueColumns(table.Columns),
	}

	return execTemplate("tvp_proc.sql.tpl", data)
}

func TVPType(schemaName string, table *schema.Table) string {
	data := &tvpProcQueryBuilder{
		Type:    TVPTableType(schemaName, table.Name),
		Columns: GetDefinitions(table.Columns, true),
	}

	return execTemplate("tvp_type.sql.tpl", data)
}
