package state

import (
	"context"
	"errors"

	"github.com/cloudquery/cloudquery/cli/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
)

// Provider keeps track of installed providers
type Provider struct {
	Source  string `db:"source"`
	Name    string `db:"name"`
	Version string `db:"version"`

	VMajor int    `db:"v_major"`
	VMinor int    `db:"v_minor"`
	VPatch int    `db:"v_patch"`
	VPre   string `db:"v_pre"`
	VMeta  string `db:"v_meta"`

	Tables     tableList `db:"tables"`
	Signatures stringMap `db:"signatures"`

	ParsedVersion *version.Version `db:"-"`
}

type Tx struct {
	execution.TXQueryExecer

	c *Client
}

func (p *Provider) Registry() registry.Provider {
	return registry.Provider{
		Source:  p.Source,
		Name:    p.Name,
		Version: p.Version,
	}
}

// GetProvider gets state about given provider, or returns nil, nil.
func (c *Client) GetProvider(ctx context.Context, p registry.Provider) (*Provider, error) {
	q := goqu.Dialect("postgres").
		Select("source", "name", "version", "v_major", "v_minor", "v_patch", "v_pre", "v_meta", "tables", "signatures").
		From("cloudquery.providers").
		Where(goqu.Ex{"source": p.Source, "name": p.Name}).
		Limit(1)
	sql, _, err := q.ToSQL()
	if err != nil {
		return nil, err
	}
	var data Provider
	err = pgxscan.Get(ctx, c.db, &data, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if v, err := version.NewVersion(data.Version); err == nil {
		data.ParsedVersion = v
	}

	return &data, nil
}

func (c *Client) ProviderSync(ctx context.Context) (*Tx, error) {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{
		TXQueryExecer: tx,
		c:             c,
	}, nil
}

// ProviderFromRegistry returns a Provider struct with info filled from a registry.Provider
func ProviderFromRegistry(r registry.Provider) *Provider {
	p := &Provider{
		Source:     r.Source,
		Name:       r.Name,
		Version:    r.Version,
		Tables:     make(tableList),
		Signatures: make(stringMap),
	}
	if r.Version != "" {
		if ver, err := version.NewVersion(r.Version); err == nil {
			sg := ver.Segments()
			if len(sg) > 0 {
				p.VMajor = sg[0]
			}
			if len(sg) > 1 {
				p.VMinor = sg[1]
			}
			if len(sg) > 2 {
				p.VPatch = sg[2]
			}
			p.VPre, p.VMeta, p.ParsedVersion = ver.Prerelease(), ver.Metadata(), ver
		}
	}
	return p
}

// InstallProvider saves state about given provider
func (t *Tx) InstallProvider(ctx context.Context, p *Provider) error {
	q := goqu.Dialect("postgres").Insert("cloudquery.providers").Rows(p)
	sql, args, err := q.ToSQL()
	if err != nil {
		return err
	}
	return t.TXQueryExecer.Exec(ctx, sql, args...)
}

// UninstallProvider removes state about given provider
func (t *Tx) UninstallProvider(ctx context.Context, p registry.Provider) error {
	q := goqu.Dialect("postgres").Delete("cloudquery.providers").Where(goqu.Ex{"source": p.Source, "name": p.Name})
	sql, args, err := q.ToSQL()
	if err != nil {
		return err
	}
	return t.TXQueryExecer.Exec(ctx, sql, args...)
}
