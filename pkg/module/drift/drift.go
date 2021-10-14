package drift

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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

				d.logger.Info("do the table", "table", pr.Name, "ids", res.Identifiers, "ignore", res.IgnoreAttributes, "iac_name", iacData.Name, "iac_type", iacData.Type)
				// do the table iac_name=users iac_type=aws_iam_user ids=["account_id","id"] ignore=["password_last_used"] table=aws_iam_users

				// Drift per resource
				var dres *Result
				switch iacProv.Name {
				case "terraform":
					dres, err = d.driftTerraform(ctx, conn, pr, res, iacData)
				default:
					ret.Error = fmt.Errorf("no suitable handler found for %q", iacProv.Name)
					return
				}
				if err != nil {
					ret.Error = fmt.Errorf("drift failed for (%s,%s): %w", prov.Name, resName, err)
					return
				} else if dres != nil {
					dres.Provider = pr.Name
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

func (d *DriftImpl) driftTerraform(ctx context.Context, conn *pgxpool.Conn, cloudProv *schema.Table, resData *ResourceConfig, iacData *IACConfig) (*Result, error) {
	res := &Result{
		Different: nil,
		Equal:     nil,
		Missing:   nil,
		Extra:     nil,
	}

	// TODO compare in SQL

	// Get from IAC
	// SELECT i.instance_id, i.attributes from tf_resource_instances i JOIN tf_resources r ON r.cq_id=i.resource_id JOIN tf_data d ON d.cq_id=r.running_id WHERE d.backend_name='mylocal' AND r.provider='aws' AND r.mode='managed' AND r.type='aws_s3_bucket' AND r.name='s3_bucket';

	// Get from provider
	// SELECT account_id, region, name, arn FROM aws_s3_buckets;

	// TODO

	return res, nil
}

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
