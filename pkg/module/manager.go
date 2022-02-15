package module

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
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

	// Instance of client to query module info
	requester moduleInfoRequester
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

type moduleInfoRequester interface {
	GetProviderModule(ctx context.Context, providerName string, req cqproto.GetModuleRequest) (*cqproto.GetModuleResponse, error)
}

// NewManager returns a new manager instance.
func NewManager(pool execution.QueryExecer, logger hclog.Logger, r moduleInfoRequester) *ManagerImpl {
	return &ManagerImpl{
		modules:   make(map[string]Module),
		pool:      pool,
		logger:    logger,
		requester: r,
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

	protoVersion, modInfo, err := m.collectProviderInfo(ctx, mod, execReq.Providers)
	if err != nil {
		return nil, fmt.Errorf("protocol negotiation failed: %w", err)
	}

	if err := mod.Configure(ctx, Info{
		UserConfig:      cfg,
		ProtocolVersion: protoVersion,
		ProviderData:    modInfo,
	}, execReq.Params); err != nil {
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

func (m *ManagerImpl) collectProviderInfo(ctx context.Context, mod Module, provs []*cqproto.GetProviderSchemaResponse) (uint32, map[string]cqproto.ModuleInfo, error) {
	var (
		providerVersionInfo = make(map[uint32]map[string]cqproto.ModuleInfo) // version vs provider vs info
		allVersions         = make(map[string][]uint32, len(provs))          // used for debug info
	)

	rq := cqproto.GetModuleRequest{
		Module:            mod.ID(),
		PreferredVersions: mod.ProtocolVersions(),
	}

	for _, p := range provs {
		data, err := m.requester.GetProviderModule(ctx, p.Name, rq)
		if err != nil {
			return 0, nil, fmt.Errorf("GetProviderModule %s: %w", p.Name, err)
		} else if data.Diagnostics.HasDiags() {
			return 0, nil, data.Diagnostics
		}
		allVersions[p.Name] = data.AvailableVersions

		for v := range data.Data {
			inf, ok := providerVersionInfo[v]
			if !ok {
				inf = make(map[string]cqproto.ModuleInfo)
			}
			inf[p.Name] = data.Data[v]
			providerVersionInfo[v] = inf
		}
	}

	// negotiate: through each version supported by the module, in order of preference, and try to find an entry which all providers can satisfy
	for _, preferredVersion := range mod.ProtocolVersions() {
		list := providerVersionInfo[preferredVersion]
		if len(list) < len(provs) {
			m.logger.Debug("skipping preferred module protocol version", "preferred_version", preferredVersion)
			continue
		}

		// use that version
		m.logger.Info("negotiating module protocol version", "version", preferredVersion)
		return preferredVersion, list, nil
	}

	return 0, nil, fmt.Errorf("version mismatch between module and providers, please upgrade your provider and/or cloudquery")
}
