package drift

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DriftImpl struct {
	logger hclog.Logger

	config *BaseConfig
}

func New(logger hclog.Logger) *DriftImpl {
	return &DriftImpl{
		logger: logger,
	}
}

func (d *DriftImpl) ID() string {
	return "drift"
}

func (d *DriftImpl) Prepare(ctx context.Context, config hcl.Body) error {
	p := NewParser("")

	theCfg, diags := p.Decode(config, nil)
	if diags.HasErrors() {
		return diags
	}

	d.config = theCfg
	return nil
}

func (d *DriftImpl) Execute(ctx context.Context, req *model.ExecuteRequest) (ret *model.ExecutionResult) {
	ret = &model.ExecutionResult{}

	cb, _ := json.Marshal(d.config)
	d.logger.Debug("executing with config", "config", string(cb), "request", req.String())

	provs, err := req.Providers()
	if err != nil {
		ret.Error = err
		return
	}

	var iacProv *cqproto.GetProviderSchemaResponse
	for _, p := range provs {
		if p.Name == "terraform" { // TODO add more iac provider names
			if iacProv != nil {
				ret.Error = fmt.Errorf("only single IAC provider is supported at a time")
				return
			}
			iacProv = p
		}
	}
	if iacProv == nil {
		ret.Error = fmt.Errorf("no IAC provider detected, can't continue")
		return
	}

	conn, err := req.Conn()
	if err != nil {
		ret.Error = fmt.Errorf("no connection: %w", err)
		return
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
				ret.Error = diags
				return
			}
			if !ok {
				continue
			}

			found = true

			d.logger.Info("processing for provider", "provider", prov.Name, "config", cfg)

			for resName, res := range cfg.Resources {
				if res == nil {
					continue // skipped
				}
				pr := prov.ResourceTables[resName]
				if pr == nil {
					d.logger.Warn("skipping resource, not found in ResourceTables", "provider", prov.Name, "resource", resName)
					continue
				}

				iacData := res.IAC[iacProv.Name]
				if iacData == nil {
					d.logger.Debug("skipping resource, iac provider not configured", "provider", prov.Name, "resource", resName, "iac_provider", iacProv.Name)
					continue
				}

				d.logger.Info("will process for provider and resource", "provider", prov.Name, "resource", resName, "iac_provider", iacProv.Name)

				res.finalInterpret()

				d.logger.Info("do the table", "table", pr.Name, "ids", res.Identifiers, "ignore", res.IgnoreAttributes, "iac_name", iacData.Name, "iac_type", iacData.Type)
				// do the table iac_name=users iac_type=aws_iam_user ids=["account_id","id"] ignore=["password_last_used"] table=aws_iam_users

				// Drift per resource
				var dres *Result
				switch iacProv.Name {
				case "terraform":
					dres, err = d.driftTerraform(ctx, conn, prov.Name, pr, res, iacData)
				default:
					ret.Error = fmt.Errorf("no suitable handler found for %q", iacProv.Name)
					return
				}
				if err != nil {
					ret.Error = fmt.Errorf("drift failed for (%s,%s): %w", prov.Name, resName, err)
					return
				} else if dres != nil {
					dres.Provider = prov.Name
					dres.ResourceType = resName
					resList = append(resList, dres)
				}
			}

			break
		}

		if !found {
			ret.Error = fmt.Errorf("no suitable provider found for %q", cfg.Name)
			return
		}
	}

	ret.Result = resList.String()
	return ret
}

func (d *DriftImpl) driftTerraform(ctx context.Context, conn *pgxpool.Conn, cloudName string, cloudTable *schema.Table, resData *ResourceConfig, iacData *IACConfig) (*Result, error) {
	res := &Result{
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	// Get from IAC
	// SELECT i.instance_id, i.attributes from tf_resource_instances i JOIN tf_resources r ON r.cq_id=i.resource_id JOIN tf_data d ON d.cq_id=r.running_id WHERE d.backend_name='mylocal' AND r.provider='aws' AND r.mode='managed' AND r.type='aws_s3_bucket' AND r.name='s3_bucket';

	// Get from provider
	// SELECT account_id, region, name, arn FROM aws_s3_buckets;

	const (
		// TODO move into module params, with defaults
		tfBackendName = "mylocal"
		tfMode        = "managed"
	)

	// TODO check if cloud provider names always match with tf_resources.provider
	tfProvider := cloudName

	tfSelect := goqu.Dialect("postgres").From(goqu.T("tf_resource_instances").As("i")).
		Select("i.instance_id", "i.attributes").
		Join(goqu.T("tf_resources").As("r"), goqu.On(goqu.Ex{"r.cq_id": goqu.I("i.resource_id")})).
		Join(goqu.T("tf_data").As("d"), goqu.On(goqu.Ex{"d.cq_id": goqu.I("r.running_id")})).
		Where(goqu.Ex{"d.backend_name": goqu.V(tfBackendName)}).
		Where(goqu.Ex{"r.provider": goqu.V(tfProvider)}).
		Where(goqu.Ex{"r.mode": goqu.V(tfMode)}).
		Where(goqu.Ex{"r.type": goqu.V(iacData.Type)}).
		Where(goqu.Ex{"r.name": goqu.V(iacData.Name)})

	// Get equal resources
	{
		q, args, err := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).Join(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select("tf.instance_id").
			ToSQL()
		if err != nil {
			return nil, fmt.Errorf("build failed: %w", err)
		}
		d.logger.Info("generated equals query", "query", q, "args", args)

		res.Equal, err = d.queryIntoResourceList(ctx, conn, q, args)
		if err != nil {
			return nil, fmt.Errorf("query failed: %w", err)
		}
	}

	// Get missing resources
	{
		q, args, err := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).RightJoin(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select("tf.instance_id").
			ToSQL()
		if err != nil {
			return nil, fmt.Errorf("build failed: %w", err)
		}
		d.logger.Info("generated missing query", "query", q, "args", args)

		res.Missing, err = d.queryIntoResourceList(ctx, conn, q, args)
		if err != nil {
			return nil, fmt.Errorf("query failed: %w", err)
		}
	}

	// Get extra resources
	{
		q, args, err := goqu.Dialect("postgres").From(goqu.T(cloudTable.Name).As("c")).
			With("tf", tfSelect).LeftJoin(goqu.T("tf"), goqu.On(goqu.Ex{"tf.instance_id": goqu.I("c." + resData.Identifiers[0])})).
			Select(goqu.I("c." + resData.Identifiers[0])).Where(goqu.Ex{"tf.instance_id": nil}).
			ToSQL()
		if err != nil {
			return nil, fmt.Errorf("build failed: %w", err)
		}
		d.logger.Info("generated extras query", "query", q, "args", args)

		res.Extra, err = d.queryIntoResourceList(ctx, conn, q, args)
		if err != nil {
			return nil, fmt.Errorf("query failed: %w", err)
		}
	}

	// TODO get different resources

	// TODO

	return res, nil
}

func (d *DriftImpl) queryIntoResourceList(ctx context.Context, conn *pgxpool.Conn, query string, args []interface{}) ([]*Resource, error) {
	var list []string
	if err := pgxscan.Select(ctx, conn, &list, query, args...); err != nil {
		return nil, err
	}

	ret := make([]*Resource, len(list))
	for i := range list {
		ret[i] = &Resource{
			ID: list[i],
		}
	}

	return ret, nil
}

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
