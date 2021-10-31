package drift

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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
	Deep              *bool    `hcl:"deep,optional"`         // Check attributes if true, otherwise just match identifiers
	ParentMatch       string   `hcl:"parent_match,optional"` // Only valid for relational resources, name of column to match to parent's cq_id
	Filters           []string `hcl:"filters,optional"`      // SQL filters to exclude cloud providers default resources

	IAC map[string]*IACConfig

	defRange *hcl.Range
}

type IACConfig struct {
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
// for ignoreIdentifiers, ignoreAttributes and IAC.attributeMap it adds on to the list of values, always growing the list
func (res *ResourceConfig) applyWildResource(wild *ResourceConfig) {
	if wild == nil {
		return
	}

	if len(res.Identifiers) == 0 {
		res.Identifiers = wild.Identifiers
	}
	if len(res.Attributes) == 0 {
		res.Attributes = wild.Attributes
	}
	if res.Deep == nil {
		res.Deep = wild.Deep
	}
	if res.ParentMatch == "" {
		res.ParentMatch = wild.ParentMatch
	}

	// add on ignoreIdentifiers, ignoreAttributes and filters values from wild
	res.IgnoreIdentifiers = mergeDedupSlices(res.IgnoreIdentifiers, wild.IgnoreIdentifiers)
	res.IgnoreAttributes = mergeDedupSlices(res.IgnoreAttributes, wild.IgnoreAttributes)
	res.Filters = mergeDedupSlices(res.Filters, wild.Filters)

	if len(res.IAC) == 0 {
		res.IAC = wild.IAC
		return
	}

	// add on attributeMap values from wild
	for k, v := range res.IAC {
		if wild.IAC[k] == nil {
			continue
		}
		for kk, vv := range wild.IAC[k].attributeMap {
			if _, ok := v.attributeMap[kk]; !ok {
				res.IAC[k].attributeMap[kk] = vv
			}
		}
	}
}

func mergeDedupSlices(a ...[]string) []string {
	dupes := make(map[string]struct{})
	for i := range a {
		for v := range a[i] {
			dupes[a[i][v]] = struct{}{}
		}
	}
	ret := make([]string, 0, len(dupes))
	for k := range dupes {
		ret = append(ret, k)
	}
	return ret
}

// finalInterpret removes each element in IgnoredIdentifiers from Identifiers
// this has to be done after all apply/macro work has finished
func (res *ResourceConfig) finalInterpret() {
	// apply res.IgnoreIdentifiers: remove matching identifiers from res.Identifiers
	res.Identifiers = removeIgnored(res.Identifiers, res.IgnoreIdentifiers)
	// apply res.IgnoreAttributes: remove matching identifiers from res.Attributes
	res.Attributes = removeIgnored(res.Attributes, res.IgnoreAttributes)
}

func removeIgnored(list []string, ignored []string) []string {
	ignoredMap := make(map[string]struct{}, len(ignored))
	for i := range ignored {
		ignoredMap[ignored[i]] = struct{}{}
	}

	// remove item from slice
	idx := 0
	for i := range list {
		if _, ok := ignoredMap[list[i]]; ok {
			continue
		}
		list[idx] = list[i]
		idx++
	}
	return list[:idx]
}

// applyProvider tries to apply the given config for the given provider, trying to match provider name and version constraints.
// Returns true if the given config is valid for the given provider and cfg is changed to resolve macros.
func (d *Drift) applyProvider(cfg *ProviderConfig, p *cqproto.GetProviderSchemaResponse) (bool, hcl.Diagnostics) {
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

		tbl := d.lookupResource(resName, p)
		if tbl == nil {
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

		// {$sql:*} identifiers are still not replaced
	}

	return true, diags
}

func (d *Drift) lookupResource(resName string, prov *cqproto.GetProviderSchemaResponse) *provResource {
	if d.tableMap == nil {
		var setTableMap func(res *schema.Table, parent *provResource)
		setTableMap = func(res *schema.Table, parent *provResource) {
			d.tableMap[res.Name] = &provResource{
				Table:  res,
				Parent: parent,
			}
			for _, rel := range res.Relations {
				setTableMap(rel, d.tableMap[res.Name])
			}
		}

		d.tableMap = make(map[string]*provResource)
		for resId, res := range prov.ResourceTables {
			d.tableMap[resId] = &provResource{Table: res}
			setTableMap(res, nil)
		}
	}

	tbl, ok := d.tableMap[resName]
	if !ok {
		d.logger.Warn("Skipping resource, not found in ResourceTables", "resource", resName)
		return nil
	}

	return tbl
}
