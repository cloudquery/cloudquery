package config

import (
	"github.com/hashicorp/hcl/v2"
)

/*





 */
type Config struct {
	CloudQuery CloudQuery `hcl:"cloudquery,block"`
	Providers  []Provider `hcl:"provider,block"`
}

type CloudQuery struct {
	PluginDirectory string             `hcl:"plugin_directory,optional"`
	Providers       []RequiredProvider `hcl:"provider,block"`
}

type RequiredProvider struct {
	Name    string `hcl:"name,label"`
	Source  string `hcl:"source"`
	Version string `hcl:"version"`
}

type Provider struct {
	Name        string
	Resources   []string
	Meta hcl.Block
}
