package model

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
	// Module is the module that should be executed.
	Module Module
	// Additional args for the module
	Args []string

	// Providers is the callback to use to access to a list of providers to process
	Providers func() ([]*cqproto.GetProviderSchemaResponse, error)

	// Conn() is the callback to use to access a pg conn
	Conn func() (*pgxpool.Conn, error)
}

type ExecutionResult struct {
	Results []string
	Error   error
}

func (e *ExecuteRequest) String() string {
	if e.Module == nil {
		return fmt.Sprintf("[execute module <nil> with args %+v]", e.Args)
	}

	return fmt.Sprintf("[execute module %s with args %+v]", e.Module.ID(), e.Args)
}
