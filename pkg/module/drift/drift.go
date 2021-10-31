package drift

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Drift struct {
	logger hclog.Logger
	config *BaseConfig

	params RunParams

	tableMap map[string]*provResource   // provider table names vs. table defs, initiated on first use
	resMap   map[string]*ResourceConfig // resource name vs. interpreted/finalized config
}

type provResource struct {
	*schema.Table
	Parent *provResource
}

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
	var err error
	ret.Result, err = d.run(ctx, req)
	if err != nil {
		ret.Error = err.Error()
	}

	return ret
}

func (d *Drift) run(ctx context.Context, req *module.ExecuteRequest) (Results, error) {
	var iacProv *cqproto.GetProviderSchemaResponse
	for _, p := range req.Providers {
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

	var resList Results

	for _, cfg := range d.config.Providers {
		if cfg.Name == iacProv.Name {
			continue
		}

		var found bool
		for _, prov := range req.Providers {
			ok, diags := d.applyProvider(cfg, prov)
			if diags.HasErrors() {
				return nil, diags
			}
			if !ok {
				continue
			}

			found = true

			d.logger.Debug("Processing for provider", "provider", prov.Name, "config", cfg)

			// Always process in the same order
			resourceKeys := make([]string, 0, len(cfg.Resources))
			for i := range cfg.Resources {
				resourceKeys = append(resourceKeys, i)
			}
			sort.Strings(resourceKeys)

			d.resMap = make(map[string]*ResourceConfig, len(resourceKeys))
			for _, resName := range resourceKeys {
				res := cfg.Resources[resName]
				if res == nil {
					continue // skipped
				}
				iacData := res.IAC[iacProv.Name]
				if iacData == nil {
					d.logger.Debug("Will skip resource, iac provider not configured", "provider", prov.Name, "resource", resName, "iac_provider", iacProv.Name)
					continue
				}

				res.finalInterpret()
				d.resMap[resName] = res
			}

			for _, resName := range resourceKeys {
				res := d.resMap[resName]
				if res == nil {
					continue // skipped
				}
				pr := d.lookupResource(resName, prov)
				if pr == nil {
					continue
				}

				iacData := res.IAC[iacProv.Name]

				d.logger.Info("Running for provider and resource", "provider", prov.Name+":"+resName, "table", pr.Name, "ids", res.Identifiers, "attributes", res.Attributes, "iac_type", iacData.Type)

				// Drift per resource
				var (
					dres *Result
					err  error
				)
				switch iacProvider(iacProv.Name) {
				case iacTerraform:
					dres, err = d.driftTerraform(ctx, req.Conn, prov.Name, pr, resName, iacData)
				default:
					return nil, fmt.Errorf("no suitable handler found for %q", iacProv.Name)
				}
				if err != nil {
					return nil, fmt.Errorf("drift failed for (%s:%s): %w", prov.Name, resName, err)
				} else if dres != nil {
					dres.Provider = prov.Name
					dres.ResourceType = resName
					resList = append(resList, dres)
				}
			}

			break
		}

		if !found {
			return nil, fmt.Errorf("no suitable provider found for %q", cfg.Name)
		}
	}

	return resList, nil
}

func (d *Drift) queryIntoResourceList(ctx context.Context, conn *pgxpool.Conn, sel *goqu.SelectDataset, what string, exclude ResourceList) (ResourceList, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build(%s) failed: %w", what, err)
	}
	d.logger.Debug("generated query", "type", what, "query", query, "args", args)

	var list []*string
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		// ERROR: relation %q does not exist
		if strings.Contains(err.Error(), "SQLSTATE 42P01") && strings.Contains(err.Error(), "does not exist") {
			return nil, fmt.Errorf("terraform provider tables don't exist: Did you run `cloudquery fetch`?")
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
	d.logger.Debug("generated query", "type", what, "query", query, "args", args)

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

func (d *Drift) handleSubresource(sel *goqu.SelectDataset, pr *provResource, res *ResourceConfig) *goqu.SelectDataset {
	if res.ParentMatch == "" {
		if len(d.params.AccountIDs) > 0 {
			sel = sel.Where(goqu.Ex{"c.account_id": d.params.AccountIDs})
		}

		return sel
	}
	if pr.Parent == nil {
		d.logger.Warn("parent_match set but no parent for table", "table", pr.Table.Name)
		return sel
	}

	// Join all parents up the chain, topmost parent has account_id

	parentCounter := 0
	parentTableName := "parent"
	childTableName := "c"
	for pr.Parent != nil {
		res = d.resMap[pr.Name]
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
					goqu.I(childTableName+"."+res.ParentMatch),
				),
			),
		)

		parentCounter++
		childTableName = parentTableName
		pr = pr.Parent
	}

	if len(d.params.AccountIDs) > 0 {
		sel = sel.Where(goqu.Ex{parentTableName + ".account_id": d.params.AccountIDs})
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

// Make sure we satisfy the interface
var _ module.Module = (*Drift)(nil)
