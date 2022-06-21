package config

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/spf13/viper"
	"github.com/xo/dburl"
)

type Provider struct {
	Name                          string   `yaml:"name,omitempty" json:"name,omitempty" hcl:"name,label"`
	Alias                         string   `yaml:"alias,omitempty" json:"alias,omitempty" hcl:"alias,optional"`
	Resources                     []string `yaml:"resources,omitempty" json:"resources,omitempty" hcl:"resources,optional"`
	SkipResources                 []string `yaml:"skip_resources,omitempty" json:"skip_resources,omitempty" hcl:"skip_resources,optional"`
	Env                           []string `yaml:"env,omitempty" json:"env,omitempty" hcl:"env,optional"`
	Configuration                 []byte   `yaml:"-" json:"-"`
	MaxParallelResourceFetchLimit uint64   `yaml:"max_parallel_resource_fetch_limit,omitempty" json:"max_parallel_resource_fetch_limit,omitempty" hcl:"max_parallel_resource_fetch_limit"`
	MaxGoroutines                 uint64   `yaml:"max_goroutines,omitempty" json:"max_goroutines,omitempty" hcl:"max_goroutines"`
	ResourceTimeout               uint64   `yaml:"resource_timeout,omitempty" json:"resource_timeout,omitempty" hcl:"resource_timeout"`
}

type Providers []*Provider

type Config struct {
	CloudQuery CloudQuery `hcl:"cloudquery,block" yaml:"cloudquery" json:"cloudquery"`
	Providers  Providers  `hcl:"provider,block" yaml:"providers" json:"providers"`
}

type CloudQuery struct {
	Logger     *logging.Config   `yaml:"logging,omitempty" json:"logging,omitempty" hcl:"logging,block"`
	Providers  RequiredProviders `yaml:"providers,omitempty" json:"providers,omitempty" hcl:"provider,block"`
	Connection *Connection       `yaml:"connection,omitempty" json:"connection,omitempty" hcl:"connection,block"`
	Policy     *Policy           `yaml:"policy,omitempty" json:"policy,omitempty" hcl:"policy,block"`

	// Used internally
	PluginDirectory string `yaml:"-" json:"-"`
	PolicyDirectory string `yaml:"-" json:"-"`
}

type Connection struct {
	DSN string `yaml:"dsn,omitempty" json:"dsn,omitempty" hcl:"dsn,optional"`

	// These params are mutually exclusive with DSN
	Type     string   `yaml:"type,omitempty" json:"type,omitempty" hcl:"type,optional"`
	Username string   `yaml:"username,omitempty" json:"username,omitempty" hcl:"username,optional"`
	Password string   `yaml:"password,omitempty" json:"password,omitempty" hcl:"password,optional"`
	Host     string   `yaml:"host,omitempty" json:"host,omitempty" hcl:"host,optional"`
	Port     int      `yaml:"port,omitempty" json:"port,omitempty" hcl:"port,optional"`
	Database string   `yaml:"database,omitempty" json:"database,omitempty" hcl:"database,optional"`
	SSLMode  string   `yaml:"sslmode,omitempty" json:"sslmode,omitempty" hcl:"sslmode,optional"`
	Extras   []string `yaml:"extras,omitempty" json:"extras,omitempty" hcl:"extras,optional"`
}

type Policy struct {
	DBPersistence bool `yaml:"db_persistence,omitempty" json:"db_persistence,omitempty" hcl:"db_persistence,optional"`
}

type RequiredProvider struct {
	Name    string  `yaml:"name,omitempty" json:"name,omitempty" hcl:"name,label"`
	Source  *string `yaml:"source,omitempty" json:"source,omitempty" hcl:"source,optional"`
	Version string  `yaml:"version,omitempty" json:"version,omitempty" hcl:"version"`
}

type RequiredProviders []*RequiredProvider

func (pp Providers) Names() []string {
	pNames := make([]string, len(pp))
	for i, p := range pp {
		pNames[i] = p.Name
	}
	return pNames
}

func (c Config) GetProvider(name string) (*Provider, error) {
	for _, p := range c.Providers {
		if name == p.Alias {
			return p, nil
		}
	}
	return nil, fmt.Errorf("provider %s does not exist", name)
}

func (c CloudQuery) GetRequiredProvider(name string) (*RequiredProvider, error) {
	for _, p := range c.Providers {
		if name == p.Name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("provider %s does not exist", name)
}

func (c Connection) IsAnyConnParamsSet() bool {
	return c.Type != "" || c.Username != "" || c.Password != "" || c.Host != "" || c.Port != 0 || c.Database != "" || c.SSLMode != "" || len(c.Extras) > 0
}

func (c *Connection) BuildFromConnParams() error {
	if c.Port == 0 {
		c.Port = 5432
	}
	if c.Type == "" {
		c.Type = "postgres"
	}
	if c.Host == "" {
		return errors.New("missing host")
	}
	if c.Database == "" {
		return errors.New("missing database")
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

	c.DSN = (&dburl.URL{OriginalScheme: c.Type, URL: u}).String()

	return nil
}

func (r RequiredProvider) String() string {
	var source string
	if r.Source != nil {
		source = *r.Source + "/"
	}
	return fmt.Sprintf("%scq-provider-%s@%s", source, r.Name, r.Version)
}

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

func handleConnectionConfig(c *Connection) error {
	if ds := viper.GetString("dsn"); ds != "" {
		c.DSN = ds
		return nil
	}

	if c.DSN != "" {
		if c.IsAnyConnParamsSet() {
			return errors.New("DSN specified along with explicit attributes, only one type is supported")
		}
		return nil
	}

	return c.BuildFromConnParams()
}
