package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v4"
)

// this really cool query is take from https://github.com/go-gorm/postgres/blob/master/migrator.go
// return the following:
//table_name |  index_name   | column_name | non_unique | primary
//------------+---------------+-------------+------------+---------
// ptable2    | ptable2_pkey  | a           | t          | t
// ptable2    | ptable2_pkey  | b           | t          | t

const indexSql = `
select
    t.relname as table_name,
    i.relname as index_name,
    a.attname as column_name,
    ix.indisunique as non_unique,
	ix.indisprimary as primary
from
    pg_class t,
    pg_class i,
    pg_index ix,
    pg_attribute a
where
    t.oid = ix.indrelid
    and i.oid = ix.indexrelid
    and a.attrelid = t.oid
    and a.attnum = ANY(ix.indkey)
    and t.relkind = 'r'
    and t.relname = ?
`

const sqlSelectTableConstraints = `SELECT 
	c.column_name, constraint_name, constraint_type
FROM information_schema.table_constraints tc
	JOIN information_schema.constraint_column_usage AS ccu USING (constraint_schema, constraint_name)
	JOIN information_schema.columns AS c ON c.table_schema = tc.constraint_schema AND tc.table_name = c.table_name AND ccu.column_name = c.column_name
WHERE constraint_type IN ('PRIMARY KEY', 'UNIQUE') AND c.table_catalog = $1 AND c.table_schema = $2 AND c.table_name = $3`

const sqlSelectColumnTypes = `SELECT a.attname as column_name, format_type(a.atttypid, a.atttypmod) AS data_type
FROM pg_attribute a JOIN pg_class b ON a.attrelid = b.relfilenode AND relnamespace = (SELECT oid FROM pg_catalog.pg_namespace WHERE nspname = $1)
WHERE a.attnum > 0 -- hide internal columns
AND NOT a.attisdropped -- hide deleted columns
AND b.relname = $2`

const sqlSelectPrimaryKeys = `select tc.constraint_name, tc.table_name                
from information_schema.table_constraints tc
  join information_schema.key_column_usage kc 
    on kc.table_name = tc.table_name and kc.table_schema = tc.table_schema and kc.constraint_name = tc.constraint_name
where tc.constraint_type = 'PRIMARY KEY'
  and kc.ordinal_position is not null and tc.table_name = 'gcp_compute_addresses'
 group by tc.constraint_name, tc.table_name;`

const sqlAlterTableAddColumn = "alter table "

const sqlAddColumn = "add column $1"

const sqlAlterTableDropColumn = "alter table $1 drop column $2"

const sqlAlterTableAddUniqueConstraint = "alter table $1 add constraint unique ($2)"

const sqlAlterTableDropUniqueConstraint = "alter table $1 drop constraint "

const sqlAlterTableDropCQPrimaryKeyConstraint = "alter table $1 drop constraint if exists cq_pk"

const sqlAlterTableAddCQPrimaryKeyConstraint = "alter table $1 add constraint cq_pk primary key ($2)"

const sqlDropTable = "drop table if exists "

const isTableExistSQL = "select count(*) from information_schema.tables where table_name = $1"

var defaultColumns = []string{"_cq_raw"}

type PostgreSqlSpec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	PgxLogLevel      string `json:"pgx_log_level,omitempty"`
}

type Client struct {
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	spec                specs.Destination
	currentDatabaseName string
	currentSchemaName   string
}

func NewClient(logger zerolog.Logger) *Client {
	return &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
}

func (p *Client) Name() string {
	return "postgresql"
}

func (p *Client) Version() string {
	// change it with builtin-cliversion
	return "v0.0.1"
}

func (p *Client) SetLogger(logger zerolog.Logger) {
	p.logger = logger
}

func (p *Client) Initialize(ctx context.Context, spec specs.Destination) error {
	var specPostgreSql PostgreSqlSpec
	p.spec = spec
	if err := spec.UnmarshalSpec(&specPostgreSql); err != nil {
		return fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}

	logLevel, err := pgx.LogLevelFromString(specPostgreSql.PgxLogLevel)
	if err != nil {
		return fmt.Errorf("failed to parse pgx log level %s: %w", specPostgreSql.PgxLogLevel, err)
	}
	p.logger.Info().Str("pgx_log_level", specPostgreSql.PgxLogLevel).Msg("Initializing postgresql destination")

	pgxConfig, err := pgxpool.ParseConfig(specPostgreSql.ConnectionString)
	if err != nil {
		return errors.Wrap(err, "failed to parse connection string")
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{Value: &pgxUUID.UUID{}, Name: "uuid", OID: pgtype.UUIDOID})
		return nil
	}
	l := zerologadapter.NewLogger(p.logger)
	pgxConfig.ConnConfig.Logger = l
	pgxConfig.ConnConfig.LogLevel = logLevel
	p.conn, err = pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return errors.Wrap(err, "failed to connect to postgresql")
	}

	p.currentDatabaseName, err = p.currentDatabase()
	if err != nil {
		return fmt.Errorf("failed to get current database: %w", err)
	}
	p.currentSchemaName = "public"
	return nil
}

func (p *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)
	primaryKeys := []string{}
	for i, c := range table.Columns {
		pgType, err := SchemaTypeToPg(c.Type)
		if err != nil {
			return errors.Wrap(err, "failed to convert schema type to postgresql type")
		}
		var columnName pgx.Identifier = []string{c.Name}
		fieldDef := columnName.Sanitize() + " " + pgType
		if c.CreationOptions.Unique {
			fieldDef += " UNIQUE"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.CreationOptions.PrimaryKey {
			primaryKeys = append(primaryKeys, c.Name)
		}
		// if write mode is append only we include the internal _cq_fetch_time in
		// the primary key
		// if c.Name == "_cq_fetch_time" && p.spec.WriteMode == specs.WriteModeAppend {
		// 	primaryKeys = append(primaryKeys, c.Name)
		// }

	}
	if len(primaryKeys) > 0 {
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cq_pk PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	} else {
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cq_pk PRIMARY KEY (_cq_id)")
	}
	sb.WriteString(")")
	_, err := p.conn.Exec(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}
	return nil
}

func (p *Client) currentDatabase() (string, error) {
	var db string
	err := p.conn.QueryRow(context.Background(), "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	return db, nil
}

type pgColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func getPgColumnByName(columns []pgColumn, name string) *pgColumn {
	for _, c := range columns {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

func (p *Client) getPgColumns(ctx context.Context, tableName string) ([]pgColumn, error) {
	var columns []pgColumn
	rows, err := p.conn.Query(ctx, sqlSelectColumnTypes, p.currentSchemaName, tableName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var column pgColumn
		if err := rows.Scan(&column.Name, &column.Type); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}

func (p *Client) getTableColumnsConstraints(ctx context.Context, tableName string) ([]pgColumnConstraint, error) {
	var columns []pgColumnConstraint
	rows, err := p.conn.Query(ctx, sqlSelectTableConstraints, p.currentDatabaseName, p.currentSchemaName, tableName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var column pgColumnConstraint
		if err := rows.Scan(&column.Name, &column.ConstraintName, &column.ConstraintType); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}

func (p *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	var err error
	var pgColumns []pgColumn
	var pgColumnsConstraints constraints
	if pgColumns, err = p.getPgColumns(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	if pgColumnsConstraints, err = p.getTableColumnsConstraints(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns constraints: %w", table.Name, err)
	}
	reCreatePrimaryKeys := false

	for _, c := range table.Columns {
		columnType, err := SchemaTypeToPg(c.Type)
		if err != nil {
			return fmt.Errorf("failed to convert schema type %s to postgresql type: %w", c.Type, err)
		}

		pgColumn := getPgColumnByName(pgColumns, c.Name)

		if pgColumn == nil {
			p.logger.Info().Str("table", table.Name).Str("column", c.Name).Msg("Column doesn't exist, creating")
			// create the new column as it doesn't exist
			var tableName pgx.Identifier = []string{table.Name}
			var columnName pgx.Identifier = []string{c.Name}
			sql := "alter table " + tableName.Sanitize() + " add column " + columnName.Sanitize() + " " + columnType
			if c.CreationOptions.Unique {
				sql += " unique"
			}
			if c.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			if _, err := p.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", c.Name, table.Name, err)
			}
		} else if pgColumn.Type != columnType {
			p.logger.Info().Str("table", table.Name).Str("column", c.Name).Str("old_type", pgColumn.Type).Str("new_type", columnType).Msg("Column exist but type is different, re-creating")
			// column exists but type is different

			// if this column contains primary key we will need to recreate the primary key
			if c.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			// right now we will drop the column and re-create. in the future we will have an option to automigrate
			if _, err := p.conn.Exec(ctx, sqlAlterTableDropColumn, table.Name, c.Name); err != nil {
				return fmt.Errorf("failed to drop column %s on table %s: %w", c.Name, table.Name, err)
			}
			sql := sqlAlterTableDropColumn
			if c.CreationOptions.Unique {
				sql += " unique"
			}
			if _, err := p.conn.Exec(ctx, sql, table.Name, c.Name, columnType); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", c.Name, table.Name, err)
			}
		} else {
			p.logger.Info().Str("table", table.Name).Str("column", c.Name).Str("type", c.Type.String()).Msg("Column exists with the same type")
			// if column is the same check if any difference on constraints
			if pgColumnsConstraints.isColumnUnique(c.Name) && !c.CreationOptions.Unique {
				p.logger.Info().Str("table", table.Name).Str("column", c.Name).Msg("Column exist with unique constraint, removing")
				// we are using default pg unique constraint nameing
				constraint_name := fmt.Sprintf("%s_%s_key", table.Name, c.Name)
				if _, err := p.conn.Exec(ctx, sqlAlterTableAddUniqueConstraint+constraint_name, table.Name, c.Name); err != nil {
					return fmt.Errorf("failed to add unique constraint on column %s on table %s: %w", c.Name, table.Name, err)
				}
			}

			if !pgColumnsConstraints.isColumnUnique(c.Name) && c.CreationOptions.Unique {
				p.logger.Info().Str("table", table.Name).Str("column", c.Name).Msg("Column exist without unique constraint, adding")
				constraint_name := fmt.Sprintf("%s_%s_key", table.Name, c.Name)
				if _, err := p.conn.Exec(ctx, sqlAlterTableDropUniqueConstraint+constraint_name, table.Name, c.Name); err != nil {
					return fmt.Errorf("failed to drop unique constraint on column %s on table %s: %w", c.Name, table.Name, err)
				}
			}

			if pgColumnsConstraints.isColumnPrimaryKey(c.Name) != c.CreationOptions.PrimaryKey {
				p.logger.Info().Str("table", table.Name).Str("column", c.Name).Msg("Column exist with different primary keys")
				reCreatePrimaryKeys = true
			}
		}
	}
	if reCreatePrimaryKeys {
		p.logger.Info().Str("table", table.Name).Msg("recreating primary keys")
		var tableName pgx.Identifier = []string{table.Name}
		constraintName := pgx.Identifier([]string{table.Name + "_pkey"}).Sanitize()
		sql := "alter table " + tableName.Sanitize() + " drop constraint if exists " + constraintName
		if _, err := p.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
		}
		sql = "alter table " + tableName.Sanitize() + " add constraint " + constraintName + " primary key (" + strings.Join(table.PrimaryKeys(), ",") + ")"
		if _, err := p.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
		}
	}
	return nil
}

// This is the responsability of the CLI of the client to lock before running migration
func (p *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	p.logger.Info().Strs("tables", tables.TableNames()).Msg("Migrating tables")
	for _, table := range tables {
		if len(table.Columns) == 0 {
			p.logger.Info().Str("table", table.Name).Msg("Table with not columns, skiping")
			continue
		}
		tableExist := 0
		if err := p.conn.QueryRow(ctx, isTableExistSQL, table.Name).Scan(&tableExist); err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExist == 0 {
			p.logger.Info().Str("table", table.Name).Msg("Table exists, creating")
			if err := p.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		} else {
			p.logger.Info().Str("table", table.Name).Msg("Table doesn't exist, auto-migrating")
			if err := p.autoMigrateTable(ctx, table); err != nil {
				return err
			}
		}
		if err := p.Migrate(ctx, table.Relations); err != nil {
			return err
		}
	}
	return nil
}

func upsert(table string, data map[string]interface{}) (string, []interface{}) {
	var sb strings.Builder

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	for c, v := range data {
		columns = append(columns, `"`+c+`"`)
		values = append(values, v)
	}

	sb.WriteString("insert into ")
	sb.WriteString(table)
	sb.WriteString(" (")
	for i, c := range columns {
		sb.WriteString(c)
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i, _ := range values {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < len(values)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	constraintName := fmt.Sprintf("%s_cq_pk", table)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, c := range columns {
		sb.WriteString(c)
		sb.WriteString("=")
		sb.WriteString(table)
		sb.WriteString(".")
		sb.WriteString(c)
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String(), values
}

func (p *Client) Write(ctx context.Context, table string, data map[string]interface{}) error {
	p.logger.Info().Str("table", table).Msg("Writing data")
	// columns := make([]string, 0, len(data))
	// values := make([]interface{}, 0, len(data))
	// for c, v := range data {
	// 	columns = append(columns, `"`+c+`"`)
	// 	values = append(values, v)
	// }
	sql, values := upsert(table, data)
	_, err := p.conn.Exec(ctx, sql, values...)
	if err != nil {
		return fmt.Errorf("failed to insert data with sql '%s': %w", sql, err)
	}
	return nil
}

func (p *Client) ExampleConfig() string {
	return `
connection_string: "postgresql://user:password@localhost:5432/dbname"
`
}

func (p *Client) Drop(ctx context.Context, tables schema.Tables) error {
	p.logger.Info().Strs("tables", tables.TableNames()).Msg("Dropping tables")
	for _, table := range tables {
		var tableName pgx.Identifier = []string{table.Name}
		if _, err := p.conn.Exec(ctx, sqlDropTable+tableName.Sanitize()); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
		}
	}
	return nil
}

func SchemaTypeToPg(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "boolean", nil
	case schema.TypeInt:
		return "bigint", nil
	case schema.TypeFloat:
		return "real", nil
	case schema.TypeUUID:
		return "uuid", nil
	case schema.TypeString:
		return "text", nil
	case schema.TypeStringArray:
		return "text[]", nil
	case schema.TypeTimestamp:
		return "timestamp without time zone", nil
	case schema.TypeJSON:
		return "json", nil
	case schema.TypeUUIDArray:
		return "uuid[]", nil
	case schema.TypeInetArray:
		return "inet[]", nil
	case schema.TypeCIDR:
		return "cidr", nil
	case schema.TypeCIDRArray:
		return "cidr[]", nil
	case schema.TypeMacAddr:
		return "macaddr", nil
	case schema.TypeMacAddrArray:
		return "macaddr[]", nil
	case schema.TypeInet:
		return "inet", nil
	case schema.TypeIntArray:
		return "bigint[]", nil
	default:
		return "", errors.Errorf("unsupported schema type: %s", t)
	}
}
