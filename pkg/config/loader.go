package config

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/blobfs"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// EnvVarPrefix is a prefix for environment variable names to be exported for HCL substitution.
const EnvVarPrefix = "CQ_VAR_"

// Parser is the main interface to read configuration files and other related
// files from disk.
//
// It retains a cache of all files that are loaded so that they can be used
// to create source code snippets in diagnostics, etc.
type Parser struct {
	fs         afero.Afero
	p          *hclparse.Parser
	HCLContext hcl.EvalContext
}

type Option func(*Parser)

func WithFS(aferoFs afero.Fs) Option {
	return func(p *Parser) {
		p.fs = afero.Afero{Fs: aferoFs}
	}
}

// WithEnvironmentVariables fills hcl.Context with values of environment variables given in vars.
// Only variables that start with given prefix are considered. Prefix is removed from the name and
// the name is lower cased then.
func WithEnvironmentVariables(prefix string, vars []string) Option {
	return func(p *Parser) {
		EnvToHCLContext(&p.HCLContext, prefix, vars)
	}
}

// WithFileFunc adds the file() function to the parser.
func WithFileFunc(basePath string) Option {
	return func(p *Parser) {
		p.HCLContext.Functions["file"] = convert.MakeFileFunc(basePath)
	}
}

// NewParser creates and returns a new Parser.
func NewParser(options ...Option) *Parser {
	p := Parser{
		fs: afero.Afero{Fs: afero.OsFs{}},
		p:  hclparse.NewParser(),
		HCLContext: hcl.EvalContext{
			Variables: make(map[string]cty.Value),
			Functions: make(map[string]function.Function),
		},
	}

	for _, opt := range options {
		opt(&p)
	}
	return &p
}

// LoadFile is a low-level method that reads the file at the given path
func (p *Parser) LoadFile(path string) ([]byte, diag.Diagnostics) {
	var contents []byte
	// Example of path supported paths:
	// `./local/relative/path/to/config.hcl`
	// `/absolute/path/to/config.hcl`
	// `s3://object/in/remote/location/absolute/path/to/config.hcl`
	sanitizedPath, err := url.Parse(path)
	if err != nil {
		return nil, diag.FromError(err, diag.USER, diag.WithSummary("Failed to load config file: invalid path"), diag.WithDetails("The file %q could not be read", path))
	}

	if sanitizedPath.Scheme == "" {
		contents, err = p.fs.ReadFile(path)
	} else {
		contents, err = loadRemoteFile(path)
	}

	if err != nil {
		if e, ok := err.(*fs.PathError); ok {
			if errors.Is(err, fs.ErrNotExist) {
				err = fmt.Errorf("%s. Hint: Try `cloudquery init <provider>`", e.Err.Error())
			} else {
				err = fmt.Errorf(e.Err.Error())
			}
		}
		return nil, diag.FromError(err, diag.USER, diag.WithSummary("Failed to read file"), diag.WithDetails("The file %q could not be read", path))
	}
	if len(contents) == 0 {
		return nil, diag.FromError(err, diag.USER, diag.WithSummary("Failed to read file"), diag.WithDetails("The file %q is empty", path))
	}

	return contents, nil
}

func (p *Parser) LoadFromSource(name string, data []byte) (hcl.Body, diag.Diagnostics) {
	file, diags := p.p.ParseHCL(data, name)
	// If the returned file or body is nil, then we'll return a non-nil empty
	// body so we'll meet our contract that nil means an error reading the file.
	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), hclToSdkDiags(diags)
	}

	return file.Body, hclToSdkDiags(diags)
}

func EnvToHCLContext(evalContext *hcl.EvalContext, prefix string, vars []string) {
	for _, e := range vars {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], prefix) {
			evalContext.Variables[strings.TrimPrefix(pair[0], prefix)] = cty.StringVal(pair[1])
		}
	}
}

func hclToSdkDiags(hd hcl.Diagnostics) diag.Diagnostics {
	var sd diag.Diagnostics
	for _, dd := range hd {
		sv := diag.ERROR
		if dd.Severity == hcl.DiagWarning {
			sv = diag.WARNING
		}
		sd = sd.Add(diag.FromError(dd, diag.USER, diag.WithSeverity(sv), diag.WithSummary("%s", dd.Summary), diag.WithDetails("%s", dd.Detail)))
	}
	return sd
}

func loadRemoteFile(path string) ([]byte, error) {
	mux := fsimpl.NewMux()
	mux.Add(filefs.FS)
	mux.Add(blobfs.FS)

	sanitizedPath, _ := url.Parse(path)

	// go-fsimpl is looking for a "directory" where it is the full path without the actual file
	// but must have all of the query strings
	directory := path[:strings.LastIndex(path, "/")] + "?" + sanitizedPath.RawQuery
	fileName := sanitizedPath.Path[strings.LastIndex(sanitizedPath.Path, "/")+1:]

	fsys, err := mux.Lookup(directory)
	if err != nil {
		return nil, err
	}

	f, err := fsys.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	contents, err := io.ReadAll(f)
	return contents, err
}
