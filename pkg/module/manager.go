package module

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/core"

	"github.com/cloudquery/cloudquery/internal/logging"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/plugin"

	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	modules  map[string]Module
	modOrder []string
	storage  database.Storage
	pm       *plugin.Manager
}

// Manager is the interface that describes the interaction with the module manager.
// Implemented by ManagerImpl.
type Manager interface {
	// Register is used to register a module into the manager.
	Register(mod Module)
	// Execute executes the given module, validating the given module name and config first.
	Execute(ctx context.Context, execReq *ExecuteRequest) (*ExecutionResult, error)
	// ExampleConfigs returns a list of example module configs from loaded modules
	ExampleConfigs(providers []string) []string
}

// NewManager returns a new manager instance.
func NewManager(storage database.Storage, pm *plugin.Manager) *ManagerImpl {
	return &ManagerImpl{
		modules: make(map[string]Module),
		storage: storage,
		pm:      pm,
	}
}

// Register is used to register a module into the manager.
func (m *ManagerImpl) Register(mod Module) {
	id := mod.ID()
	if _, ok := m.modules[id]; ok {
		panic("module " + id + " already registered")
	}
	m.modules[id] = mod
	m.modOrder = append(m.modOrder, id)
}

// Execute executes the given module, validating the given module name and config first.
func (m *ManagerImpl) Execute(ctx context.Context, req *ExecutionOptions) (*ExecutionResult, error) {
	mod, ok := m.modules[req.Module]
	if !ok {
		return nil, fmt.Errorf("module not found %q", req.Module)
	}
	schemas := make([]*core.ProviderSchema, len(req.Providers))
	for i, p := range req.Providers {
		s, diags := core.GetProviderSchema(ctx, m.pm, &core.GetProviderSchemaOptions{Provider: p})
		if diags.HasErrors() {
			return nil, diags
		}
		schemas[i] = s
	}

	protoVersion, modInfo, err := m.collectProviderInfo(ctx, mod, schemas)
	if err != nil {
		return nil, fmt.Errorf("protocol negotiation failed: %w", err)
	}

	if err := mod.Configure(ctx, Info{
		UserConfig:      req.ProfileConfig,
		ProtocolVersion: protoVersion,
		ProviderData:    modInfo,
	}, req.Params); err != nil {
		return nil, fmt.Errorf("module configuration failed: %w", err)
	}
	// TODO: this is still bad
	conn, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "database-drift"), m.storage.DSN())
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to new database")
		return nil, err
	}
	defer conn.Close()

	// TODO: this is very weird behavior, execute should return an error
	result := mod.Execute(ctx, &ExecuteRequest{
		Module:        req.Module,
		ProfileConfig: req.ProfileConfig,
		Params:        req.Params,
		Schemas:       schemas,
		Conn:          conn,
	})
	log.Info().Str("module", req.Module).Msg("module execution finished")
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("module execution failed with error")
	}
	// TODO: print this nicely, not sure it makes sense in log
	log.Debug().Str("module", req.Module).Interface("data", result.Result).Msg("module execution results")
	return result, nil
}

// ExampleConfigs returns a list of example module configs from loaded modules
func (m *ManagerImpl) ExampleConfigs(providers []string) []string {
	ret := make([]string, 0, len(m.modules))
	for _, i := range m.modOrder {
		cfg := m.modules[i].ExampleConfig(providers)
		if cfg == "" {
			continue
		}
		ret = append(ret, cfg)
	}
	return ret
}

func (m *ManagerImpl) collectProviderInfo(ctx context.Context, mod Module, provs []*core.ProviderSchema) (uint32, map[string]cqproto.ModuleInfo, error) {
	var (
		providerVersionInfo = make(map[uint32]map[string]cqproto.ModuleInfo) // version vs provider vs info
		allVersions         = make(map[string][]uint32, len(provs))          // used for debug info
	)
	for _, p := range provs {
		data, err := GetProviderModule(ctx, m.pm, &GetModuleOptions{
			Provider: registry.Provider{
				Name:    p.Name,
				Version: p.Version,
				Source:  registry.DefaultOrganization, // TODO: won't work with community providers
			},
			Request: cqproto.GetModuleRequest{
				Module:            mod.ID(),
				PreferredVersions: mod.ProtocolVersions(),
			}})
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
			log.Debug().Uint32("preferred_version", preferredVersion).Msg("skipping preferred module protocol version")
			continue
		}
		// use that version
		log.Info().Uint32("version", preferredVersion).Msg("negotiating module protocol version")
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
