package module

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
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
	// ProviderData is the provider supplied module info: Provider vs. data from the provider
	ProviderData map[string]ProviderData
}

// ProviderData is info-key vs. list of files under that key
type ProviderData map[string][]*cqproto.ModuleFile

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

func (p ProviderData) GetHCL(key string) (hcl.Body, hcl.Diagnostics) {
	files := p[key]
	if len(files) == 0 {
		return nil, nil
	}

	pa := hclparse.NewParser()
	hf := make([]*hcl.File, len(files))
	var diags hcl.Diagnostics
	for i := range files {
		f, d := pa.ParseHCL(files[i].Contents, files[i].Name)
		hf[i] = f
		diags = append(diags, d...)
	}
	if diags.HasErrors() {
		return nil, diags
	}

	return hcl.MergeFiles(hf), nil
}
