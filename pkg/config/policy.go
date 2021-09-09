package config

type PolicyWrapper struct {
	Policies []*Policy `hcl:"policy,block"`
}

type Policy struct {
	Name        string         `hcl:"name,label"`
	Description string         `hcl:"description,optional"`
	Config      *Configuration `hcl:"configuration,block"`

	Policies []*Policy `hcl:"policy,block"`
	Queries  []*Query  `hcl:"query,block"`
	Views    []*View   `hcl:"view,block"`
}

type View struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,optional"`

	Query *Query `hcl:"query,block"`
}

type Configuration struct {
	Providers []*PolicyProvider `hcl:"provider,block"`
}

type PolicyProvider struct {
	Type    string `hcl:"type,label"`
	Version string `hcl:"version,optional"`
}

type QueryType string

const (
	ManualQuery    QueryType = "manual"
	AutomaticQuery QueryType = "automatic"
)

type Query struct {
	Name         string    `hcl:"name,label"`
	Description  string    `hcl:"description,optional"`
	ExpectOutput bool      `hcl:"expect_output,optional"`
	Type         QueryType `hcl:"type,optional"`
	Query        string    `hcl:"query"`
}
