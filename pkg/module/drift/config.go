package drift

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/cqproto"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
)

type BaseConfig struct {
	WildProvider *ProviderConfig
	Providers    []*ProviderConfig `hcl:"provider,block"`
}

type ProviderConfig struct {
	WildResource *ResourceConfig

	Name          string                     `hcl:"name,label"`
	Resources     map[string]*ResourceConfig `hcl:"resource,block"`
	Source        string                     `hcl:"source,optional"`
	Version       string                     `hcl:"version,optional"`
	SkipResources []string                   `hcl:"skip_resources,optional"`

	versionConstraints version.Constraints
}

type ResourceConfig struct {
	Identifiers       []string `hcl:"identifiers,optional"`
	IgnoreIdentifiers []string `hcl:"ignore_identifiers,optional"`
	Attributes        []string `hcl:"attributes,optional"`
	IgnoreAttributes  []string `hcl:"ignore_attributes,optional"`

	IAC map[string]*IACConfig

	defRange *hcl.Range
}

type IACConfig struct {
	Name string `hcl:"name,optional"`
	Type string `hcl:"type,optional"`

	AttributeMap []string `hcl:"attribute_map,optional"`

	attributeMap map[string]string // cloud vs. iac
	defRange     *hcl.Range
}

const wildcard = "*"

func (prov *ProviderConfig) applyWildProvider(wild *ProviderConfig) {
	if wild == nil {
		return
	}

	if len(prov.SkipResources) == 0 {
		prov.SkipResources = wild.SkipResources
	}
}

// applyWildResource sets the missing values from res if they are provided in the wild resource
// for IAC.attributeMap it adds on to the list of attributes
func (res *ResourceConfig) applyWildResource(wild *ResourceConfig) {
	if wild == nil {
		return
	}

	if len(res.Identifiers) == 0 {
		res.Identifiers = wild.Identifiers
	}
	if len(res.IgnoreIdentifiers) == 0 {
		res.IgnoreIdentifiers = wild.IgnoreIdentifiers
	}
	if len(res.Attributes) == 0 {
		res.Attributes = wild.Attributes
	}
	if len(res.IgnoreAttributes) == 0 {
		res.IgnoreAttributes = wild.IgnoreAttributes
	}
	if len(res.IAC) == 0 {
		res.IAC = wild.IAC
	} else {
		// add on attributeMap values from wild
		for k, v := range res.IAC {
			if wild.IAC[k] != nil {
				for kk, vv := range wild.IAC[k].attributeMap {
					if _, ok := v.attributeMap[kk]; !ok {
						res.IAC[k].attributeMap[kk] = vv
					}
				}
			}
		}
	}

}

// finalInterpret removes each element in IgnoredIdentifiers from Identifiers
// this has to be done after all apply/macro work has finished
func (res *ResourceConfig) finalInterpret() {
	// apply res.IgnoreIdentifiers: remove matching identifiers from res.Identifiers
	ignoreIdMap := make(map[string]struct{}, len(res.IgnoreIdentifiers))
	for i := range res.IgnoreIdentifiers {
		ignoreIdMap[res.IgnoreIdentifiers[i]] = struct{}{}
	}

	// remove item from slice
	idx := 0
	for i := range res.Identifiers {
		if _, ok := ignoreIdMap[res.Identifiers[i]]; ok {
			continue
		}
		res.Identifiers[idx] = res.Identifiers[i]
		idx++
	}
	res.Identifiers = res.Identifiers[:idx]
}

// applyProvider tries to apply the given config for the given provider, trying to match provider name and version constraints.
// Returns true if the given config is valid for the given provider and cfg is changed to resolve macros.
func (d *DriftImpl) applyProvider(cfg *ProviderConfig, p *cqproto.GetProviderSchemaResponse) (bool, hcl.Diagnostics) {
	if p.Name != cfg.Name {
		return false, nil // not the correct provider: names don't match
	}

	if len(cfg.versionConstraints) > 0 {
		pver, err := version.NewSemver(p.Version)
		if err != nil {
			return false, []*hcl.Diagnostic{
				{
					Severity: hcl.DiagError,
					Summary:  `Invalid provider version`,
					Detail:   fmt.Sprintf("could not parse provider version %q: %v", p.Version, err),
				},
			}
		}
		if !cfg.versionConstraints.Check(pver) {
			d.logger.Warn("provider is blocked by constraint", "provider", p.Name+"@"+p.Version, "constraint", cfg.Version)
			return false, nil // not the correct provider: versions don't match
		}
	}

	var diags hcl.Diagnostics

	skips := make(map[string]struct{}, len(cfg.SkipResources))
	for i := range cfg.SkipResources {
		skips[cfg.SkipResources[i]] = struct{}{}
	}

	for resName, res := range cfg.Resources {
		if _, ok := skips[resName]; ok {
			cfg.Resources[resName] = nil
			continue
		}

		tbl, ok := p.ResourceTables[resName]
		if !ok {
			d.logger.Warn("resource is not defined by the provider", "resource", resName)
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Specified resource not in provider`,
				Detail:   fmt.Sprintf("resource %q is not defined by the provider", resName),
				Subject:  res.defRange,
			})
			continue
		}

		for k, v := range map[placeholder][]string{
			placeholderResourceKey:             {resName},
			placeholderResourceName:            {tbl.Name},
			placeholderResourceColumnNames:     tbl.ColumnNames(),
			placeholderResourceOptsPrimaryKeys: tbl.PrimaryKeys(),
		} {
			res.Identifiers = replacePlaceholderInSlice(k, v, res.Identifiers)
			res.IgnoreIdentifiers = replacePlaceholderInSlice(k, v, res.IgnoreIdentifiers)
			res.Attributes = replacePlaceholderInSlice(k, v, res.Attributes)
			res.IgnoreAttributes = replacePlaceholderInSlice(k, v, res.IgnoreAttributes)
		}
	}

	return true, diags
}
