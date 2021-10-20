package drift

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type DriftImpl struct {
	logger hclog.Logger
	config *BaseConfig

	// flags
	debug bool

	tfBackendName, tfMode, tfProvider string
	forceDeep                         bool
}

func New(logger hclog.Logger) *DriftImpl {
	return &DriftImpl{
		logger: logger,
	}
}

func (d *DriftImpl) ID() string {
	return "drift"
}

func (d *DriftImpl) Configure(ctx context.Context, config hcl.Body) error {
	p := NewParser("")

	theCfg, diags := p.Decode(config, nil)
	if diags.HasErrors() {
		return diags
	}

	d.config = theCfg
	return nil
}

func (d *DriftImpl) Execute(ctx context.Context, req *model.ExecuteRequest) *model.ExecutionResult {
	ret := &model.ExecutionResult{}

	rootCmd := &cobra.Command{
		Use:   "drift",
		Short: "Drift Module",
		Long:  "Drift Module",
		Args:  cobra.MinimumNArgs(1),
	}
	rootCmd.SetUsageTemplate(`Usage:
  cloudquery [CQ-PARAMS] module drift [COMMAND] -- [FLAGS]{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] -- --help" for more information about a command.{{end}}
`)
	rootCmd.PersistentFlags().BoolVar(&d.debug, "debug", false, "Show debug output")
	rootCmd.SetArgs(req.Args)

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Detect drifts",
		Long:  "Detect drifts between cloud provider and IaC",
		Run: func(cmd *cobra.Command, args []string) {
			cb, _ := json.Marshal(d.config)
			d.logger.Debug("executing with config", "config", string(cb), "request", req.String())
			res, err := d.run(ctx, req)
			ret.Result = res
			if err != nil {
				ret.Error = err.Error()
			}
		},
	}
	runCmd.Flags().StringVar(&d.tfBackendName, "tf-backend-name", "mylocal", "Set Terraform backend name")
	runCmd.Flags().StringVar(&d.tfMode, "tf-mode", "managed", "Set Terraform mode")
	runCmd.Flags().StringVar(&d.tfProvider, "tf-provider", "", "Set Terraform provider (defaults to cloud provider name)")
	runCmd.Flags().BoolVar(&d.forceDeep, "deep", false, "Force deep mode")
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		if ret.Error == "" {
			ret.Error = err.Error()
		}
	}

	return ret
}

func (d *DriftImpl) run(ctx context.Context, req *model.ExecuteRequest) (Results, error) {
	provs, err := req.Providers()
	if err != nil {
		return nil, err
	}

	var iacProv *cqproto.GetProviderSchemaResponse
	for _, p := range provs {
		if p.Name == "terraform" { // TODO add more iac provider names
			if iacProv != nil {
				return nil, fmt.Errorf("only single IAC provider is supported at a time")
			}
			iacProv = p
		}
	}
	if iacProv == nil {
		return nil, fmt.Errorf("no IAC provider detected, can't continue")
	}

	conn, err := req.Conn()
	if err != nil {
		return nil, fmt.Errorf("no connection: %w", err)
	}

	var resList Results

	for _, cfg := range d.config.Providers {
		if cfg.Name == iacProv.Name {
			continue
		}

		var found bool
		for _, prov := range provs {
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

			for _, resName := range resourceKeys {
				res := cfg.Resources[resName]
				if res == nil {
					continue // skipped
				}
				pr := prov.ResourceTables[resName]
				if pr == nil {
					d.logger.Warn("Skipping resource, not found in ResourceTables", "provider", prov.Name, "resource", resName)
					continue
				}

				iacData := res.IAC[iacProv.Name]
				if iacData == nil {
					d.logger.Debug("Skipping resource, iac provider not configured", "provider", prov.Name, "resource", resName, "iac_provider", iacProv.Name)
					continue
				}

				res.finalInterpret()

				d.logger.Info("Running for provider and resource", "provider", prov.Name+":"+resName, "table", pr.Name, "ids", res.Identifiers, "attributes", res.Attributes, "iac_name", iacData.Name, "iac_type", iacData.Type)

				// Drift per resource
				var dres *Result
				switch iacProv.Name {
				case "terraform":
					dres, err = d.driftTerraform(ctx, conn, prov.Name, pr, res, iacData)
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

func (d *DriftImpl) driftTerraform(ctx context.Context, conn *pgxpool.Conn, cloudName string, cloudTable *schema.Table, resData *ResourceConfig, iacData *IACConfig) (*Result, error) {
	res := &Result{
		IAC:       "Terraform",
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	// Get from IAC
	// SELECT i.instance_id, i.attributes from tf_resource_instances i JOIN tf_resources r ON r.cq_id=i.resource_id JOIN tf_data d ON d.cq_id=r.running_id WHERE d.backend_name='mylocal' AND r.provider='aws' AND r.mode='managed' AND r.type='aws_s3_bucket' AND r.name='s3_bucket';

	// Get from provider
	// SELECT account_id, region, name, arn FROM aws_s3_buckets;

	tfProvider := d.tfProvider
	if tfProvider == "" {
		tfProvider = cloudName
	}

	tfAttributes := make([]string, len(resData.Attributes))
	tfQueryItems := make([]string, len(resData.Attributes))
	for i, a := range resData.Attributes {
		if mapped := iacData.attributeMap[a]; mapped != "" {
			tfAttributes[i] = mapped
		} else {
			tfAttributes[i] = a
		}
		switch cloudTable.Column(a).Type {
		case schema.TypeString:
			tfQueryItems[i] = fmt.Sprintf(`COALESCE(i.attributes->>'%s','')`, tfAttributes[i])
		case schema.TypeJSON:
			tfQueryItems[i] = fmt.Sprintf(`(i.attributes->>'%s')::json`, tfAttributes[i])
		default:
			tfQueryItems[i] = fmt.Sprintf(`i.attributes->>'%s'`, tfAttributes[i])
		}
	}

	cloudQueryItems := make([]string, len(resData.Attributes))
	for i := range resData.Attributes {
		if cloudTable.Column(resData.Attributes[i]).Type == schema.TypeString {
			cloudQueryItems[i] = fmt.Sprintf(`COALESCE("c"."%s",'')`, resData.Attributes[i])
		} else {
			cloudQueryItems[i] = fmt.Sprintf(`"c"."%s"`, resData.Attributes[i])
		}
	}

	tfAttrQuery := goqu.L("JSONB_BUILD_ARRAY(" + strings.Join(tfQueryItems, ",") + ")")
	cloudAttrQuery := goqu.L("JSONB_BUILD_ARRAY(" + strings.Join(cloudQueryItems, ",") + ")")

	if len(resData.Attributes) == 0 {
		tfAttrQuery = goqu.L("''")
		cloudAttrQuery = goqu.L("''")
	}

	tfSelect := goqu.Dialect("postgres").From(goqu.T("tf_resource_instances").As("i")).
		Select("i.instance_id", tfAttrQuery.As("attlist")).
		Join(goqu.T("tf_resources").As("r"), goqu.On(goqu.Ex{"r.cq_id": goqu.I("i.resource_id")})).
		Join(goqu.T("tf_data").As("d"), goqu.On(goqu.Ex{"d.cq_id": goqu.I("r.running_id")})).
		Where(goqu.Ex{"d.backend_name": goqu.V(d.tfBackendName)}).
		Where(goqu.Ex{"r.provider": goqu.V(tfProvider)}).
		Where(goqu.Ex{"r.mode": goqu.V(d.tfMode)}).
		Where(goqu.Ex{"r.type": goqu.V(iacData.Type)}).
		Where(goqu.Ex{"r.name": goqu.V(iacData.Name)})

	deepMode := d.forceDeep || (resData.Deep != nil && *resData.Deep)

	var err error

	if !deepMode {
		// Get equal resources
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).Join(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select("tf.instance_id")
		res.Equal, err = d.queryIntoResourceList(ctx, conn, q, "equals", nil)
		if err != nil {
			return nil, err
		}
	}

	{
		// Get missing resources
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).LeftJoin(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select("tf.instance_id").Where(goqu.Ex{"c.cq_id": nil})
		res.Missing, err = d.queryIntoResourceList(ctx, conn, q, "missing", nil)
		if err != nil {
			return nil, err
		}
	}

	{
		// Get extra resources
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).LeftJoin(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select(goqu.I("c." + resData.Identifiers[0])).Where(goqu.Ex{"tf.instance_id": nil})
		res.Extra, err = d.queryIntoResourceList(ctx, conn, q, "extras", nil)
		if err != nil {
			return nil, err
		}
	}

	if deepMode {
		// Get different resources
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).LeftJoin(goqu.T("tf"),
			goqu.On(
				goqu.Ex{
					"tf.instance_id": goqu.I("c." + resData.Identifiers[0]),
				},
				goqu.L("? @> ?", goqu.I("tf.attlist"), cloudAttrQuery),
				goqu.L("? <@ ?", goqu.I("tf.attlist"), cloudAttrQuery),
			),
		).
			Select(goqu.I("c." + resData.Identifiers[0])).Where(goqu.Ex{"tf.instance_id": nil})

		res.Different, err = d.queryIntoResourceList(ctx, conn, q, "differs", append(res.Missing, res.Extra...))
		if err != nil {
			return nil, err
		}

		if d.debug && len(res.Different) > 0 {
			// get tf side
			sel := goqu.Dialect("postgres").From("tf").With("tf", tfSelect).Select(goqu.I("tf.instance_id").As("id"), goqu.I("tf.attlist").As("attlist")).Where(
				goqu.Ex{
					"tf.instance_id": res.Different.IDs(),
				})
			tfAttList, err := d.queryIntoAttributeList(ctx, conn, sel, "attlist-tf")
			if err != nil {
				return nil, err
			}

			// get cloud side
			sel = goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).Select(goqu.I("c."+resData.Identifiers[0]).As("id"), cloudAttrQuery.As("attlist")).Where(
				exp.NewBooleanExpression(exp.InOp, goqu.I("c."+resData.Identifiers[0]), res.Different.IDs()),
			)
			cloudAttList, err := d.queryIntoAttributeList(ctx, conn, sel, "attlist-cloud")
			if err != nil {
				return nil, err
			}

			makeTable := func(title string) *tablewriter.Table {
				fmt.Println(title)
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{strings.ToUpper(cloudName) + " EXPR", strings.ToUpper(cloudName) + " VAL", "TERRAFORM VAL", "TERRAFORM EXPR"})
				table.SetBorder(true)
				return table
			}

			for k, tfAttrs := range tfAttList {
				cloudAttrs, ok := cloudAttList[k]
				if !ok {
					continue // Resource exists only in TF. This is already handled by the "Missing" resource/check
				}
				table := makeTable(fmt.Sprintf("DIFF RESOURCE: %s", k))
				var (
					matchingAttr []string
					matchingVal  []string
				)
				for i := range tfAttrs {
					if reflect.DeepEqual(cloudAttrs[i], tfAttrs[i]) {
						table.Append([]string{
							cloudQueryItems[i],
							fmt.Sprintf("%v", cloudAttrs[i]),
							fmt.Sprintf("%v", tfAttrs[i]),
							tfQueryItems[i],
						})
					} else {
						matchingAttr = append(matchingAttr, `"`+resData.Attributes[i]+`"`)
						matchingVal = append(matchingVal, fmt.Sprintf("%v", cloudAttrs[i]))
					}
				}
				table.Render()
				fmt.Println("Matching attributes " + strings.Join(matchingAttr, ", "))
				table = tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"ATTRIBUTE", "MATCHING VALUE"})
				table.SetBorder(true)
				for i := range matchingAttr {
					table.Append([]string{strings.Trim(matchingAttr[i], `"`), matchingVal[i]})
				}
				table.Render()
			}
		}

	}

	if deepMode {
		// Get deepequal resources
		q := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).Join(goqu.T("tf"),
			goqu.On(
				goqu.Ex{
					"tf.instance_id": goqu.I("c." + resData.Identifiers[0]),
				},
				goqu.L("? @> ?", goqu.I("tf.attlist"), cloudAttrQuery),
				goqu.L("? <@ ?", goqu.I("tf.attlist"), cloudAttrQuery),
			),
		).
			Select("tf.instance_id")
		res.DeepEqual, err = d.queryIntoResourceList(ctx, conn, q, "deepequal", nil)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (d *DriftImpl) queryIntoResourceList(ctx context.Context, conn *pgxpool.Conn, sel *goqu.SelectDataset, what string, exclude ResourceList) (ResourceList, error) {
	query, args, err := sel.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu build(%s) failed: %w", what, err)
	}
	d.logger.Debug("generated query", "type", what, "query", query, "args", args)

	var list []string
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		return nil, fmt.Errorf("goqu select(%s) failed: %w", what, err)
	}

	exList := make(map[string]struct{}, len(exclude))
	for _, e := range exclude {
		exList[e.ID] = struct{}{}
	}

	ret := make([]*Resource, 0, len(list))
	for i := range list {
		if _, ok := exList[list[i]]; ok {
			continue // exclude
		}
		ret = append(ret, &Resource{
			ID: list[i],
		})
	}

	return ret, nil
}

func (d *DriftImpl) queryIntoAttributeList(ctx context.Context, conn *pgxpool.Conn, sel *goqu.SelectDataset, what string) (map[string][]interface{}, error) {
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

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
