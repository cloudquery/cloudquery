package destinations

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type PostgreSqlSpec struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
}

type PostgreSqlPlugin struct {
	conn   *pgxpool.Pool
	logger zerolog.Logger
}

func (p *PostgreSqlPlugin) Configure(ctx context.Context, spec specs.DestinationSpec) error {
	var specPostgreSql PostgreSqlSpec
	if err := spec.Spec.Decode(&specPostgreSql); err != nil {
		return errors.Wrap(err, "failed to decode spec")
	}
	pgxConfig, err := pgxpool.ParseConfig(specPostgreSql.ConnectionString)
	if err != nil {
		return errors.Wrap(err, "failed to parse connection string")
	}
	// pgxConfig.ConnConfig.Logger = zerologadapter.NewLogger(opts.Logger)
	p.conn, err = pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return errors.Wrap(err, "failed to connect to postgresql")
	}
	return nil
}

func (p *PostgreSqlPlugin) Save(ctx context.Context, resources []*schema.Resource) error {
	for _, resource := range resources {
		sql, values, err := sq.Insert(resource.TableName()).Columns(resource.Columns()...).Values([]interface{}{""}).ToSql()
		if err != nil {
			return errors.Wrap(err, "failed to generate insert sql")
		}
		_, err = p.conn.Exec(ctx, sql, values...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgreSqlPlugin) DropTables(ctx context.Context, table []*schema.Table) error {
	return nil
}

func (p *PostgreSqlPlugin) CreateTables(ctx context.Context, table []*schema.Table) error {
	for _, t := range table {
		if len(t.Columns) == 0 {
			p.logger.Info().Str("table", t.Name).Msg("skipping table creation, no columns")
			continue
		}
		var sb strings.Builder
		sb.WriteString("CREATE TABLE IF NOT EXISTS ")
		sb.WriteString(t.Name)
		sb.WriteString(" (")
		totalColumns := len(t.Columns)
		for i, c := range t.Columns {
			pgType, err := SchemaTypeToPg(c.Type)
			if err != nil {
				return errors.Wrap(err, "failed to convert schema type to postgresql type")
			}
			sb.WriteString(fmt.Sprintf("%s %s", c.Name, pgType))
			if i != totalColumns-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString(")")
		_, err := p.conn.Exec(ctx, sb.String())
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgreSqlPlugin) GetExampleConfig(ctx context.Context) string {
	return `
connection_string: "postgresql://user:password@localhost:5432/dbname"
`
}

func SchemaTypeToPg(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "BOOLEAN", nil
	case schema.TypeSmallInt, schema.TypeInt, schema.TypeBigInt:
		return "INTEGER", nil
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
