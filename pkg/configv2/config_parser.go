package configv2

import (
	"fmt"
	"net/url"

	_ "embed"

	"github.com/hashicorp/hcl/v2"
	"github.com/spf13/afero"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
)

//go:embed schema.json
var schema []byte

// ReadConfigFile reads cloudquery config files from either local or remote locations
// and reutrn bytes and diags
// Example of path supported paths:
// `./local/relative/path/to/config.hcl`
// `/absolute/path/to/config.hcl`
// `s3://object/in/remote/location/absolute/path/to/config.hcl`
func ReadConfigFile(path string) ([]byte, error) {
	// Example of path supported paths:
	// `./local/relative/path/to/config.hcl`
	// `/absolute/path/to/config.hcl`
	// `s3://object/in/remote/location/absolute/path/to/config.hcl`
	sanitizedPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	fs := afero.Afero{Fs: afero.OsFs{}}
	var content []byte
	if sanitizedPath.Scheme == "" {
		content, err = fs.ReadFile(path)
	} else {
		content, err = loadRemoteFile(path)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return content, nil
}

func UnmarshalConfig(data []byte) (*Config, *gojsonschema.Result, error) {
	c := Config{}
	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(c)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate config: %w", err)
	}
	return &c, result, err
}

// ReadModuleConfigProfiles separates the module config from the modules block, where block identifier is the module name.
func ReadModuleConfigProfiles(module string, block hcl.Body) (map[string]hcl.Body, error) {
	if block == nil {
		return nil, nil
	}

	content, _, diags := block.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       module,
				LabelNames: []string{"name"},
			},
		},
	})
	if diags.HasErrors() {
		return nil, diags
	}

	ret := make(map[string]hcl.Body, len(content.Blocks))
	for i := range content.Blocks {
		if _, ok := ret[content.Blocks[i].Labels[0]]; ok {
			return nil, hcl.Diagnostics{
				{
					Severity: hcl.DiagError,
					Summary:  "Duplicate profile name",
					Detail:   fmt.Sprintf("Profile name %q already defined", content.Blocks[i].Labels[0]),
					Subject:  content.Blocks[i].DefRange.Ptr(),
				},
			}
		}

		ret[content.Blocks[i].Labels[0]] = content.Blocks[i].Body
	}
	return ret, nil
}

// func decodeCloudQueryBlock(block *hcl.Block, ctx *hcl.EvalContext) (CloudQuery, hcl.Diagnostics) {
// 	var cq CloudQuery
// 	// Pre-populate with existing values
// 	cq.Logger = &logging.GlobalConfig
// 	var diags hcl.Diagnostics
// 	diags = diags.Extend(gohcl.DecodeBody(block.Body, ctx, &cq))

// 	// TODO: decode in a more generic way
// 	if cq.Connection == nil {
// 		cq.Connection = &Connection{}
// 	}

// 	if err := handleConnectionBlock(cq.Connection); err != nil {
// 		diags = append(diags, &hcl.Diagnostic{
// 			Severity: hcl.DiagError,
// 			Summary:  "Invalid DSN configuration",
// 			Detail:   err.Error(),
// 			Subject:  &block.DefRange,
// 		})
// 	}

// 	datadir := viper.GetString("data-dir")

// 	if dir := viper.GetString("plugin-dir"); dir != "" {
// 		if dir == "." {
// 			if dir, err := os.Getwd(); err == nil {
// 				cq.PluginDirectory = dir
// 			}
// 		} else {
// 			cq.PluginDirectory = dir
// 		}
// 	} else if datadir != "" {
// 		cq.PluginDirectory = filepath.Join(datadir, "providers")
// 	}

// 	if dir := viper.GetString("policy-dir"); dir != "" {
// 		if dir == "." {
// 			if dir, err := os.Getwd(); err == nil {
// 				cq.PolicyDirectory = dir
// 			}
// 		} else {
// 			cq.PolicyDirectory = dir
// 		}
// 	} else if datadir != "" {
// 		cq.PolicyDirectory = filepath.Join(datadir, "policies")
// 	}

// 	// validate provider versions
// 	for _, cp := range cq.Providers {
// 		if cp.Version != "latest" && !strings.HasPrefix(cp.Version, "v") {
// 			diags = append(diags, &hcl.Diagnostic{
// 				Severity: hcl.DiagError,
// 				Summary:  fmt.Sprintf("Provider %s version %s is invalid", cp.Name, cp.Version),
// 				Detail:   "Please set to 'latest' version or valid semantic versioning starting with vX.Y.Z",
// 				Subject:  &block.DefRange,
// 			})
// 		}
// 	}
// 	return cq, diags
// }

// func handleConnectionBlock(c *Connection) error {
// 	if ds := viper.GetString("dsn"); ds != "" {
// 		c.DSN = ds
// 		return nil
// 	}
// 	if c.DSN != "" {
// 		if c.IsAnyConnParamsSet() {
// 			return errors.New("DSN specified along with explicit attributes, only one type is supported")
// 		}
// 		return nil
// 	}

// 	s, err := c.BuildFromConnParams()
// 	if err != nil {
// 		return err
// 	}
// 	a := s.String()
// 	c.DSN = a
// 	return nil
// }
