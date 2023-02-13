package client

import (
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

func pgErrToStr(err *pgconn.PgError) string {
	var sb strings.Builder
	sb.WriteString("severity: ")
	sb.WriteString(err.Severity)
	sb.WriteString(", code: ")
	sb.WriteString(err.Code)
	sb.WriteString(", message: ")
	sb.WriteString(err.Message)
	sb.WriteString(", detail :")
	sb.WriteString(err.Detail)
	sb.WriteString(", hint: ")
	sb.WriteString(err.Hint)
	sb.WriteString(", position: ")
	sb.WriteString(strconv.FormatInt(int64(err.Position), 10))
	sb.WriteString(", internal_position: ")
	sb.WriteString(strconv.FormatInt(int64(err.InternalPosition), 10))
	sb.WriteString(", internal_query: ")
	sb.WriteString(err.InternalQuery)
	sb.WriteString(", where: ")
	sb.WriteString(err.Where)
	sb.WriteString(", schema_name: ")
	sb.WriteString(err.SchemaName)
	sb.WriteString(", table_name: ")
	sb.WriteString(err.TableName)
	sb.WriteString(", column_name: ")
	sb.WriteString(err.ColumnName)
	sb.WriteString(", data_type_name: ")
	sb.WriteString(err.DataTypeName)
	sb.WriteString(", constraint_name: ")
	sb.WriteString(err.ConstraintName)
	sb.WriteString(", file: ")
	sb.WriteString(err.File)
	sb.WriteString(", line: ")
	sb.WriteString(strconv.FormatUint(uint64(err.Line), 10))
	sb.WriteString(", routine: ")
	sb.WriteString(err.Routine)
	return sb.String()
}
