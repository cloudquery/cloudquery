package drift

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/afero"
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

func (d *Drift) run(ctx context.Context, req *module.ExecuteRequest) (*Results, error) {
	iacProv, iacStates, err := readIACStates(string(iacTerraform), d.params.StateFiles)
	if err != nil {
		return nil, err
	}

	resList := &Results{
		ListManaged: d.params.ListManaged,
		Debug:       d.params.Debug,
	}

	for _, cfg := range d.config.Providers {
		schema, err := d.findProvider(cfg, req.Providers)
		if err != nil {
			return nil, err
		} else if schema == nil {
			continue
		}

		d.logger.Debug("Processing for provider", "provider", schema.Name, "config", cfg)

		resources := cfg.interpolatedResourceMap(iacProv, d.logger)

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

			d.logger.Debug("Running for provider and resource", "provider", schema.Name+":"+resName, "table", pr.Name, "ids", res.Identifiers, "attributes", res.Attributes, "iac_type", res.IAC[iacProv].Type)

			// Drift per resource
			var (
				dres *Result
				err  error
			)
			switch iacProv {
			case iacTerraform:
				dres, err = driftTerraform(ctx, d.logger, req.Conn, schema.Name, pr, resName, resources, res.IAC[iacProv], iacStates.([]*terraform.Data), d.params)
			default:
				err = fmt.Errorf("no suitable handler found for %q", iacProv)
			}
			if err != nil {
				return nil, fmt.Errorf("drift failed for (%s:%s): %w", schema.Name, resName, err)
			}

			dres.Provider = schema.Name
			dres.ResourceType = resName
			resList.Data = append(resList.Data, dres)
		}
	}

	resList.process()

	return resList, nil
}

func queryIntoResourceList(ctx context.Context, logger hclog.Logger, conn *pgxpool.Conn, sel *goqu.SelectDataset) (ResourceList, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build failed: %w", err)
	}
	logger.Trace("generated query", "query", query, "args", args)

	var list []struct {
		ID      *string       `db:"id"`
		AttList []interface{} `db:"attlist"`
	}
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UndefinedTable {
			// ERROR: relation %q does not exist
			return nil, fmt.Errorf("cloud provider tables don't exist: Did you run `cloudquery fetch`? %w", pgErr)
		}

		logger.Warn("query failed with error", "query", query, "args", args, "error", err)
		return nil, fmt.Errorf("goqu select failed: %w", err)
	}

	ret := make([]*Resource, 0, len(list))
	for i := range list {
		ret = append(ret, &Resource{
			ID: func(s *string) string {
				if s == nil {
					return "<null id>"
				}
				return *s
			}(list[i].ID),
			Attributes: list[i].AttList,
		})
	}

	return ret, nil
}

func handleSubresource(logger hclog.Logger, sel *goqu.SelectDataset, pr *traversedTable, resources map[string]*ResourceConfig, accountIDs []string) *goqu.SelectDataset {
	parentColumn := pr.ParentIDColumn()

	if parentColumn == "" {
		if pr.Parent != nil {
			logger.Error("parent set but no parentColumn for table", "table", pr.Table.Name)
		}

		if len(accountIDs) > 0 {
			accountIDColumn := pr.AccountIDColumn()

			if accountIDColumn != "" {
				sel = sel.Where(goqu.Ex{"c." + accountIDColumn: accountIDs})
			}
		}

		return sel
	}
	if pr.Parent == nil {
		logger.Warn("parentColumn set but no parent for table", "table", pr.Table.Name)
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
			logger.Warn("Found parent but no resourceConfig", "table", pr.Table.Name)
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

	if len(accountIDs) > 0 {
		accountIDColumn := pr.AccountIDColumn()

		sel = sel.Where(goqu.Ex{parentTableName + "." + accountIDColumn: accountIDs})
	}

	return sel
}

func handleFilters(sel *goqu.SelectDataset, res *ResourceConfig) *goqu.SelectDataset {
	for _, f := range res.Filters {
		sel = sel.Where(goqu.L(f))
	}

	return sel
}

var idRegEx = regexp.MustCompile(`(?ms)^\$\{sql:(.+?)\}$`)

const idSeparator = "|"

// handleIdentifiers returns an SQL expression given one or multiple identifiers. the `sql(<query>)` is also handled here.
// Given multiple identifiers, each of them are concatenated using the idSeparator
func handleIdentifiers(identifiers []string) (exp.Expression, error) {
	idLen := len(identifiers)
	if idLen == 0 {
		return nil, fmt.Errorf("no identifiers to match")
	}

	concatArgs := make([]string, 0, len(identifiers)*2)
	for i, id := range identifiers {
		usingVariable := false

		if ma := idRegEx.FindStringSubmatch(id); len(ma) == 2 {
			id = ma[1]
			usingVariable = true
		}

		if strings.Contains(id, "${") {
			return nil, fmt.Errorf("identifier %d still contains variable", i)
		}

		if !usingVariable && !strings.Contains(id, ".") {
			id = "c." + `"` + id + `"`
		}

		concatArgs = append(concatArgs, id, "'"+idSeparator+"'")
	}

	if idLen == 1 {
		return goqu.L(concatArgs[0] + " AS id"), nil
	}

	return goqu.L("CONCAT(" + strings.Join(concatArgs[:len(concatArgs)-1], ",") + ") AS id"), nil
}

func readIACStates(iacID string, stateFiles []string) (iacProvider, interface{}, error) {
	if len(stateFiles) == 0 {
		return "", nil, fmt.Errorf("state files for %s not specified", iacID)
	}

	switch iacProvider(iacID) {
	case iacTerraform:
		ret := make([]*terraform.Data, len(stateFiles))

		fs := afero.NewOsFs()
		for idx, fn := range stateFiles {
			fh, err := fs.Open(fn)
			if err != nil {
				return "", nil, err
			}
			data, err := terraform.LoadState(fh)
			_ = fh.Close()
			if err != nil {
				return "", nil, fmt.Errorf("parse %s: %w", fn, err)
			}
			ret[idx] = data
		}
		return iacTerraform, ret, nil
	default:
		return "", nil, fmt.Errorf("unknown IAC %q", iacID)
	}
}

// Make sure we satisfy the interface
var _ module.Module = (*Drift)(nil)
