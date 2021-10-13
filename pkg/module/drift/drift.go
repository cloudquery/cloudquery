package drift

import (
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/module/model"
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
	d.logger.Debug("executing with config", "config", string(cb), "request", req)

	provs, err := req.Providers()
	if err != nil {
		ret.Error = err
		return
	}

	for _, cfg := range d.config.Providers {
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

				d.logger.Info("will process for provider and resource", "provider", prov.Name, "resource", resName)

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
