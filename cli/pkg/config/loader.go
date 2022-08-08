package config

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/blobfs"
	"github.com/hairyhenderson/go-fsimpl/filefs"
	"github.com/spf13/afero"
)

// EnvVarPrefix is a prefix for environment variable names to be exported for HCL substitution.
const EnvVarPrefix = "CQ_VAR_"

// Parser is the main interface to read configuration files and other related
// files from disk.
//
// It retains a cache of all files that are loaded so that they can be used
// to create source code snippets in diagnostics, etc.
type Parser struct {
	fs afero.Afero

	variables map[string]string
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
		for _, v := range vars {
			pair := strings.SplitN(v, "=", 2)
			if strings.HasPrefix(pair[0], prefix) {
				var varVal string
				if len(pair) == 2 {
					varVal = pair[1]
				}
				p.variables[strings.TrimPrefix(pair[0], prefix)] = varVal
			}
		}
	}
}

// NewParser creates and returns a new Parser.
func NewParser(options ...Option) *Parser {
	p := Parser{
		fs:        afero.Afero{Fs: afero.OsFs{}},
		variables: make(map[string]string),
	}

	for _, opt := range options {
		opt(&p)
	}
	return &p
}

func (p *Parser) getVariableValue(s string) string {
	return p.variables[s]
}

// LoadFile is a low-level method that reads the file at the given path
func (p *Parser) LoadFile(path string) ([]byte, diag.Diagnostics) {
	var contents []byte
	// Example of path supported paths:
	// `./local/relative/path/to/config.yml`
	// `/absolute/path/to/config.yml`
	// `s3://object/in/remote/location/absolute/path/to/config.yml`
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
