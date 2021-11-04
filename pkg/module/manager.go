package module

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/modules"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/afero"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	modules map[string]Module

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
	ExecuteModule(ctx context.Context, modName, modConfigPath string, execReq *ExecuteRequest) (*ExecutionResult, error)

	// ReadConfig reads the given module's default/builtin config
	ReadConfig(modName string) ([]byte, error)
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
	m.modules[mod.ID()] = mod
}

// ExecuteModule executes the given module, validating the given module name and config first.
func (m *ManagerImpl) ExecuteModule(ctx context.Context, modName, modConfigPath string, execReq *ExecuteRequest) (*ExecutionResult, error) {
	mod, ok := m.modules[modName]
	if !ok {
		return nil, fmt.Errorf("module not found %q", modName)
	}

	rawConfig, err := m.readDecodeConfig(mod.ID(), modConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config: %w", err)
	}

	if err := mod.Configure(ctx, rawConfig); err != nil {
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

// ReadConfig reads the given module's default/builtin config
func (m *ManagerImpl) ReadConfig(modName string) ([]byte, error) {
	// Try to find $modName.hcl
	filename := fmt.Sprintf("%s.hcl", modName)

	contents, err := modules.FS.ReadFile("configs/" + filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded config for module %s: %w", modName, err)
	}

	return contents, nil
}

// readDecodeConfig reads the module config. Uses configPath if set, if not, it will try to get the default module config
func (m *ManagerImpl) readDecodeConfig(modName string, configPath string) (hcl.Body, error) {
	if configPath != "" {
		osFs := file.NewOsFs()
		if info, err := osFs.Stat(configPath); err != nil || info.IsDir() {
			return nil, fmt.Errorf("could not find the given config %q", configPath)
		}
		parser := config.NewParser()
		configRaw, diags := parser.LoadHCLFile(configPath)
		if diags != nil && diags.HasErrors() {
			return nil, fmt.Errorf("failed to load config file: %#v", diags.Error())
		}
		inner, diags := parser.DecodeModuleConfig(configRaw, modName)
		if diags.HasErrors() {
			return nil, fmt.Errorf("DecodeModuleConfig: %w", diags)
		} else if inner == nil {
			return nil, fmt.Errorf("could not find valid module block in config")
		}
		return inner, nil
	}

	// Try to find $modName.hcl
	filename := fmt.Sprintf("%s.hcl", modName)

	contents, err := modules.FS.ReadFile("configs/" + filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded config for module %s: %w", modName, err)
	}

	fs := afero.NewMemMapFs()
	fp, err := fs.Create(filename)
	if err != nil {
		return nil, err
	}
	if _, err := fp.Write(contents); err != nil {
		_ = fp.Close()
		return nil, err
	}
	if err := fp.Close(); err != nil {
		return nil, err
	}

	parser := config.NewParser(config.WithFS(fs))
	configRaw, diags := parser.LoadHCLFile(filename)
	if diags != nil && diags.HasErrors() {
		return nil, fmt.Errorf("failed to load embedded config: %#v", diags.Error())
	}

	inner, diags := parser.DecodeModuleConfig(configRaw, modName)
	if diags.HasErrors() {
		return nil, fmt.Errorf("DecodeModuleConfig: %w", diags)
	} else if inner == nil {
		return nil, fmt.Errorf("could not find valid module block in embedded config")
	}
	return inner, nil
}
