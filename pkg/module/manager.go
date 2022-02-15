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
	GetProviderModule(ctx context.Context, providerName string, req *cqproto.GetModuleRequest) (*cqproto.GetModuleResponse, error)
}

var errNegotiationFailed = fmt.Errorf("version mismatch between module and providers, please upgrade your provider and/or cloudquery")

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

	protoVersion, modInfo, err := m.collectProviderInfo(ctx, mod, execReq.Providers, 0)
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

func (m *ManagerImpl) collectProviderInfo(ctx context.Context, mod Module, provs []*cqproto.GetProviderSchemaResponse, forceVersion uint32) (uint32, map[string]ProviderData, error) {
	var doVersions []uint32
	if forceVersion > 0 {
		doVersions = []uint32{forceVersion}
	} else {
		doVersions = mod.ProtocolVersions()
	}

	var (
		ret               = make(map[string]ProviderData, len(provs)) // provider vs. info-key vs. files provided by provider under that key
		supportedVersions = make(map[string][]uint32, len(provs))
		foundUnsupported  bool
	)

	// Do initial requests
	for _, p := range provs {
		data, err := m.requester.GetProviderModule(ctx, p.Name, &cqproto.GetModuleRequest{
			Module:            mod.ID(),
			PreferredVersions: doVersions,
		})
		if err != nil {
			return 0, nil, fmt.Errorf("GetProviderModule %s: %w", p.Name, err)
		} else if data.Diagnostics.HasDiags() {
			return 0, nil, data.Diagnostics
		}
		supportedVersions[p.Name] = data.SupportedVersions

		if data.Version == 0 {
			foundUnsupported = true
			continue
		}

		ret[p.Name] = data.Info
		doVersions = []uint32{data.Version}
	}

	if !foundUnsupported {
		// happy path
		return doVersions[0], ret, nil
	}

	if forceVersion > 0 {
		return 0, nil, errNegotiationFailed
	}

	// negotiate: through each version supported by the module, in order of prefence, and try to find an entry which all providers can satisfy
	availableVersions := compileAvailableVersions(supportedVersions)
	for _, preferredVersion := range mod.ProtocolVersions() {
		list := availableVersions[preferredVersion]
		if len(list) < len(provs) {
			m.logger.Debug("skipping preferred module protocol version, available providers", "preferred_version", preferredVersion, "prov_versions", list)
			continue
		}

		// force that version
		m.logger.Info("negotiating module protocol version", "version", preferredVersion)
		return m.collectProviderInfo(ctx, mod, provs, preferredVersion)
	}
	return 0, nil, errNegotiationFailed
}

// compileAvailableVersions makes a matrix of all available versions by all providers
// gets a map of supported versions per provider, returns list of providers per version
func compileAvailableVersions(supportedVersions map[string][]uint32) map[uint32][]string {
	vers := make(map[uint32]map[string]struct{})
	for prov, versions := range supportedVersions {
		for _, v := range versions {
			mp, ok := vers[v]
			if !ok {
				mp = make(map[string]struct{})
			}
			mp[prov] = struct{}{}
			vers[v] = mp
		}
	}

	// convert from map[string]struct{} to []string
	ret := make(map[uint32][]string, len(vers))
	for k, v := range vers {
		list := make([]string, 0, len(v))
		for prov := range v {
			list = append(list, prov)
		}
		ret[k] = list
	}
	return ret
}
