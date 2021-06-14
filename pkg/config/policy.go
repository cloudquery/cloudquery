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

	Queries []*Query `hcl:"query,block"`
}

type Configuration struct {
	Providers []*PolicyProvider `hcl:"provider,block"`
}

type PolicyProvider struct {
	Type    string `hcl:"type,label"`
	Version string `hcl:"version,optional"`
}

type Query struct {
	Name         string `hcl:"name,label"`
	Description  string `hcl:"description,optional"`
	ExpectOutput bool   `hcl:"expect_output,optional"`
	Query        string `hcl:"query"`
}
