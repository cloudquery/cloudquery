package drift

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Drift struct {
	logger hclog.Logger
	config *BaseConfig

	params RunParams

	tableMap map[string]resourceMap // one map per provider, initiated on first use
}

type resourceMap map[string]*traversedTable // table names vs. table definitions

type iacProvider string

const (
	iacTerraform      iacProvider = "terraform"
	iacCloudformation iacProvider = "cloudformation"
)

func New(logger hclog.Logger) *Drift {
	return &Drift{
		logger: logger,
	}
}

func (d *Drift) ID() string {
	return "drift"
}

func (d *Drift) Configure(ctx context.Context, config hcl.Body) error {
	p := NewParser("")

	cfg, diags := p.Decode(config, nil)
	if diags.HasErrors() {
		return diags
	}

	d.config = cfg
	return nil
}

func (d *Drift) Execute(ctx context.Context, req *module.ExecuteRequest) *module.ExecutionResult {
	d.params = req.Params.(RunParams)

	ret := &module.ExecutionResult{}
	ret.Result, ret.Error = d.run(ctx, req)
	if ret.Error != nil {
		ret.ErrorMsg = ret.Error.Error()
	}

	return ret
}

func (d *Drift) run(ctx context.Context, req *module.ExecuteRequest) (Results, error) {
	iacProv, err := getIACProvider(req.Providers)
	if err != nil {
		return nil, err
	}

	var resList Results

	for _, cfg := range d.config.Providers {
		schema, err := d.findProvider(cfg, req.Providers, iacProv.Name)
		if err != nil {
			return nil, err
		} else if schema == nil {
			continue
		}

		d.logger.Debug("Processing for provider", "provider", schema.Name, "config", cfg)

		resources := cfg.interpolatedResourceMap(iacProv.Name, d.logger)

		// Always process in the same order so both results and error messages are consistent
		for _, resName := range cfg.resourceKeys() {
			res := resources[resName]
			if res == nil {
				continue // skipped
			}
			pr := d.lookupResource(resName, schema)
			if pr == nil {
				d.logger.Warn("Skipping resource, lookup failed", "resource", resName)
				continue
			}

			d.logger.Debug("Running for provider and resource", "provider", schema.Name+":"+resName, "table", pr.Name, "ids", res.Identifiers, "attributes", res.Attributes, "iac_type", res.IAC[iacProv.Name].Type)

			// Drift per resource
			var (
				dres *Result
				err  error
			)
			switch iacProvider(iacProv.Name) {
			case iacTerraform:
				dres, err = d.driftTerraform(ctx, req.Conn, schema.Name, pr, resName, resources, res.IAC[iacProv.Name])
			default:
				err = fmt.Errorf("no suitable handler found for %q", iacProv.Name)
			}
			if err != nil {
				return nil, fmt.Errorf("drift failed for (%s:%s): %w", schema.Name, resName, err)
			}

			dres.Provider = schema.Name
			dres.ResourceType = resName
			resList = append(resList, dres)
		}

	}

	return resList, nil
}

func (d *Drift) queryIntoResourceList(ctx context.Context, conn *pgxpool.Conn, sel *goqu.SelectDataset, what string, exclude ResourceList) (ResourceList, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build(%s) failed: %w", what, err)
	}
	d.logger.Trace("generated query", "type", what, "query", query, "args", args)

	var list []*string
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UndefinedTable {
			// ERROR: relation %q does not exist
			return nil, fmt.Errorf("terraform provider tables don't exist: Did you run `cloudquery fetch`?%q", pgErr.TableName)
		}

		d.logger.Warn("query failed with error", "query", query, "args", args, "type", what, "error", err)
		return nil, fmt.Errorf("goqu select(%s) failed: %w", what, err)
	}

	exList := make(map[string]struct{}, len(exclude))
	for _, e := range exclude {
		exList[e.ID] = struct{}{}
	}

	ret := make([]*Resource, 0, len(list))
	for i := range list {
		if list[i] == nil {
			ret = append(ret, &Resource{
				ID: "<null id>",
			})
			continue
		}

		if _, ok := exList[*list[i]]; ok {
			continue // exclude
		}
		ret = append(ret, &Resource{
			ID: *list[i],
		})
	}

	return ret, nil
}

func (d *Drift) queryIntoAttributeList(ctx context.Context, conn *pgxpool.Conn, sel *goqu.SelectDataset, what string) (map[string][]interface{}, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build(%s) failed: %w", what, err)
	}
	d.logger.Trace("generated query", "type", what, "query", query, "args", args)

	var list []struct {
		ID      string        `db:"id"`
		AttList []interface{} `db:"attlist"`
	}
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		return nil, fmt.Errorf("goqu select(%s) failed: %w", what, err)
	}

	ret := make(map[string][]interface{}, len(list))
	for i := range list {
		ret[list[i].ID] = list[i].AttList
	}
	return ret, nil
}

func (d *Drift) handleSubresource(sel *goqu.SelectDataset, pr *traversedTable, resources map[string]*ResourceConfig) *goqu.SelectDataset {
	parentColumn := pr.ParentIDColumn()

	if parentColumn == "" {
		if pr.Parent != nil {
			d.logger.Error("parent set but no parentColumn for table", "table", pr.Table.Name)
		}

		if len(d.params.AccountIDs) > 0 {
			accountIDColumn := pr.AccountIDColumn()

			if accountIDColumn != "" {
				sel = sel.Where(goqu.Ex{"c." + accountIDColumn: d.params.AccountIDs})
			}
		}

		return sel
	}
	if pr.Parent == nil {
		d.logger.Warn("parentColumn set but no parent for table", "table", pr.Table.Name)
		return sel
	}

	// Join all parents up the chain, topmost parent has account_id

	parentCounter := 0
	parentTableName := "parent"
	childTableName := "c"
	var res *ResourceConfig
	for pr.Parent != nil {
		res = resources[pr.Name]
		if res == nil {
			d.logger.Warn("Found parent but no resourceConfig", "table", pr.Table.Name)
			return sel // FIXME we're skipping the account_id filter here by returning
		}

		if parentCounter > 0 {
			parentTableName = fmt.Sprintf("parent%d", parentCounter)
		}

		sel = sel.Join(
			goqu.T(pr.Parent.Name).As(parentTableName),
			goqu.On(
				goqu.L("? = ?",
					goqu.I(parentTableName+".cq_id"),
					goqu.I(childTableName+"."+parentColumn),
				),
			),
		)

		parentCounter++
		childTableName = parentTableName
		pr = pr.Parent
		parentColumn = pr.ParentIDColumn()
	}

	if len(d.params.AccountIDs) > 0 {
		accountIDColumn := pr.AccountIDColumn()

		sel = sel.Where(goqu.Ex{parentTableName + "." + accountIDColumn: d.params.AccountIDs})
	}

	return sel
}

func (d *Drift) handleFilters(sel *goqu.SelectDataset, res *ResourceConfig) *goqu.SelectDataset {
	for _, f := range res.Filters {
		sel = sel.Where(goqu.L(f))
	}

	return sel
}

var idRegEx = regexp.MustCompile(`(?ms)^\$\{sql:(.+?)\}$`)

func (d *Drift) handleIdentifier(identifiers []string) (exp.Expression, exp.Expression, error) {
	switch l := len(identifiers); {
	case l == 0:
		return nil, nil, fmt.Errorf("no identifiers to match")
	case l > 1:
		return nil, nil, fmt.Errorf("multiple identifiers not supported yet")
	}

	usingVariable := false

	if ma := idRegEx.FindStringSubmatch(identifiers[0]); len(ma) == 2 {
		identifiers[0] = ma[1]
		usingVariable = true
	}

	if strings.Contains(identifiers[0], "${") {
		return nil, nil, fmt.Errorf("identifier still contains variable")
	}

	if usingVariable {
		return goqu.L(identifiers[0]), goqu.L("(" + identifiers[0] + ") AS id"), nil
	}

	return goqu.I("c." + identifiers[0]), goqu.I("c." + identifiers[0]).As("id"), nil
}

func getIACProvider(provs []*cqproto.GetProviderSchemaResponse) (*cqproto.GetProviderSchemaResponse, error) {
	var iacProv *cqproto.GetProviderSchemaResponse
	for _, p := range provs {
		if p.Name == string(iacTerraform) {
			if iacProv != nil {
				return nil, fmt.Errorf("only single IAC provider is supported at a time")
			}
			iacProv = p
		}
	}
	if iacProv == nil {
		return nil, fmt.Errorf("no IAC provider detected, can't continue")
	}

	return iacProv, nil
}

// Make sure we satisfy the interface
var _ module.Module = (*Drift)(nil)
