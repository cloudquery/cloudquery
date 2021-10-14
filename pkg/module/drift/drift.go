package drift

import (
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
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

func (d *DriftImpl) Prepare(config hcl.Body) error {
	p := NewParser("")

	theCfg, diags := p.Decode(config, nil)
	if diags.HasErrors() {
		return diags
	}

	d.config = theCfg
	return nil
}

func (d *DriftImpl) Execute(req *model.ExecuteRequest) (ret *model.ExecutionResult) {
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

				// TODO do drift, per resource

			}

			break
		}

		if !found {
			ret.Error = fmt.Errorf("no suitable provider found for %q", cfg.Name)
			return
		}
	}

	return ret
}

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
