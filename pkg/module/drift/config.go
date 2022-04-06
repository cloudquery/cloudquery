package drift

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
)

//go:embed drift.hcl
var builtinConfig []byte

type BaseConfig struct {
	WildProvider *ProviderConfig
	Providers    []*ProviderConfig      `hcl:"provider,block"`
	Terraform    *TerraformSourceConfig `hcl:"terraform,block"`
}

type ProviderConfig struct {
	WildResource *ResourceConfig

	Name            string                     `hcl:"name,label"`
	Resources       map[string]*ResourceConfig `hcl:"resource,block"`
	Version         string                     `hcl:"version,optional"`
	IgnoreResources ResourceSelectors          `hcl:"ignore_resources,optional"`
	CheckResources  ResourceSelectors          `hcl:"check_resources,optional"`
	AccountIDs      []string                   `hcl:"account_ids,optional"`

	versionConstraints version.Constraints
}

type ResourceConfig struct {
	Identifiers       []string `hcl:"identifiers,optional"`
	IgnoreIdentifiers []string `hcl:"ignore_identifiers,optional"`
	Attributes        []string `hcl:"attributes,optional"`
	IgnoreAttributes  []string `hcl:"ignore_attributes,optional"`
	Deep              *bool    `hcl:"deep,optional"`    // Check attributes if true, otherwise just match identifiers
	Filters           []string `hcl:"filters,optional"` // SQL filters to exclude cloud providers default resources
	Sets              []string `hcl:"sets,optional"`    // Unordered list-attributes where item order doesn't matter

	IAC map[iacProvider]*IACConfig

	defRange *hcl.Range

	// check/ignore selectors
	acl ResourceACL
}

// ResourceACL manages resource allow and ignore lists
type ResourceACL struct {
	// AllowEnabled determines whether allow list is enabled (Allow contains entries) or not
	AllowEnabled bool
	// Allow contains the allow list for resources
	Allow ResourceSelectors
	// Ignore contains the ignore list for resources
	Ignore ResourceSelectors
}

// ShouldSkip gets a resource and compares it to the ACL, returning whether the given resource should be skipped or not
func (r ResourceACL) ShouldSkip(resource *Resource) bool {
	if r.AllowEnabled && !r.Allow.ContainsInstance(resource.ID) && !r.Allow.ContainsInstance("*") && !r.Allow.ContainsTags(resource.Tags) {
		return true
	}
	if r.Ignore.ContainsInstance(resource.ID) || r.Ignore.ContainsInstance("*") || r.Ignore.ContainsTags(resource.Tags) {
		return true
	}
	return false
}

// HasTagFilters returns true if the ACL contains tag filters
func (r ResourceACL) HasTagFilters() bool {
	if r.AllowEnabled {
		for _, f := range r.Allow {
			if f.Tags != nil {
				return true
			}
		}
	}

	for _, f := range r.Ignore {
		if f.Tags != nil {
			return true
		}
	}

	return false
}

type IACConfig struct {
	Type        string   `hcl:"type,optional"`
	Path        string   `hcl:"path,optional"`
	Identifiers []string `hcl:"identifiers,optional"`

	AttributeMap []string `hcl:"attribute_map,optional"`

	attributeMap map[string]string // cloud vs. iac
	defRange     *hcl.Range
}

type TerraformSourceConfig struct {
	Backend TerraformBackend `hcl:"backend"`

	// S3 backend
	Bucket  string   `hcl:"bucket,optional"`
	Keys    []string `hcl:"keys,optional"`
	Region  string   `hcl:"region,optional"`
	RoleARN string   `hcl:"role_arn,optional"`

	// Local backend
	Files []string `hcl:"files,optional"`
}

type TerraformBackend string

const (
	TFLocal TerraformBackend = "local"
	TFS3    TerraformBackend = "s3"
)

func (t TerraformBackend) Valid() bool {
	return t == TFLocal || t == TFS3
}

func (c TerraformSourceConfig) Validate() error {
	switch c.Backend {
	case TFLocal:
		if len(c.Files) == 0 {
			return fmt.Errorf("files not specified")
		}
		return nil
	case TFS3:
		if c.Bucket == "" {
			return fmt.Errorf("bucket not specified")
		}
		if len(c.Keys) == 0 {
			return fmt.Errorf("keys not specified")
		}
		return nil
	default:
		return fmt.Errorf("invalid backend type")
	}
}

type ResourceSelectors []*ResourceSelector

type ResourceSelector struct {
	Type string
	ID   *string
	Tags *map[string]string
}

func (t ResourceSelectors) ByType(resourceType string) ResourceSelectors {
	ret := make([]*ResourceSelector, 0, len(t))
	for _, s := range t {
		if s.Type == resourceType {
			ret = append(ret, s)
		}
	}
	return ret
}

func (t ResourceSelectors) AllInstances() bool {
	return t.ContainsInstance("*")
}

func (t ResourceSelectors) ContainsInstance(id string) bool {
	for _, s := range t {
		if s.ID != nil && *s.ID == id {
			return true
		}
	}
	return false
}

func (t ResourceSelectors) HasTags() bool {
	for _, s := range t {
		if s.Tags != nil {
			return true
		}
	}
	return false
}

func (t ResourceSelectors) ContainsTags(tags map[string]string) bool {
	if tags == nil {
		return false
	}
	for _, s := range t {
		if s.Tags == nil {
			continue
		}

		matches := 0
		for k, v := range *s.Tags {
			if v2, ok := tags[k]; ok && v2 == v {
				matches++
			}
		}

		if matches == len(*s.Tags) {
			return true
		}
	}

	return false
}

func parseResourceSelectors(input []string) (ResourceSelectors, error) {
	ret := make([]*ResourceSelector, 0, len(input))
	for _, s := range input {
		parts := strings.SplitN(s, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid resource selector, should be in type:id or type:[tags] format")
		}
		if parts[0] == "" {
			return nil, fmt.Errorf("type can't be empty, use * for wildcard")
		}
		r := &ResourceSelector{
			Type: parts[0],
		}
		if len(parts[1]) > 2 && strings.HasPrefix(parts[1], "[") && strings.HasSuffix(parts[1], "]") {
			var slc []string
			// Parse tags in key=value format separated by comma
			for _, tag := range strings.Split(strings.Trim(parts[1], "[]"), ",") {
				if tag == "" {
					return nil, fmt.Errorf("invalid empty tag in resource selector")
				}
				if !strings.Contains(tag, "=") {
					return nil, fmt.Errorf("invalid tag in resource selector: %q", tag)
				}
				slc = append(slc, tag)
			}
			slcTags := parseTags(slc)
			r.Tags = &slcTags
		} else {
			r.ID = &parts[1]
		}

		ret = append(ret, r)
	}
	return ret, nil
}

func parseTags(tags []string) map[string]string {
	ret := make(map[string]string, len(tags))
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		tagParts := strings.SplitN(tag, "=", 2)
		if len(tagParts) == 1 {
			ret[tagParts[0]] = ""
		} else {
			ret[tagParts[0]] = tagParts[1]
		}
	}
	return ret
}

const wildcard = "*"

func (b *BaseConfig) FindProvider(name string) *ProviderConfig {
	for i := range b.Providers {
		if b.Providers[i].Name == name {
			return b.Providers[i]
		}
	}
	return nil
}

func (prov *ProviderConfig) applyWildProvider(wild *ProviderConfig) {
	if wild == nil {
		return
	}

	if len(prov.IgnoreResources) == 0 {
		prov.IgnoreResources = wild.IgnoreResources
	}
	if len(prov.CheckResources) == 0 {
		prov.CheckResources = wild.CheckResources
	}
	if len(prov.AccountIDs) == 0 {
		prov.AccountIDs = wild.AccountIDs
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

	// add on ignoreIdentifiers, ignoreAttributes and filters values from wild
	res.IgnoreIdentifiers = mergeDedupSlices(res.IgnoreIdentifiers, wild.IgnoreIdentifiers)
	res.IgnoreAttributes = mergeDedupSlices(res.IgnoreAttributes, wild.IgnoreAttributes)
	res.Filters = mergeDedupSlices(res.Filters, wild.Filters)
	res.Sets = mergeDedupSlices(res.Sets, wild.Sets)

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

func (prov *ProviderConfig) resourceKeys() []string {
	k := make([]string, 0, len(prov.Resources))
	for i := range prov.Resources {
		k = append(k, i)
	}
	sort.Strings(k)
	return k
}

func (prov *ProviderConfig) interpolatedResourceMap(iacProvider iacProvider, logger hclog.Logger) map[string]*ResourceConfig {
	resourceKeys := prov.resourceKeys()
	ret := make(map[string]*ResourceConfig, len(resourceKeys))

	for _, resName := range resourceKeys {
		res := prov.Resources[resName]
		if res == nil {
			continue // skipped
		}
		iacData := res.IAC[iacProvider]
		if iacData == nil {
			logger.Debug("Will skip resource, iac provider not configured", "provider", prov.Name, "resource", resName, "iac_provider", iacProvider)
			continue
		}

		// Remove each element in IgnoredIdentifiers from Identifiers
		// This has to be done after all apply/macro work has finished

		// apply res.IgnoreIdentifiers: remove matching identifiers from res.Identifiers
		res.Identifiers = removeIgnored(res.Identifiers, res.IgnoreIdentifiers)
		// apply res.IgnoreAttributes: remove matching identifiers from res.Attributes
		res.Attributes = removeIgnored(res.Attributes, res.IgnoreAttributes)

		ret[resName] = res
	}

	return ret
}

func (d *Drift) findProvider(cfg *ProviderConfig, schemas []*cqproto.GetProviderSchemaResponse) (*cqproto.GetProviderSchemaResponse, error) {
	for _, schema := range schemas {
		if ok, diags := d.applyProvider(cfg, schema); diags.HasErrors() {
			return nil, diags
		} else if ok {
			return schema, nil
		}
	}

	return nil, fmt.Errorf("no suitable provider found for %q", cfg.Name)
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
	sort.Strings(ret)
	return ret
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
// Returns true if the given config is valid for the given provider and cfg is changed to resolve macros and acl processing
func (d *Drift) applyProvider(cfg *ProviderConfig, p *cqproto.GetProviderSchemaResponse) (bool, hcl.Diagnostics) {
	if p.Name != cfg.Name {
		return false, nil // not the correct provider: names don't match
	}

	if len(cfg.versionConstraints) > 0 {
		pver, err := version.NewSemver(p.Version)
		if err == nil {
			if pr := pver.Prerelease(); pr != "" && strings.HasPrefix(pr, "SNAPSHOT") {
				// re-parse without prerelease info
				v := strings.SplitN(p.Version, "-", 2)
				pver, err = version.NewVersion(v[0])
			}
		}
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

	allIgs := cfg.IgnoreResources.ByType("*")
	allChecks := cfg.CheckResources.ByType("*")
	checkEnabled := len(cfg.CheckResources) > 0

	for resName, res := range cfg.Resources {
		cleanRes, _ := SplitHashedResource(resName)

		// CheckResources / IgnoreResources broad strokes...
		if checkEnabled {
			res.acl.AllowEnabled = true
			res.acl.Allow = append(cfg.CheckResources.ByType(cleanRes), allChecks...)
			if !res.acl.Allow.AllInstances() && !res.acl.Allow.HasTags() {
				delete(cfg.Resources, resName)
				continue
			}
		}

		res.acl.Ignore = append(cfg.IgnoreResources.ByType(cleanRes), allIgs...)
		if res.acl.Ignore.AllInstances() {
			delete(cfg.Resources, resName)
			continue
		}

		tbl := d.lookupResource(resName, p)
		if tbl == nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Specified resource not in provider`,
				Detail:   fmt.Sprintf("resource %q is not defined by the provider", cleanRes),
				Subject:  res.defRange,
			})
			continue
		}

		for k, v := range map[placeholder][]string{
			placeholderResourceKey:             {cleanRes},
			placeholderResourceName:            {tbl.Name},
			placeholderResourceColumnNames:     tbl.NonCQColumns(),
			placeholderResourceOptsPrimaryKeys: tbl.NonCQPrimaryKeys(),
		} {
			res.Identifiers = replacePlaceholderInSlice(k, v, res.Identifiers)
			res.IgnoreIdentifiers = replacePlaceholderInSlice(k, v, res.IgnoreIdentifiers)
			res.Attributes = replacePlaceholderInSlice(k, v, res.Attributes)
			res.IgnoreAttributes = replacePlaceholderInSlice(k, v, res.IgnoreAttributes)
			res.Sets = replacePlaceholderInSlice(k, v, res.Sets)
		}

		// {$sql:*} identifiers are still not replaced
	}

	return true, diags
}

func (d *Drift) lookupResource(resName string, prov *cqproto.GetProviderSchemaResponse) *traversedTable {
	if d.tableMap == nil {
		d.tableMap = make(map[string]resourceMap)
	}

	if d.tableMap[prov.Name] == nil {
		d.tableMap[prov.Name] = traverseResourceTable(prov.ResourceTables)
	}

	res, _ := SplitHashedResource(resName)
	return d.tableMap[prov.Name][res]
}

// SplitHashedResource splits a given resource name and returns the resource and hash elements separately.
func SplitHashedResource(configResName string) (string, string) {
	resParts := strings.SplitN(configResName, "#", 2)
	if len(resParts) == 1 {
		return resParts[0], ""
	}
	return resParts[0], resParts[1]
}
