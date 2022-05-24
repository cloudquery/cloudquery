package config

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/xo/dburl"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/policy"
)

type Providers []*Provider

func (pp Providers) Names() []string {
	pNames := make([]string, len(pp))
	for i, p := range pp {
		pNames[i] = p.Name
	}
	return pNames
}

type Config struct {
	CloudQuery CloudQuery      `hcl:"cloudquery,block"`
	Providers  Providers       `hcl:"provider,block"`
	Policies   policy.Policies `hcl:"policy,block"`
	Modules    hcl.Body        `hcl:"modules,block"`
}

func (c Config) GetProvider(name string) (*Provider, error) {
	for _, p := range c.Providers {
		if name == p.Alias {
			return p, nil
		}
	}
	return nil, fmt.Errorf("provider %s does not exist", name)
}

type CloudQuery struct {
	PluginDirectory string            `hcl:"plugin_directory,optional"`
	PolicyDirectory string            `hcl:"policy_directory,optional"`
	Logger          *logging.Config   `hcl:"logging,block"`
	Providers       RequiredProviders `hcl:"provider,block"`
	Connection      *Connection       `hcl:"connection,block"`
	History         *struct {         // Deprecated
		Retention      int `hcl:"retention,optional"`
		TimeInterval   int `hcl:"interval,optional"`
		TimeTruncation int `hcl:"truncation,optional"`
	} `hcl:"history,block"`
	Policy *Policy `hcl:"policy,block"`
}

func (c CloudQuery) GetRequiredProvider(name string) (*RequiredProvider, error) {
	for _, p := range c.Providers {
		if name == p.Name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("provider %s does not exist", name)
}

type Connection struct {
	DSN string `hcl:"dsn,optional"`

	// These params are mutually exclusive with DSN
	Type     string   `hcl:"type,optional"`
	Username string   `hcl:"username,optional"`
	Password string   `hcl:"password,optional"`
	Host     string   `hcl:"host,optional"`
	Port     int      `hcl:"port,optional"`
	Database string   `hcl:"database,optional"`
	SSLMode  string   `hcl:"sslmode,optional"`
	Extras   []string `hcl:"extras,optional"`
}

func (c Connection) IsAnyConnParamsSet() bool {
	return c.Type != "" || c.Username != "" || c.Password != "" || c.Host != "" || c.Port != 0 || c.Database != "" || c.SSLMode != "" || len(c.Extras) > 0
}

func (c Connection) BuildFromConnParams() (*dburl.URL, error) {
	if c.Port == 0 {
		c.Port = 5432
	}
	if c.Type == "" {
		c.Type = "postgres"
	}
	if c.Host == "" {
		return nil, errors.New("missing host")
	}
	if c.Database == "" {
		return nil, errors.New("missing database")
	}

	u := url.URL{
		Scheme: c.Type,
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}
	if c.Username != "" && c.Password != "" {
		u.User = url.UserPassword(c.Username, c.Password)
	} else if c.Username != "" {
		u.User = url.User(c.Username)
	}

	v := url.Values{}
	if c.Extras != nil {
		for _, extra := range c.Extras {
			parts := strings.SplitN(extra, "=", 2)
			if len(parts) == 1 {
				v.Add(parts[0], "")
			} else {
				v.Add(parts[0], parts[1])
			}
		}
	}
	if c.SSLMode != "" {
		v.Set("sslmode", c.SSLMode)
	}
	u.RawQuery = v.Encode()

	return &dburl.URL{
		OriginalScheme: c.Type,
		URL:            u,
	}, nil
}

type RequiredProvider struct {
	Name    string  `hcl:"name,label"`
	Source  *string `hcl:"source,optional"`
	Version string  `hcl:"version"`
}

func (r RequiredProvider) String() string {
	var source string
	if r.Source != nil {
		source = *r.Source + "/"
	}
	return fmt.Sprintf("%scq-provider-%s@%s", source, r.Name, r.Version)
}

type RequiredProviders []*RequiredProvider

// Distinct returns one name per provider
func (r RequiredProviders) Distinct() RequiredProviders {
	ret := make(RequiredProviders, 0, len(r))
	dupes := make(map[string]struct{}, len(r))
	for _, p := range r {
		if _, ok := dupes[p.Name]; ok {
			continue
		}
		dupes[p.Name] = struct{}{}

		ret = append(ret, p)
	}
	return ret
}

func (r RequiredProviders) Names() []string {
	ret := make([]string, 0, len(r))
	dupes := make(map[string]struct{}, len(r))
	for _, p := range r {
		if _, ok := dupes[p.Name]; ok {
			continue
		}
		dupes[p.Name] = struct{}{}

		ret = append(ret, p.Name)
	}
	sort.Strings(ret)
	return ret
}

func (r RequiredProviders) Get(name string) *RequiredProvider {
	for _, p := range r {
		if p.Name == name {
			return p
		}
	}
	return nil
}

type Policy struct {
	DBPersistence bool `hcl:"db_persistence,optional"`
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
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
		{
			Type: "modules",
		},
	},
}
