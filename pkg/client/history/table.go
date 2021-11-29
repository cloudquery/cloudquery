package history

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CreateHyperTableResult struct {
	HypertableId int    `db:"hypertable_id"`
	SchemaName   string `db:"schema_name"`
	TableName    string `db:"table_name"`
	Created      bool   `db:"created"`
}

type TableCreator struct {
	log hclog.Logger
	cfg *Config
}

func NewHistoryTableCreator(cfg *Config, l hclog.Logger) (*TableCreator, error) {
	return &TableCreator{
		l,
		cfg,
	}, nil
}

func (h TableCreator) CreateTable(ctx context.Context, conn *pgxpool.Conn, t, p *schema.Table) error {
	sql, err := h.buildTableSQL(t, p)
	if err != nil {
		return err
	}
	h.log.Debug("creating table if not exists", "table", t.Name)
	if _, err := conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	if err := h.createHyperTable(ctx, t, p, conn); err != nil {
		return fmt.Errorf("failed to created hypertable for table: %s: %w", t.Name, err)
	}

	if _, err := conn.Exec(ctx, fmt.Sprintf(createTableView, t.Name)); err != nil {
		return fmt.Errorf("failed to created view for table: %s: %w", t.Name, err)
	}

	if p != nil {
		if err := h.buildCascadeTrigger(ctx, conn, t, p); err != nil {
			return fmt.Errorf("table build %s failed: %w", t.Name, err)
		}
	}

	// Create relation tables
	for _, r := range t.Relations {
		h.log.Debug("creating table relation", "table", r.Name)
		if err := h.CreateTable(ctx, conn, r, t); err != nil {
			return err
		}
	}

	return nil
}

func (h TableCreator) createHyperTable(ctx context.Context, t, p *schema.Table, conn *pgxpool.Conn) error {
	var hyperTable CreateHyperTableResult
	tName := fmt.Sprintf(`"history"."%s"`, t.Name)
	if err := pgxscan.Get(ctx, conn, &hyperTable, fmt.Sprintf(createHyperTable, h.cfg.TimeInterval), tName); err != nil {
		return fmt.Errorf("failed to create hypertable: %w", err)
	}
	h.log.Debug("created hyper table for table", "table", hyperTable.TableName, "id", hyperTable.HypertableId, "created", hyperTable.Created)
	if p != nil {
		return nil
	}
	if _, err := conn.Exec(ctx, fmt.Sprintf(dataRetentionPolicy, h.cfg.Retention), tName); err != nil {
		return err
	}
	h.log.Debug("created data retention policy", "table", hyperTable.TableName, "days", h.cfg.Retention)
	return nil
}

// TODO: create unique index on fetch_date parent_cq_id
func (h TableCreator) buildCascadeTrigger(ctx context.Context, conn *pgxpool.Conn, t, p *schema.Table) error {
	c := h.findParentIdColumn(t)
	if c == nil {
		return fmt.Errorf("failed to find parent cq id column for %s", t.Name)
	}
	_, err := conn.Exec(ctx, "SELECT history.build_trigger($1, $2, $3);", p.Name, t.Name, c.Name)
	if err != nil {
		return fmt.Errorf("failed to create trigger: %w", err)
	}
	return nil
}

func (h TableCreator) findParentIdColumn(t *schema.Table) *schema.Column {
	for _, c := range t.Columns {
		if c.Meta().Resolver != nil && c.Meta().Resolver.Name == "ParentIdResolver" {
			return &c
		}
	}
	// Support old school columns instead of meta, this is backwards compatibility for providers using SDK prior v0.5.0
	for _, c := range t.Columns {
		if strings.HasSuffix(c.Name, "cq_id") && c.Name != "cq_id" {
			return &c
		}
	}

	return nil
}

func (h TableCreator) buildTableSQL(table, _ *schema.Table) (string, error) {
	// Build SQL to create a table.
	ctb := sqlbuilder.CreateTable(fmt.Sprintf("history.%s", table.Name)).IfNotExists()
	var uniques []string
	for _, c := range schema.GetDefaultSDKColumns() {
		ctb.Define(c.Name, schema.GetPgTypeFromType(c.Type))
		if c.CreationOptions.Unique {
			uniques = append(uniques, c.Name)
		}
	}
	// TODO check fetch_date not defined
	ctb.Define("fetch_date", schema.GetPgTypeFromType(schema.TypeTimestamp))
	h.buildColumns(ctb, table.Columns)
	for _, s := range uniques {
		ctb.Define(fmt.Sprintf("UNIQUE (fetch_date, %s)", s))
	}
	// TODO: check fetch_date not in PKs
	allKeys := append([]string{"fetch_date"}, table.PrimaryKeys()...)
	ctb.Define(fmt.Sprintf("constraint %s_pk primary key(%s)", schema.TruncateTableConstraint(table.Name), strings.Join(allKeys, ",")))

	sql, _ := ctb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	h.log.Trace("creating table if not exists", "table", table.Name)
	return sql, nil
}

func (h TableCreator) buildColumns(ctb *sqlbuilder.CreateTableBuilder, cc []schema.Column) {
	for _, c := range cc {
		defs := []string{strconv.Quote(c.Name), schema.GetPgTypeFromType(c.Type)}
		if c.CreationOptions.Unique {
			defs = []string{strconv.Quote(c.Name), schema.GetPgTypeFromType(c.Type), "unique"}
		}
		ctb.Define(defs...)
	}
}
