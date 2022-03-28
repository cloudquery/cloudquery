package drift

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift/terraform"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	cqschema "github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

	protoVersion = 1
)

func (i iacProvider) String() string {
	switch i {
	case iacTerraform:
		return "Terraform"
	case iacCloudformation:
		return "Cloudformation"
	default:
		return "unknown"
	}
}

func New(logger hclog.Logger) *Drift {
	return &Drift{
		logger: logger,
	}
}

func (d *Drift) ID() string {
	return "drift"
}

func (d *Drift) ProtocolVersions() []uint32 {
	return []uint32{protoVersion}
}

func (d *Drift) Configure(ctx context.Context, info module.Info, runParams module.ModuleRunParams) error {
	d.params = runParams.(RunParams)

	builtin, err := d.readBaseConfig(info.ProtocolVersion, info.ProviderData)
	if err != nil {
		return fmt.Errorf("builtin config failed: %w", err)
	}

	d.config, err = d.readProfileConfig(builtin, info.UserConfig)
	if err != nil {
		return fmt.Errorf("read config failed: %w", err)
	}

	return nil
}

func (d *Drift) Execute(ctx context.Context, req *module.ExecuteRequest) *module.ExecutionResult {
	ret := &module.ExecutionResult{}
	ret.Result, ret.Error = d.run(ctx, req)
	if ret.Error != nil {
		ret.ErrorMsg = ret.Error.Error()
	}

	return ret
}

func (d *Drift) ExampleConfig(providers []string) string {
	hasAws := false
	for i := range providers {
		if providers[i] == "aws" {
			hasAws = true
			break
		}
	}
	if !hasAws {
		return ""
	}

	return `// drift configuration block
drift "drift-example" {
  // state block defines from where to access the state
  terraform {
    // backend: "local" or "s3"
    backend  = "local"

    // local backend options
    // files: list of tfstate files
    files = [ "/path/to.tfstate" ]

    // s3 backend options
    // bucket   = "<tfstate bucket>"
    // keys     = [ "<tfstate key>" ]
    // region   = "us-east-1"
    // role_arn = ""
  }

    // provider "aws" {
    //   account_ids      = ["123456789"]
    //   check_resources   = ["ec2.instances:*"]
    //   ignore_resources = ["ec2.instances:i-123456789", "aws_cloudwatchlogs_filters:*"]
    // }
}`
}

func (d *Drift) readBaseConfig(version uint32, providerData map[string]cqproto.ModuleInfo) (*BaseConfig, error) {
	if version != protoVersion {
		return nil, fmt.Errorf("unsupported module protocol version %d", version)
	}

	configRaw, diags := hclparse.NewParser().ParseHCL(builtinConfig, "")
	if diags.HasErrors() {
		return nil, diags
	}

	content, diags := configRaw.Body.Content(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type: "config",
			},
		},
	})
	if diags.HasErrors() {
		return nil, diags
	}
	if l := len(content.Blocks); l != 1 {
		return nil, fmt.Errorf("unexpected number of blocks (%d)", l)
	}

	p := NewParser("")
	cfg, diags := p.Decode(content.Blocks[0].Body, "", nil)
	if diags.HasErrors() {
		return nil, diags
	}

	for provider, modInfo := range providerData {
		hc, diags := module.GetCombinedHCL(modInfo)
		provCfg, diags := p.Decode(hc, provider, diags)
		if diags.HasErrors() {
			return nil, diags
		}
		if l := len(provCfg.Providers); l != 1 {
			return nil, fmt.Errorf("unexpected number of provider blocks (%s: %d)", provider, l)
		}
		cfg.Providers = append(cfg.Providers, provCfg.Providers...)
	}

	if diags.HasErrors() {
		return nil, diags
	}

	return cfg, nil
}

func (d *Drift) readProfileConfig(base *BaseConfig, body hcl.Body) (*BaseConfig, error) {
	p := NewParser("")

	if body == nil {
		if diags := p.interpret(base); diags.HasErrors() {
			return nil, diags
		}
		return base, nil
	}

	cfg, diags := p.Decode(body, "", nil)
	if diags.HasErrors() {
		return nil, diags
	}

	// keep base config and apply profile config into it, as we don't know which provider version from the builtin config is going to be selected
	base.Terraform = cfg.Terraform // always override this one

	if cfg.WildProvider != nil {
		base.WildProvider.applyWildProvider(cfg.WildProvider)

		if cfg.WildProvider.WildResource != nil {
			cfg.WildProvider.WildResource.applyWildResource(base.WildProvider.WildResource)
			base.WildProvider.WildResource = cfg.WildProvider.WildResource
		}
	}

	for _, prov := range base.Providers {
		cp := cfg.FindProvider(prov.Name)
		if cp == nil {
			continue
		}
		prov.applyWildProvider(cp)

		if cp.WildResource != nil {
			cp.WildResource.applyWildResource(prov.WildResource)
			prov.WildResource = cp.WildResource
		}

		for resName, res := range prov.Resources {
			if cres, ok := cp.Resources[resName]; ok {
				cres.applyWildResource(res)
				prov.Resources[resName] = cres
			}
		}
	}

	if diags := p.interpret(base); diags.HasErrors() {
		return nil, diags
	}

	return base, nil
}

func (d *Drift) run(ctx context.Context, req *module.ExecuteRequest) (*Results, error) {
	iacProv, iacStates, err := readIACStates(d.logger, string(iacTerraform), d.config.Terraform, d.params.StateFiles)
	if err != nil {
		return nil, err
	}

	resList := &Results{
		IACName:     iacProv.String(),
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
		if d.params.Debug {
			listUnimplementedResources(d.logger, resources, schema)
		}

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
				dres, err = driftTerraform(ctx, d.logger, req.Conn, schema.Name, pr, resName, resources, res.IAC[iacProv], iacStates.([]*terraform.Data), d.params, cfg.AccountIDs)
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

func listUnimplementedResources(logger hclog.Logger, resources map[string]*ResourceConfig, provSchema *cqproto.GetProviderSchemaResponse) {
	var (
		res, subRes []string
	)
	for nm, tabl := range provSchema.ResourceTables {
		if _, found := resources[nm]; !found {
			res = append(res, nm)
		}
		for _, rel := range tabl.Relations {
			subRes = append(subRes, listUnimplementedResourcesInner(resources, nm+":", rel)...)
		}
	}
	sort.Strings(res)
	sort.Strings(subRes)
	logger.Debug("Not implemented resources", "list", res)
	logger.Debug("Not implemented subresources", "list", subRes)
}

func listUnimplementedResourcesInner(resources map[string]*ResourceConfig, upper string, t *cqschema.Table) []string {
	var ret []string
	if _, found := resources[t.Name]; !found {
		ret = append(ret, upper+t.Name)
	}
	for _, rel := range t.Relations {
		ret = append(ret, listUnimplementedResourcesInner(resources, upper+t.Name+":", rel)...)
	}
	return ret
}

func queryIntoResourceList(ctx context.Context, logger hclog.Logger, conn execution.QueryExecer, sel *goqu.SelectDataset) (ResourceList, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build failed: %w", err)
	}
	logger.Trace("generated query", "query", query, "args", args)

	var list []struct {
		ID      *string           `db:"id"`
		AttList []interface{}     `db:"attlist"`
		Tags    map[string]string `db:"tags"`
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
			Tags:       list[i].Tags,
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

func readIACStates(logger hclog.Logger, iacID string, tf *TerraformSourceConfig, stateFiles []string) (iacProvider, interface{}, error) {
	if iacProvider(iacID) != iacTerraform {
		return "", nil, fmt.Errorf("unknown IAC %q", iacID)
	}

	if len(stateFiles) == 0 {
		// no override: read TF config here
		if tf == nil {
			return "", nil, fmt.Errorf("terraform configuration not found: either specify state files or edit config.hcl")
		}
		switch tf.Backend {
		case TFLocal:
			stateFiles = tf.Files
		case TFS3:
			states, err := loadIACStatesFromS3(iacID, tf.Bucket, tf.Keys, tf.Region, tf.RoleARN)
			if err != nil {
				return "", nil, err
			}
			return iacTerraform, states, nil
		default:
			return "", nil, fmt.Errorf("unsupported backend")
		}
	}

	if len(stateFiles) == 0 {
		return "", nil, fmt.Errorf("state files for %s not specified", iacID)
	}

	ret := make([]*terraform.Data, 0, len(stateFiles))

	fs := afero.NewOsFs()
	for _, fnGlob := range stateFiles {
		matches, err := afero.Glob(fs, fnGlob)
		if err != nil {
			return "", nil, err
		}
		for _, fn := range matches {
			fh, err := fs.Open(fn)
			if err != nil {
				return "", nil, err
			}
			data, err := terraform.LoadState(fh)
			_ = fh.Close()
			if err != nil {
				return "", nil, fmt.Errorf("parse %s: %w", fn, err)
			}
			if ok, err := terraform.ValidateStateVersion(data); err != nil {
				if !ok {
					return "", nil, fmt.Errorf("validate %s: %w", fn, err)
				}
				logger.Warn("ValidateStateVersion", "warning", err.Error())
			} else if !ok {
				return "", nil, fmt.Errorf("validate %s: failed", fn)
			}

			ret = append(ret, data)
		}
	}

	if len(ret) == 0 {
		return "", nil, fmt.Errorf("no matches for specified %s state patterns", iacID)
	}

	return iacTerraform, ret, nil
}

// Make sure we satisfy the interface
var _ module.Module = (*Drift)(nil)
