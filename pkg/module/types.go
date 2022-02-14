package module

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/hashicorp/hcl/v2"
)

type Module interface {
	// ID returns the name of the module
	ID() string
	// Supported module protocol versions, in order of preference
	ProtocolVersions() []uint32
	// Configure configures the module to run
	Configure(context.Context, Info, ModuleRunParams) error
	// Execute executes the module, using given args in ExecuteRequest
	Execute(context.Context, *ExecuteRequest) *ExecutionResult
	// ExampleConfig returns an example configuration to be put in config.hcl
	ExampleConfig() string
}

// Info about user supplied configs and provider supplied module info
type Info struct {
	// UserConfig is the config supplied by the user in config.hcl
	UserConfig hcl.Body

	// ProtocolVersion of the provider supplied module info
	ProtocolVersion uint32
	// ProviderData is the provider supplied module info, one map[string][]byte per provider
	ProviderData map[string]map[string][]byte
}

type ModuleRunParams interface{}

type ExecuteRequest struct {
	// Params are the invocation parameters specific to the module
	Params ModuleRunParams

	// Providers is the list of providers to process
	Providers []*cqproto.GetProviderSchemaResponse
	// Conn is the db connection to use
	Conn execution.QueryExecer
}

type ExecutionResult struct {
	Result   interface{} `json:"result"`
	Error    error       `json:"-"`
	ErrorMsg string      `json:"error,omitempty"`
}

type ExitCoder interface {
	ExitCode() int
}
