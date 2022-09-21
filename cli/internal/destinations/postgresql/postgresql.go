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
	"github.com/rs/zerolog"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v4"
)

type PostgreSqlSpec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
}

type Client struct {
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	spec                specs.Destination
	currentDatabaseName string
	currentSchemaName   string
}

type pgTablePrimaryKeys struct {
	name    string
	columns []string
}

type pgColumn struct {
	name string
	typ  string
}

type pgTableColumns struct {
	name    string
	columns []pgColumn
}

// https://wiki.postgresql.org/wiki/Retrieve_primary_key_columns
const sqlSelectPrimaryKeys = `
SELECT a.attname as pkey FROM pg_index i       
JOIN   pg_attribute a ON a.attrelid = i.indrelid
  AND a.attnum = ANY(i.indkey)
WHERE  i.indrelid = $1::regclass
AND    i.indisprimary;
`

const sqlSelectColumnTypes = `SELECT a.attname as column_name, format_type(a.atttypid, a.atttypmod) AS data_type
FROM pg_attribute a JOIN pg_class b ON a.attrelid = b.relfilenode AND relnamespace = (SELECT oid FROM pg_catalog.pg_namespace WHERE nspname = $1)
WHERE a.attnum > 0 -- hide internal columns
AND NOT a.attisdropped -- hide deleted columns
AND b.relname = $2`

const sqlDropTable = "drop table if exists "

const isTableExistSQL = "select count(*) from information_schema.tables where table_name = $1"

func NewClient(logger zerolog.Logger) *Client {
	return &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
}

func (*Client) Name() string {
	return "postgresql"
}

func (*Client) Version() string {
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

	logLevel, err := pgx.LogLevelFromString(specPostgreSql.PgxLogLevel.String())
	if err != nil {
		return fmt.Errorf("failed to parse pgx log level %s: %w", specPostgreSql.PgxLogLevel, err)
	}
	p.logger.Info().Str("pgx_log_level", specPostgreSql.PgxLogLevel.String()).Msg("Initializing postgresql destination")

	pgxConfig, err := pgxpool.ParseConfig(specPostgreSql.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to parse connection string %w", err)
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
		return fmt.Errorf("failed to connect to postgresql: %w", err)
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
	tableName := pgx.Identifier{table.Name}.Sanitize()
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(tableName)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, c := range table.Columns {
		pgType, err := SchemaTypeToPg(c.Type)
		if err != nil {
			return fmt.Errorf("failed to convert schema type %s to pg type: %w", c.Type, err)
		}
		columnName := pgx.Identifier{c.Name}.Sanitize()
		fieldDef := columnName + " " + pgType
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.CreationOptions.PrimaryKey {
			primaryKeys = append(primaryKeys, c.Name)
		}
	}

	if len(primaryKeys) > 0 {
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	} else {
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(table.Name)
		sb.WriteString("_cqpk PRIMARY KEY (_cq_id)")
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

func (p *Client) getPgTableColumns(ctx context.Context, tableName string) (*pgTableColumns, error) {
	tc := pgTableColumns{
		name: tableName,
	}
	rows, err := p.conn.Query(ctx, sqlSelectColumnTypes, p.currentSchemaName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var typ string
		if err := rows.Scan(&name, &typ); err != nil {
			return nil, err
		}
		tc.columns = append(tc.columns, pgColumn{
			name: strings.ToLower(name),
			typ:  strings.ToLower(typ),
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (p *Client) getPgTablePrimaryKeys(ctx context.Context, tableName string) (*pgTablePrimaryKeys, error) {
	pks := pgTablePrimaryKeys{
		name: tableName,
	}
	rows, err := p.conn.Query(ctx, sqlSelectPrimaryKeys, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		pks.columns = append(pks.columns, strings.ToLower(column))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &pks, nil
}

func (p *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	var err error
	var pgColumns *pgTableColumns
	var pgPKs *pgTablePrimaryKeys
	// create the new column as it doesn't exist
	tableName := pgx.Identifier{table.Name}.Sanitize()
	if pgColumns, err = p.getPgTableColumns(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns types: %w", table.Name, err)
	}

	if pgPKs, err = p.getPgTablePrimaryKeys(ctx, table.Name); err != nil {
		return fmt.Errorf("failed to get table %s columns primary keys: %w", table.Name, err)
	}
	reCreatePrimaryKeys := false

	for _, c := range table.Columns {
		columnName := pgx.Identifier{c.Name}.Sanitize()
		columnType, err := SchemaTypeToPg(c.Type)
		if err != nil {
			return fmt.Errorf("failed to convert schema type %s to postgresql type: %w", c.Type, err)
		}
		pgColumn := pgColumns.getPgColumn(c.Name)

		switch {
		case pgColumn == nil:
			p.logger.Info().Str("table", table.Name).Str("column", c.Name).Msg("Column doesn't exist, creating")

			sql := "alter table " + tableName + " add column " + columnName + " " + columnType
			if c.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			if _, err := p.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", c.Name, table.Name, err)
			}
		case pgColumn.typ != columnType:
			p.logger.Info().Str("table", table.Name).Str("column", c.Name).Str("old_type", pgColumn.typ).Str("new_type", columnType).Msg("Column exist but type is different, re-creating")
			// column exists but type is different

			// if this column contains primary key we will need to recreate the primary key
			if c.CreationOptions.PrimaryKey {
				reCreatePrimaryKeys = true
			}
			sql := "alter table " + tableName + " drop column " + columnName
			// right now we will drop the column and re-create. in the future we will have an option to automigrate
			if _, err := p.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to drop column %s on table %s: %w", c.Name, table.Name, err)
			}
			sql = "alter table " + tableName + " add column " + columnName + " " + columnType
			if _, err := p.conn.Exec(ctx, sql); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", c.Name, table.Name, err)
			}
		default:
			// column exists and type is the same but constraint might differ
			if pgPKs.columnExist(c.Name) != c.CreationOptions.PrimaryKey {
				p.logger.Info().Str("table", table.Name).Str("column", c.Name).Bool("pk", c.CreationOptions.PrimaryKey).Msg("Column exist with different primary keys")
				reCreatePrimaryKeys = true
			}
		}
	}
	if reCreatePrimaryKeys {
		p.logger.Info().Str("table", table.Name).Msg("recreating primary keys")
		constraintName := pgx.Identifier{table.Name + "_cqpk"}.Sanitize()
		sql := "alter table " + tableName + " drop constraint if exists " + constraintName
		if _, err := p.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
		}
		sql = "alter table " + tableName + " add constraint " + constraintName + " primary key (" + strings.Join(table.PrimaryKeys(), ",") + ")"
		if _, err := p.conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
		}
	}
	return nil
}

// This is the responsibility of the CLI of the client to lock before running migration
func (p *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		p.logger.Info().Str("table", table.Name).Msg("Migrating table")
		if len(table.Columns) == 0 {
			p.logger.Info().Str("table", table.Name).Msg("Table with not columns, skiping")
			continue
		}
		tableExist := 0
		if err := p.conn.QueryRow(ctx, isTableExistSQL, table.Name).Scan(&tableExist); err != nil {
			return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
		}
		if tableExist == 0 {
			p.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist creating")
			if err := p.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		} else {
			p.logger.Info().Str("table", table.Name).Msg("Table exist, auto-migrating")
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
	for i := range values {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < len(values)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	constraintName := fmt.Sprintf("%s_cqpk", table)
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
	sql, values := upsert(table, data)
	_, err := p.conn.Exec(ctx, sql, values...)
	if err != nil {
		return fmt.Errorf("failed to insert data with sql '%s': %w", sql, err)
	}
	return nil
}

func (*Client) ExampleConfig() string {
	return `
connection_string: "postgresql://postgres:pass@localhost:5432/postgres"
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
		return "jsonb", nil
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
		return "", fmt.Errorf("unknown type %s", t)
	}
}

func (p *pgTablePrimaryKeys) columnExist(column string) bool {
	for _, c := range p.columns {
		if c == column {
			return true
		}
	}
	return false
}

func (p *pgTableColumns) getPgColumn(column string) *pgColumn {
	for _, c := range p.columns {
		if c.name == column {
			return &c
		}
	}
	return nil
}
