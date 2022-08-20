package postgresql

import (
	"context"
	"fmt"
	"strings"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type PostgreSqlSpec struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
}

type Client struct {
	conn   *pgxpool.Pool
	logger zerolog.Logger
	m      sync.Mutex
}

func NewClient(logger zerolog.Logger) *Client {
	return &Client{
		logger: logger,
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

	if err := spec.UnmarshalSpec(&specPostgreSql); err != nil {
		return fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}

	pgxConfig, err := pgxpool.ParseConfig(specPostgreSql.ConnectionString)
	if err != nil {
		return errors.Wrap(err, "failed to parse connection string")
	}
	l := zerologadapter.NewLogger(p.logger)
	pgxConfig.ConnConfig.Logger = l
	pgxConfig.ConnConfig.LogLevel = pgx.LogLevelWarn
	p.conn, err = pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return errors.Wrap(err, "failed to connect to postgresql")
	}
	return nil
}

const isTableExistSQL = "select count(*) from information_schema.tables where table_name = $1"

func (p *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(table.Name)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)
	for i, c := range table.Columns {
		pgType, err := SchemaTypeToPg(c.Type)
		if err != nil {
			return errors.Wrap(err, "failed to convert schema type to postgresql type")
		}
		sb.WriteString(fmt.Sprintf("\"%s\" %s", c.Name, pgType))
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")
	_, err := p.conn.Exec(ctx, sb.String())
	if err != nil {
		return err
	}
	return nil
}

// This is the responsability of the CLI of the client to lock before running migration
func (p *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		if len(table.Columns) == 0 {
			p.logger.Info().Str("table", table.Name).Msg("skipping table creation, no columns")
			continue
		}
		tableExist := 0
		if err := p.conn.QueryRow(ctx, isTableExistSQL, table.Name).Scan(&tableExist); err != nil {
			return errors.Wrap(err, "failed to query information_schema.tables")
		}
		if tableExist == 0 {
			if err := p.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		}
		if err := p.Migrate(ctx, table.Relations); err != nil {
			return err
		}
	}
	return nil
}

func (p *Client) Write(ctx context.Context, resource *schema.Resource) error {
	columns := make([]string, 0, len(resource.Data))
	values := make([]interface{}, 0, len(resource.Data))
	for c, v := range resource.Data {
		columns = append(columns, `"`+c+`"`)
		values = append(values, v)
	}

	sql, values, err := sq.Insert(resource.TableName).Columns(columns...).Values(values...).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed to generate insert sql '%s': %w", sql, err)
	}
	_, err = p.conn.Exec(ctx, sql, values...)
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

func SchemaTypeToPg(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "BOOLEAN", nil
	case schema.TypeSmallInt, schema.TypeInt, schema.TypeBigInt:
		return "BIGINT", nil
	case schema.TypeFloat:
		return "REAL", nil
	case schema.TypeUUID:
		return "UUID", nil
	case schema.TypeString:
		return "TEXT", nil
	case schema.TypeByteArray:
		return "BYTEA", nil
	case schema.TypeStringArray:
		return "TEXT[]", nil
	case schema.TypeTimestamp:
		return "TIMESTAMP", nil
	case schema.TypeJSON:
		return "JSON", nil
	case schema.TypeUUIDArray:
		return "UUID[]", nil
	case schema.TypeInetArray:
		return "INET[]", nil
	case schema.TypeCIDR:
		return "CIDR", nil
	case schema.TypeCIDRArray:
		return "CIDR[]", nil
	case schema.TypeMacAddr:
		return "MACADDR", nil
	case schema.TypeMacAddrArray:
		return "MACADDR[]", nil
	case schema.TypeInet:
		return "INET", nil
	case schema.TypeIntArray:
		return "INET[]", nil
	default:
		return "", errors.Errorf("unsupported schema type: %s", t)
	}
}
