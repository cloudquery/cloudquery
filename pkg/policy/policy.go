package policy

import (
	"fmt"
	"strings"
)

type Policies []*Policy

func (pp Policies) All() []string {
	policyNames := make([]string, len(pp))
	for i, p := range pp {
		policyNames[i] = p.Name
	}
	return policyNames
}

func (pp Policies) Properties() map[string]interface{} {
	usedCustom := 0
	policies := make([]*Meta, 0)
	for _, p := range pp {
		if p.meta == nil || p.meta.Type != "hub" {
			usedCustom++
			// we don't add info about custom policies that were executed
			continue
		}
		policies = append(policies, p.meta)
	}
	return map[string]interface{}{
		"used_custom": usedCustom,
		"policies":    policies,
	}
}

type Policy struct {
	// Name of the policy
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// Short human-readable title about the policy
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	// Full documentation about the policy, this will be shown in the hub
	Doc    string         `yaml:"doc,omitempty" json:"doc,omitempty"`
	Config *Configuration `yaml:"configuration,omitempty" json:"configuration,omitempty"`

	Policies Policies `yaml:"policy,omitempty" json:"policy,omitempty"`
	Checks   []*Check `yaml:"check,omitempty" json:"check,omitempty"`
	Views    []*View  `yaml:"view,omitempty" json:"view,omitempty"`

	// Link to policy in filesystem/hub/git etc' to use, if source flag is set, all other attributes aren't allowed.
	Source string `yaml:"source,omitempty" json:"source,omitempty"`

	// List of identifiers that all checks and sub-policies must have, unless sub-policy overrides.
	Identifiers []string

	// Meta is information added to the policy after it was loaded
	meta *Meta
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
			return policy.Filter(nextPolicy)
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

type Meta struct {
	Type      string `json:"type,omitempty"`
	Version   string `json:"version,omitempty"`
	SubPolicy string `json:"sub_policy,omitempty"`
	Directory string `json:"directory,omitempty"`
}

type View struct {
	Name  string `yaml:"name,omitempty" json:"name,omitempty"`
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	Query string `yaml:"query" json:"query"`
}

type Configuration struct {
	Providers []*Provider `yaml:"providers,omitempty"`
}

type Provider struct {
	Type    string `yaml:"type,omitempty" json:"type,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

type QueryType string

const (
	ManualQuery    QueryType = "manual"
	AutomaticQuery QueryType = "automatic"
)

type Check struct {
	Name         string    `yaml:"name,omitempty" json:"name,omitempty"`
	Title        string    `yaml:"title,omitempty" json:"title,omitempty"`
	Doc          string    `yaml:"doc,omitempty" json:"doc,omitempty"`
	ExpectOutput bool      `yaml:"expect_output,omitempty" json:"expect_output,omitempty"`
	Type         QueryType `yaml:"type,omitempty" json:"type,omitempty"`
	Query        string    `yaml:"query,omitempty" json:"query,omitempty"`
	Reason       string    `yaml:"reason,omitempty" json:"reason,omitempty"`
}
