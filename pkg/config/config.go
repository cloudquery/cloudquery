package config

import (
	"errors"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/hashicorp/hcl/v2"
)

type Config struct {
	CloudQuery CloudQuery  `hcl:"cloudquery,block"`
	Providers  []*Provider `hcl:"provider,block"`
}

func (c Config) GetProvider(name string) (*Provider, error) {
	for _, p := range c.Providers {
		if name == p.Name {
			return p, nil
		}
	}
	return nil, errors.New("provider does not exist")
}

type CloudQuery struct {
	PluginDirectory string              `hcl:"plugin_directory,optional"`
	Logger          *logging.Config     `hcl:"logging,block"`
	Providers       []*RequiredProvider `hcl:"provider,block"`
	Connection      Connection          `hcl:"connection,block"`
}

type Connection struct {
	DSN string `hcl:"dsn,attr"`
}

type RequiredProvider struct {
	Name    string `hcl:"name,label"`
	Source  string `hcl:"source"`
	Version string `hcl:"version"`
}

// configFileSchema is the schema for the top-level of a config file. We use
// the low-level HCL API for this level so we can easily deal with each
// block type separately with its own decoding logic.
var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "cloudquery",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
	},
}

// cloudqueryBlockSchema is the schema for a top-level "cloudquery" block in
// a configuration file.
var cloudqueryBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "plugin_directory"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "connection",
		},
		{
			Type: "required_providers",
		},
	},
}
