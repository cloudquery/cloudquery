package drift

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
)

type DriftImpl struct {
	logger hclog.Logger

	config hcl.Body
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
	// TODO parse config, if it isn't valid return error
	// TODO put config into d
	d.config = config
	return nil
}

func (d *DriftImpl) Execute(req *model.ExecuteRequest) *model.ExecutionResult {
	// TODO run
	fmt.Printf("config is %+v\n", d.config)
	fmt.Printf("execute is %+v\n", req)
	return nil
}

// Make sure we satisfy the interface
var _ model.Module = (*DriftImpl)(nil)
