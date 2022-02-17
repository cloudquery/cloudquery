package module

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/hashicorp/go-hclog"
)

type LowLevelQueryExecer interface {
	execution.Copier
	execution.QueryExecer
}

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
	ExecuteModule(ctx context.Context, execReq *ExecuteRequest) (*ExecutionResult, error)

	// ExampleConfigs returns a list of example module configs from loaded modules
	ExampleConfigs() []string
}

type moduleInfoRequester interface {
	GetProviderModule(ctx context.Context, providerName string, req cqproto.GetModuleRequest) (*cqproto.GetModuleResponse, error)
}

// NewManager returns a new manager instance.
func NewManager(pool LowLevelQueryExecer, logger hclog.Logger, r moduleInfoRequester) *ManagerImpl {
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
func (m *ManagerImpl) ExecuteModule(ctx context.Context, execReq *ExecuteRequest) (*ExecutionResult, error) {
	mod, ok := m.modules[execReq.Module]
	if !ok {
		return nil, fmt.Errorf("module not found %q", execReq.Module)
	}

	protoVersion, modInfo, err := m.collectProviderInfo(ctx, mod, execReq.Providers)
	if err != nil {
		return nil, fmt.Errorf("protocol negotiation failed: %w", err)
	}

	if err := mod.Configure(ctx, Info{
		UserConfig:      execReq.ProfileConfig,
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

	return 0, nil, versionError(mod.ID(), mod.ProtocolVersions(), allVersions)
}

func versionError(modName string, modVersions []uint32, provVersions map[string][]uint32) error {
	var (
		unsupportingProviders []string
		olderProviders        []string
		newerProviders        []string

		minRequired = minUint32(modVersions)
	)

	for p, versions := range provVersions {
		if len(versions) == 0 {
			unsupportingProviders = append(unsupportingProviders, p)
			continue
		}

		if maxSupplied := maxUint32(versions); minRequired > maxSupplied {
			olderProviders = append(olderProviders, p)
		} else if minRequired < maxSupplied {
			newerProviders = append(newerProviders, p)
		}
	}

	sort.Strings(unsupportingProviders)
	sort.Strings(olderProviders)
	sort.Strings(newerProviders)

	if l := len(unsupportingProviders); l == 1 {
		return fmt.Errorf("provider %s doesn't support %s yet", unsupportingProviders[0], modName)
	} else if l > 1 {
		return fmt.Errorf("providers %s don't support %s yet", strings.Join(unsupportingProviders, ", "), modName)
	}

	if l := len(olderProviders); l == 1 {
		return fmt.Errorf("provider %s seems to support an older version of %s, which is incompatible with your cloudquery version", olderProviders[0], modName)
	} else if l > 1 {
		return fmt.Errorf("providers %s seem to support an older version of %s, which is incompatible with your cloudquery version", strings.Join(olderProviders, ", "), modName)
	}

	if l := len(newerProviders); l == 1 {
		return fmt.Errorf("provider %s seems to support a newer version of %s, which is incompatible with your cloudquery version", newerProviders[0], modName)
	} else if l > 1 {
		return fmt.Errorf("providers %s seem to support a newer version of %s, which is incompatible with your cloudquery version", strings.Join(newerProviders, ", "), modName)
	}

	return fmt.Errorf("version mismatch between module and providers, please upgrade your providers and/or cloudquery")
}

func minUint32(v []uint32) uint32 {
	var smallest uint32
	for i := range v {
		if smallest == 0 || v[i] < smallest {
			smallest = v[i]
		}
	}
	return smallest
}

func maxUint32(v []uint32) uint32 {
	var biggest uint32
	for i := range v {
		if v[i] > biggest {
			biggest = v[i]
		}
	}
	return biggest
}
