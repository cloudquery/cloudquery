package queries

type colQueryBuilder struct {
	Table  string
	Column *ColumnDefinition
}

func AddColumn(table string, column *ColumnDefinition) string {
	return execTemplate("col_add.sql.tpl", &colQueryBuilder{Table: table, Column: column})
}

func ModifyColumn(table string, column *ColumnDefinition) string {
	return execTemplate("col_mod.sql.tpl", &colQueryBuilder{Table: table, Column: column})
}
