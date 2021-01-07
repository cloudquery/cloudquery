package sqlserver

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

type Config struct {
	DriverName string
	DSN        string
	Conn       gorm.ConnPool
}

type Dialector struct {
	*Config
}

func (dialector Dialector) Name() string {
	return "sqlserver"
}

func Open(dsn string) gorm.Dialector {
	return &Dialector{Config: &Config{DSN: dsn}}
}

func New(config Config) gorm.Dialector {
	return &Dialector{Config: &config}
}

func (dialector Dialector) Initialize(db *gorm.DB) (err error) {

	// register callbacks
	callbacks.RegisterDefaultCallbacks(db, &callbacks.Config{})
	db.Callback().Create().Replace("gorm:create", Create)

	if dialector.DriverName == "" {
		dialector.DriverName = "sqlserver"
	}

	if dialector.Conn != nil {
		db.ConnPool = dialector.Conn
	} else {
		db.ConnPool, err = sql.Open(dialector.DriverName, dialector.DSN)
		if err != nil {
			return err
		}
	}

	for k, v := range dialector.ClauseBuilders() {
		db.ClauseBuilders[k] = v
	}
	return
}

func (dialector Dialector) ClauseBuilders() map[string]clause.ClauseBuilder {
	return map[string]clause.ClauseBuilder{
		"LIMIT": func(c clause.Clause, builder clause.Builder) {
			if limit, ok := c.Expression.(clause.Limit); ok {
				if stmt, ok := builder.(*gorm.Statement); ok {
					if _, ok := stmt.Clauses["ORDER BY"]; !ok {
						if stmt.Schema != nil && stmt.Schema.PrioritizedPrimaryField != nil {
							builder.WriteString("ORDER BY ")
							builder.WriteQuoted(stmt.Schema.PrioritizedPrimaryField.DBName)
							builder.WriteByte(' ')
						} else {
							builder.WriteString("ORDER BY (SELECT NULL) ")
						}
					}
				}

				if limit.Offset > 0 {
					builder.WriteString("OFFSET ")
					builder.WriteString(strconv.Itoa(limit.Offset))
					builder.WriteString(" ROWS")
				}

				if limit.Limit > 0 {
					if limit.Offset == 0 {
						builder.WriteString("OFFSET 0 ROW")
					}
					builder.WriteString(" FETCH NEXT ")
					builder.WriteString(strconv.Itoa(limit.Limit))
					builder.WriteString(" ROWS ONLY")
				}
			}
		},
	}
}

func (dialector Dialector) DefaultValueOf(field *schema.Field) clause.Expression {
	return clause.Expr{SQL: "NULL"}
}

func (dialector Dialector) Migrator(db *gorm.DB) gorm.Migrator {
	return Migrator{migrator.Migrator{Config: migrator.Config{
		DB:                          db,
		Dialector:                   dialector,
		CreateIndexAfterCreateTable: true,
	}}}
}

func (dialector Dialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	writer.WriteString("@p")
	writer.WriteString(strconv.Itoa(len(stmt.Vars)))
}

func (dialector Dialector) QuoteTo(writer clause.Writer, str string) {
	writer.WriteByte('"')
	if strings.Contains(str, ".") {
		for idx, str := range strings.Split(str, ".") {
			if idx > 0 {
				writer.WriteString(`."`)
			}
			writer.WriteString(str)
			writer.WriteByte('"')
		}
	} else {
		writer.WriteString(str)
		writer.WriteByte('"')
	}
}

var numericPlaceholder = regexp.MustCompile("@p(\\d+)")

func (dialector Dialector) Explain(sql string, vars ...interface{}) string {
	for idx, v := range vars {
		if b, ok := v.(bool); ok {
			if b {
				vars[idx] = 1
			} else {
				vars[idx] = 0
			}
		}
	}

	return logger.ExplainSQL(sql, numericPlaceholder, `'`, vars...)
}

func (dialector Dialector) DataTypeOf(field *schema.Field) string {
	switch field.DataType {
	case schema.Bool:
		return "bit"
	case schema.Int, schema.Uint:
		var sqlType string
		switch {
		case field.Size < 16:
			sqlType = "smallint"
		case field.Size < 31:
			sqlType = "int"
		default:
			sqlType = "bigint"
		}

		if field.AutoIncrement {
			return sqlType + " IDENTITY(1,1)"
		}
		return sqlType
	case schema.Float:
		return "float"
	case schema.String:
		size := field.Size
		hasIndex := field.TagSettings["INDEX"] != "" || field.TagSettings["UNIQUE"] != ""
		if (field.PrimaryKey || hasIndex) && size == 0 {
			size = 256
		}
		if size > 0 && size <= 4000 {
			return fmt.Sprintf("nvarchar(%d)", size)
		}
		return "nvarchar(MAX)"
	case schema.Time:
		return "datetimeoffset"
	case schema.Bytes:
		return "varbinary(MAX)"
	}

	return string(field.DataType)
}

func (dialectopr Dialector) SavePoint(tx *gorm.DB, name string) error {
	tx.Exec("SAVE TRANSACTION " + name)
	return nil
}

func (dialectopr Dialector) RollbackTo(tx *gorm.DB, name string) error {
	tx.Exec("ROLLBACK TRANSACTION " + name)
	return nil
}
