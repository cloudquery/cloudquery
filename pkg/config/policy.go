package config

import "fmt"

type PolicyWrapper struct {
	Policies []*Policy `hcl:"policy,block"`
}

// Normalize normalizes all contained policies.
func (pw *PolicyWrapper) Normalize() {
	for _, p := range pw.Policies {
		p.Normalize()
	}
}

// Validate validates all contained policies.
func (pw *PolicyWrapper) Validate() error {
	for i, p := range pw.Policies {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("policy %s (%d): %w", p.Name, i, err)
		}
	}
	return nil
}

type Policy struct {
	Name        string         `hcl:"name,label"`
	Description string         `hcl:"description,optional"`
	Config      *Configuration `hcl:"configuration,block"`

	Policies []*Policy `hcl:"policy,block"`
	Queries  []*Query  `hcl:"query,block"`
	Views    []*View   `hcl:"view,block"`
}

// Normalize normalizes the policy by calling Normalize on contained policies and queries.
func (p *Policy) Normalize() {
	for _, p := range p.Policies {
		p.Normalize()
	}
	for _, q := range p.Queries {
		q.Normalize()
	}
}

// Validate validates the policy by calling Validate on contained policies and queries.
func (p *Policy) Validate() error {
	for i, p := range p.Policies {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("policy %s (%d): %w", p.Name, i, err)
		}
	}
	for i, q := range p.Queries {
		if err := q.Validate(); err != nil {
			return fmt.Errorf("query %s (%d): %w", q.Name, i, err)
		}
	}
	return nil
}

type View struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description,optional"`

	Query *Query `hcl:"query,block"`
}

type Configuration struct {
	Providers []*PolicyProvider `hcl:"provider,block"`
}

type PolicyProvider struct {
	Type    string `hcl:"type,label"`
	Version string `hcl:"version,optional"`
}

type QueryType string

const (
	ManualQuery    QueryType = "manual"
	AutomaticQuery QueryType = "automatic"
)

type Query struct {
	Name         string    `hcl:"name,label"`
	Description  string    `hcl:"description,optional"`
	ExpectOutput bool      `hcl:"expect_output,optional"`
	Type         QueryType `hcl:"type,optional"`
	Query        string    `hcl:"query"`
}

// Normalize normalizes the query by filling empty fields with default values were applicable etc.
func (q *Query) Normalize() {
	if q.Type == "" {
		q.Type = AutomaticQuery
	}
}

// Validate validates the query by checking that all fields have expected values.
func (q *Query) Validate() error {
	switch q.Type {
	case AutomaticQuery, ManualQuery:
		return nil
	default:
		return fmt.Errorf("invalid query type: %v", q.Type)
	}
}
