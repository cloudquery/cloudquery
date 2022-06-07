package policy

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/internal/getter"
)

type Policies []*Policy

type Meta struct {
	Type      string `json:"type,omitempty"`
	Version   string `json:"version,omitempty"`
	SubPolicy string `json:"sub_policy,omitempty"`
	Directory string `json:"directory,omitempty"`
}

type View struct {
	Name  string `hcl:"name,label"`
	Title string `hcl:"title,optional"`
	Query string `hcl:"query"`
}

type Configuration struct {
	Providers []*Provider `hcl:"provider,block"`
}

type Provider struct {
	Type    string `hcl:"type,label"`
	Version string `hcl:"version,optional"`
}

type QueryType string

type Check struct {
	Name         string    `hcl:"name,label"`
	Title        string    `hcl:"title,optional"`
	Doc          string    `hcl:"doc,optional"`
	ExpectOutput bool      `hcl:"expect_output,optional"`
	Type         QueryType `hcl:"type,optional"`
	Query        string    `hcl:"query"`
	Reason       string    `hcl:"reason,optional"`
}

type Analytic struct {
	// Whether policy will persist in database
	Persistence bool
	// Whether policy was defined in configuration
	UsedConfig bool
	// Name of the policy
	Name string
	// Type of the policy i.e S3/Hub/Git
	Type string
	// The selector used for the policy
	Selector string
	// Whether policy is private
	Private bool
}

type Policy struct {
	// Name of the policy
	Name string `hcl:"name,label"`
	// Short human-readable title about the policy
	Title string `hcl:"title,optional"`
	// Full documentation about the policy, this will be shown in the hub
	Doc    string         `hcl:"doc,optional"`
	Config *Configuration `hcl:"configuration,block"`

	Policies Policies `hcl:"policy,block"`
	Checks   []*Check `hcl:"check,block"`
	Views    []*View  `hcl:"view,block"`

	// Link to policy in filesystem/hub/git etc' to use, if source flag is set, all other attributes aren't allowed.
	Source string `hcl:"source,optional"`

	// List of identifiers that all checks and sub-policies must have, unless sub-policy overrides.
	Identifiers []string

	// Meta is information added to the policy after it was loaded
	meta *Meta
}

const (
	ManualQuery    QueryType = "manual"
	AutomaticQuery QueryType = "automatic"
)

func (pp Policies) All() []string {
	policyNames := make([]string, len(pp))
	for i, p := range pp {
		policyNames[i] = p.Name
	}
	return policyNames
}

func (pp Policies) Get(policyName, subPath string) Policies {
	if subPath == "" {
		policyName, subPath = getter.ParseSourceSubPolicy(policyName)
	}

	for _, p := range pp {
		if policyName != p.Name {
			continue
		}
		if subPath != "" && p.SubPolicy() == "" {
			return Policies{
				&Policy{
					Name:   p.Name,
					Source: p.Source + "//" + subPath,
				},
			}
		}
		return Policies{p}
	}
	return nil
}

func (p Policy) Analytic(dbPersistence, usedConfig bool) Analytic {
	pa := Analytic{
		Persistence: dbPersistence,
		UsedConfig:  usedConfig,
		Name:        p.Name,
		Type:        p.SourceType(),
		Selector:    p.SubPolicy(),
		Private:     p.SourceType() != "hub",
	}
	if !p.HasMeta() {
		policyName, subPath := getter.ParseSourceSubPolicy(p.Source)
		dp, _, _ := DetectPolicy(policyName, subPath)
		pa.Type = dp.SourceType()
		pa.Selector = subPath
		pa.Name = policyName
		pa.Private = p.SourceType() != "hub"
	}
	if pa.Private {
		pa.Selector = "private"
		pa.Name = p.Sha256Hash()
	}
	return pa
}

func (p Policy) String() string {
	if p.SubPolicy() != "" {
		return fmt.Sprintf("%s//%s", p.Name, p.SubPolicy())
	}
	return p.Name
}

func (p Policy) Version() string {
	if p.meta == nil {
		return "v0.0.0"
	}
	return p.meta.Version
}

func (p Policy) SubPolicy() string {
	if p.meta == nil {
		return ""
	}
	return p.meta.SubPolicy
}

func (p Policy) SourceType() string {
	if p.meta == nil {
		return ""
	}
	return p.meta.Type
}

func (p Policy) HasMeta() bool {
	return p.meta != nil
}

func (p Policy) HasChecks() bool {
	for _, policy := range p.Policies {
		if policy.HasChecks() {
			return true
		}
	}
	return len(p.Checks) > 0
}

func (p Policy) TotalQueries() int {
	count := 0
	if len(p.Policies) > 0 {
		for _, inner := range p.Policies {
			count += inner.TotalQueries()
		}
	}
	return count + len(p.Checks)
}

// Path should not include the root of the policy
// If no policy or control matches selector then a shell policy is returned

func (p Policy) Filter(path string) Policy {
	if path == "" {
		return p
	}
	selectorPath := strings.SplitN(path, "/", 2)
	if len(selectorPath) == 0 {
		return p
	}
	var emptyPolicy Policy
	nextPolicy := ""
	if strings.Contains(path, "/") {
		nextPolicy = selectorPath[1]
	}
	for _, policy := range p.Policies {
		if policy.Name == selectorPath[0] {
			filtered := policy.Filter(nextPolicy)
			if filtered.Config == nil {
				filtered.Config = p.Config
			}
			return filtered
		}
	}

	for _, check := range p.Checks {
		if check.Name == selectorPath[0] {
			p.Checks = make([]*Check, 0)
			p.Checks = append(p.Checks, check)
			return p
		}
	}

	return emptyPolicy
}

func (p Policy) Sha256Hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", p.Policies)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (a Analytic) Properties() map[string]interface{} {
	return map[string]interface{}{
		"policy_persistence":    a.Persistence,
		"policy_name":           a.Name,
		"policy_type":           a.Type,
		"policy_is_private":     a.Private,
		"policy_selector":       a.Selector,
		"policy_used_in_config": a.UsedConfig,
	}
}
