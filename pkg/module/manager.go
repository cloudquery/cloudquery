package module

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	modules  map[string]Module
	modOrder []string

	// Instance of a database connection pool
	pool *pgxpool.Pool

	// Logger instance
	logger hclog.Logger
}

// Manager is the interface that describes the interaction with the module manager.
// Implemented by ManagerImpl.
type Manager interface {
	// RegisterModule is used to register a module into the manager.
	RegisterModule(mod Module)

	// ExecuteModule executes the given module, validating the given module name and config first.
	ExecuteModule(ctx context.Context, modName string, configBlock hcl.Body, execReq *ExecuteRequest) (*ExecutionResult, error)

	// ExampleConfigs returns a list of example module configs from loaded modules
	ExampleConfigs() []string
}

// NewManager returns a new manager instance.
func NewManager(pool *pgxpool.Pool, logger hclog.Logger) *ManagerImpl {
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
func (m *ManagerImpl) ExecuteModule(ctx context.Context, modName string, cfgBlock hcl.Body, execReq *ExecuteRequest) (*ExecutionResult, error) {
	mod, ok := m.modules[modName]
	if !ok {
		return nil, fmt.Errorf("module not found %q", modName)
	}

	profileConfig, err := m.readModuleConfigProfiles(mod.ID(), cfgBlock)
	if err != nil {
		return nil, err
	}

	if err := mod.Configure(ctx, profileConfig, execReq.Params); err != nil {
		return nil, fmt.Errorf("module configuration failed: %w", err)
	}

	// Acquire connection from the connection pool
	execReq.Conn, err = m.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from the connection pool: %w", err)
	}
	defer execReq.Conn.Release()

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

// readModuleConfigProfiles separates the module config from the modules block
func (m *ManagerImpl) readModuleConfigProfiles(modName string, block hcl.Body) (map[string]hcl.Body, error) {
	if block == nil {
		return nil, nil
	}

	cfgs, diags := config.DecodeModuleProfile(block, modName)
	if diags.HasErrors() {
		return nil, fmt.Errorf("DecodeModuleProfile: %w", diags)
	}
	return cfgs, nil
}
