package module

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/cqproto"

	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Module interface {
	// ID returns the name of the module
	ID() string
	// Configure configures the module to run
	Configure(context.Context, hcl.Body) error
	// Execute executes the module, using given args in ExecuteRequest
	Execute(context.Context, *ExecuteRequest) *ExecutionResult
}

type ExecuteRequest struct {
	// Module is the module that should be executed
	Module Module
	// Params are the invocation parameters specific to the module
	Params interface{}

	// Providers is the list of providers to process
	Providers []*cqproto.GetProviderSchemaResponse
	// Conn is the db connection to use
	Conn *pgxpool.Conn
}

type ExecutionResult struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

func (e *ExecuteRequest) String() string {
	if e.Module == nil {
		return fmt.Sprintf("[execute module <nil> with params %+v]", e.Params)
	}

	return fmt.Sprintf("[execute module %s with params %+v]", e.Module.ID(), e.Params)
}
