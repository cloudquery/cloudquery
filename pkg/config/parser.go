package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
)

type SourceType string

const (
	SourceJSON = "json"
	SourceHCL  = "hcl"
)

// Parser is the main interface to read configuration files and other related
// files from disk.
//
// It retains a cache of all files that are loaded so that they can be used
// to create source code snippets in diagnostics, etc.
type Parser struct {
	fs afero.Afero
	p  *hclparse.Parser
}

// NewParser creates and returns a new Parser that reads files from the given
// filesystem. If a nil filesystem is passed then the system's "real" filesystem
// will be used, via afero.OsFs.
func NewParser(fs afero.Fs) *Parser {
	if fs == nil {
		fs = afero.OsFs{}
	}

	return &Parser{
		fs: afero.Afero{Fs: fs},
		p:  hclparse.NewParser(),
	}
}

// LoadHCLFile is a low-level method that reads the file at the given path,
// parses it, and returns the hcl.Body representing its root. In many cases
// it is better to use one of the other Load*File methods on this type,
// which additionally decode the root body in some way and return a higher-level
// construct.
//
// If the file cannot be read at all -- e.g. because it does not exist -- then
// this method will return a nil body and error diagnostics. In this case
// callers may wish to ignore the provided error diagnostics and produce
// a more context-sensitive error instead.
//
// The file will be parsed using the HCL native syntax unless the filename
// ends with ".json", in which case the HCL JSON syntax will be used.
func (p *Parser) LoadHCLFile(path string) (hcl.Body, hcl.Diagnostics) {
	src, err := p.fs.ReadFile(path)

	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	return p.loadFromSource(path, src, SourceType(filepath.Ext(path)))
}

func (p *Parser) loadFromSource(name string, data []byte, ext SourceType) (hcl.Body, hcl.Diagnostics) {
	var file *hcl.File
	var diags hcl.Diagnostics
	switch ext {
	case SourceJSON:
		file, diags = p.p.ParseJSON(data, name)
	default:
		file, diags = p.p.ParseHCL(data, name)
	}
	// If the returned file or body is nil, then we'll return a non-nil empty
	// body so we'll meet our contract that nil means an error reading the file.
	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), diags
	}

	return file.Body, diags
}

func (p *Parser) LoadConfigFromSource(name string, data []byte) (*Config, hcl.Diagnostics) {
	body, diags := p.loadFromSource(name, data, SourceHCL)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfig(body, diags)
}

func (p *Parser) LoadConfigFile(path string) (*Config, hcl.Diagnostics) {

	body, diags := p.LoadHCLFile(path)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfig(body, diags)
}

func (p *Parser) decodeConfig(body hcl.Body, diags hcl.Diagnostics) (*Config, hcl.Diagnostics) {

	config := &Config{}

	content, contentDiags := body.Content(configFileSchema)
	diags = append(diags, contentDiags...)

	for _, block := range content.Blocks {
		switch block.Type {
		case "cloudquery":
			contentDiags = gohcl.DecodeBody(block.Body, nil, &config.CloudQuery)
			diags = append(diags, contentDiags...)
			// TODO: decode in a more generic way
			if dsn := viper.GetString("dsn"); dsn != "" {
				config.CloudQuery.Connection.DSN = dsn
			}
			if dir := viper.GetString("plugin-dir"); dir != "" {
				config.CloudQuery.PluginDirectory = dir
			}
		case "provider":
			cfg, cfgDiags := decodeProviderBlock(block)
			diags = append(diags, cfgDiags...)
			if cfg != nil {
				config.Providers = append(config.Providers, cfg)
			}
		default:
			// Should never happen because the above cases should be exhaustive
			// for all block type names in our schema.
			continue
		}
	}
	return config, diags
}
