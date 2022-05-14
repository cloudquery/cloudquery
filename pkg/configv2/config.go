package configv2

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

type Provider struct {
	Name                          string      `yaml:"name,omitempty" json:"name,omitempty"`
	Alias                         string      `yaml:"alias,omitempty" json:"alias,omitempty"`
	Resources                     []string    `yaml:"resources,omitempty" json:"resources,omitempty"`
	Env                           []string    `yaml:"env,omitempty" json:"env,omitempty"`
	Configuration                 interface{} `yaml:"-"`
	MaxParallelResourceFetchLimit uint64      `yaml:"max_parallel_resource_fetch_limit,omitempty" json:"max_parallel_resource_fetch_limit,omitempty"`
	MaxGoroutines                 uint64      `yaml:"max_goroutines,omitempty" json:"max_goroutines,omitempty"`
	ResourceTimeout               uint64      `yaml:"resource_timeout,omitempty" json:"resource_timeout,omitempty"`
}

type Providers []Provider

type Connection struct {
	DSN string `yaml:"dsn,omitempty" json:"dsn,omitempty"`

	// These params are mutually exclusive with DSN
	Type     string   `yaml:"type,omitempty" json:"type,omitempty"`
	Username string   `yaml:"username,omitempty" json:"username,omitempty"`
	Password string   `yaml:"password,omitempty" json:"password,omitempty"`
	Host     string   `yaml:"host,omitempty" json:"host,omitempty"`
	Port     int      `yaml:"port,omitempty" json:"port,omitempty"`
	Database string   `yaml:"database,omitempty" json:"database,omitempty"`
	SSLMode  string   `yaml:"sslmode,omitempty" json:"sslmode,omitempty"`
	Extras   []string `yaml:"extras,omitempty" json:"extras,omitempty"`
}

type CloudQuery struct {
	PluginDirectory string            `yaml:"plugin_directory,omitempty" json:"plugin_directory,omitempty"`
	PolicyDirectory string            `yaml:"policy_directory,omitempty" json:"policy_directory,omitempty"`
	Logger          *logging.Config   `yaml:"logging,omitempty" json:"logging,omitempty"`
	Providers       RequiredProviders `yaml:"providers,omitempty" json:"providers,omitempty"`
	Connection      *Connection       `yaml:"connection,omitempty" json:"connection,omitempty"`
}

type Config struct {
	CloudQuery *CloudQuery     `yaml:"cloudquery,omitempty" json:"cloudquery,omitempty"`
	Providers  []Provider      `yaml:"providers,omitempty" json:"providers,omitempty"`
	Policies   policy.Policies `yaml:"policy" json:"policy"`
	Modules    hcl.Body        `yaml:"modules"`
}

func (pp Providers) Names() []string {
	pNames := make([]string, len(pp))
	for i, p := range pp {
		pNames[i] = p.Name
	}
	return pNames
}

func (c Config) GetProvider(name string) (Provider, error) {
	for _, p := range c.Providers {
		if name == p.Alias {
			return p, nil
		}
	}
	return Provider{}, fmt.Errorf("provider %s does not exist", name)
}

func (c CloudQuery) GetRequiredProvider(name string) (RequiredProvider, error) {
	for _, p := range c.Providers {
		if name == p.Name {
			return p, nil
		}
	}
	return RequiredProvider{}, fmt.Errorf("provider %s does not exist", name)
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
	Name    string  `yaml:"name"`
	Source  *string `yaml:"source"`
	Version string  `yaml:"version"`
}

func (r RequiredProvider) String() string {
	var source string
	if r.Source != nil {
		source = *r.Source + "/"
	}
	return fmt.Sprintf("%scq-provider-%s@%s", source, r.Name, r.Version)
}

type RequiredProviders []RequiredProvider

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

func (r RequiredProviders) Get(name string) RequiredProvider {
	for _, p := range r {
		if p.Name == name {
			return p
		}
	}
	return RequiredProvider{}
}
