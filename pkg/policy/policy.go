package policy

import "fmt"

type Policies []*Policy

func (pp Policies) All() []string {
	policyNames := make([]string, len(pp))
	for i, p := range pp {
		policyNames[i] = p.Name
	}
	return policyNames
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
	return p.meta.subPolicy
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

type Meta struct {
	Type      string
	Version   string
	subPolicy string
	Directory string
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

const (
	ManualQuery    QueryType = "manual"
	AutomaticQuery QueryType = "automatic"
)

type Check struct {
	Name         string    `hcl:"name,label"`
	Title        string    `hcl:"title,optional"`
	Doc          string    `hcl:"doc,optional"`
	ExpectOutput bool      `hcl:"expect_output,optional"`
	Type         QueryType `hcl:"type,optional"`
	Query        string    `hcl:"query"`
}
