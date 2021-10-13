package model

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/hashicorp/hcl/v2"
)

type Module interface {
	ID() string
	Prepare(config hcl.Body) error
	Execute(*ExecuteRequest) *ExecutionResult
}

type ExecuteRequest struct {
	// Module is the module that should be executed.
	Module Module

	Args []string

	// Providers is the list of providers to process
	Providers []*config.Provider
}

type ExecutionResult struct {
	Error error
}

func (e *ExecuteRequest) String() string {
	if e.Module == nil {
		return fmt.Sprintf("[run module <nil> with args %+v]", e.Args)
	}

	return fmt.Sprintf("[run module %s with args %+v]", e.Module.ID(), e.Args)
}
