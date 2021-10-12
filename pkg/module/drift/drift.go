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

func (d *DriftImpl) Execute(req *model.ExecuteRequest) *model.ExecutionResult {
	// TODO run
	cb, _ := json.Marshal(d.config)
	fmt.Printf("config is %s\n", string(cb))
	fmt.Printf("execute is %+v\n", req)
	return nil
}

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
