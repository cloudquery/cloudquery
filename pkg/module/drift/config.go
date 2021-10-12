package drift

type BaseConfig struct {
	Providers []*ProviderConfig `hcl:"provider,block"`
}

type ProviderConfig struct {
	Name          string                     `hcl:"name,label"`
	Resources     map[string]*ResourceConfig `hcl:"resource,block"`
	Source        string                     `hcl:"source,optional"`
	Version       string                     `hcl:"version,optional"`
	SkipResources []string                   `hcl:"skip_resources,optional"`
}

type ResourceConfig struct {
	Identifiers      []string `hcl:"identifiers,optional"`
	IgnoreAttributes []string `hcl:"ignore_attributes,optional"`
	TfType           string   `hcl:"tf_type,optional"`
	TfName           string   `hcl:"tf_name,optional"`
}

const wildcard = "*"
