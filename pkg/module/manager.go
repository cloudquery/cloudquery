package module

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	modules  map[string]Module
	modOrder []string

	// Instance of database
	pool execution.QueryExecer

	// Logger instance
	logger hclog.Logger
}

// Manager is the interface that describes the interaction with the module manager.
// Implemented by ManagerImpl.
type Manager interface {
	// RegisterModule is used to register a module into the manager.
	RegisterModule(mod Module)

	// ExecuteModule executes the given module, validating the given module name and config first.
	ExecuteModule(ctx context.Context, modName string, profileConfig hcl.Body, execReq *ExecuteRequest) (*ExecutionResult, error)

	// ExampleConfigs returns a list of example module configs from loaded modules
	ExampleConfigs() []string
}

// NewManager returns a new manager instance.
func NewManager(pool execution.QueryExecer, logger hclog.Logger) *ManagerImpl {
	return &ManagerImpl{
		modules: make(map[string]Module),
		pool:    pool,
		logger:  logger,
	}
}

// RegisterModule is used to register a module into the manager.
func (m *ManagerImpl) RegisterModule(mod Module) {
	id := mod.ID()
	if _, ok := m.modules[id]; ok {
		panic("module " + id + " already registered")
	}
	m.modules[id] = mod
	m.modOrder = append(m.modOrder, id)
}

// ExecuteModule executes the given module, validating the given module name and config first.
func (m *ManagerImpl) ExecuteModule(ctx context.Context, modName string, cfg hcl.Body, execReq *ExecuteRequest) (*ExecutionResult, error) {
	mod, ok := m.modules[modName]
	if !ok {
		return nil, fmt.Errorf("module not found %q", modName)
	}

	if err := mod.Configure(ctx, cfg, execReq.Params); err != nil {
		return nil, fmt.Errorf("module configuration failed: %w", err)
	}

	execReq.Conn = m.pool

	return mod.Execute(ctx, execReq), nil
}

// ExampleConfigs returns a list of example module configs from loaded modules
func (m *ManagerImpl) ExampleConfigs() []string {
	ret := make([]string, 0, len(m.modules))
	for _, i := range m.modOrder {
		cfg := m.modules[i].ExampleConfig()
		if cfg == "" {
			continue
		}
		ret = append(ret, cfg)
	}
	return ret
}
