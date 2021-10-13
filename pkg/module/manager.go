package module

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/modules"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module/model"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/afero"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	modules map[string]model.Module

	// Instance of a database connection pool
	pool *pgxpool.Pool

	// Logger instance
	logger hclog.Logger
}

// Manager is the interface that describes the interaction with the module manager.
// Implemented by ManagerImpl.
type Manager interface {
	// RegisterModule is used to register a module into the manager.
	RegisterModule(mod model.Module)

	// ParseModuleReference parses and validates the given arguments into an execution request.
	ParseModuleReference(baseReq model.ExecuteRequest, args []string, modConfigPath string) (*model.ExecuteRequest, error)

	// RunModule runs the given module.
	RunModule(ctx context.Context, execRequest *model.ExecuteRequest) (*model.ExecutionResult, error)
}

// NewManager returns a new manager instance.
func NewManager(pool *pgxpool.Pool, logger hclog.Logger) *ManagerImpl {
	return &ManagerImpl{
		modules: make(map[string]model.Module),
		pool:    pool,
		logger:  logger,
	}
}

// RegisterModule is used to register a module into the manager.
func (m *ManagerImpl) RegisterModule(mod model.Module) {
	m.modules[mod.ID()] = mod
}

// ParseModuleReference parses and validates the given arguments into an execution request.
func (m *ManagerImpl) ParseModuleReference(baseReq model.ExecuteRequest, args []string, modConfigPath string) (*model.ExecuteRequest, error) {
	// Make sure the mandatory args are given
	if len(args) < 1 {
		return nil, fmt.Errorf("invalid module name. Module name is required but got %#v", args)
	}

	mod, ok := m.modules[args[0]]
	if !ok {
		return nil, fmt.Errorf("module not found %q", args[0])
	}

	rawConfig, err := m.readConfig(mod.ID(), modConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config: %w", err)
	}

	if err := mod.Prepare(rawConfig); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	baseReq.Module = mod
	baseReq.Args = args

	return &baseReq, nil
}

// RunModule runs the given module.
func (m *ManagerImpl) RunModule(ctx context.Context, execReq *model.ExecuteRequest) (*model.ExecutionResult, error) {

	var (
		oneDb   = &sync.Once{}
		theConn *pgxpool.Conn
	)

	execReq.Conn = func() (*pgxpool.Conn, error) {
		var err error
		oneDb.Do(func() {
			// Acquire connection from the connection pool
			theConn, err = m.pool.Acquire(ctx)
		})
		if err != nil {
			return nil, fmt.Errorf("failed to acquire connection from the connection pool: %w", err)
		}
		return theConn, nil
	}
	defer func() {
		if theConn != nil {
			theConn.Release()
		}
	}()

	res := execReq.Module.Execute(execReq)
	return res, nil
}

// readConfig reads the module config. Uses configPath if set, if not, it will try to get the default module config
func (m *ManagerImpl) readConfig(modName string, configPath string) (hcl.Body, error) {
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
