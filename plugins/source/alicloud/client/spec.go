package client

import (
	_ "embed"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Spec struct {
	Accounts          []AccountSpec `json:"accounts,omitempty"`
	BillHistoryMonths int           `json:"bill_history_months,omitempty"`
	Concurrency       int           `json:"concurrency,omitempty"`
}

type AccountSpec struct {
	Name      string   `json:"name,omitempty"`
	Regions   []string `json:"regions,omitempty"`
	AccessKey string   `json:"access_key,omitempty"`
	SecretKey string   `json:"secret_key,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 50000
	}
	if s.BillHistoryMonths == 0 {
		s.BillHistoryMonths = 12
	}
}

func (s *Spec) Validate() error {
	if len(s.Accounts) == 0 {
		return fmt.Errorf("missing alicloud accounts in configuration")
	}
	accountNames := map[string]struct{}{}
	for _, account := range s.Accounts {
		if account.Name == "" {
			return fmt.Errorf("missing alicloud account name in configuration")
		}
		if _, found := accountNames[account.Name]; found {
			return fmt.Errorf("duplicate alicloud account name %s in configuration", account.Name)
		}
		accountNames[account.Name] = struct{}{}
		if account.AccessKey == "" {
			return fmt.Errorf("missing access_key in account configuration for account %s", account.Name)
		}
		if account.SecretKey == "" {
			return fmt.Errorf("missing secret_key in account configuration for account %s", account.Name)
		}
		if len(account.Regions) == 0 {
			return fmt.Errorf("missing regions in account configuration for account %s", account.Name)
		}
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	one := uint64(1)

	accounts := *sc.Properties.Value("accounts").OneOf[0]
	accounts.MinItems = &one
	sc.Properties.Set("accounts", &accounts)

	sc.Required = append(sc.Required, "accounts")
}

func (AccountSpec) JSONSchemaExtend(sc *jsonschema.Schema) {
	one := uint64(1)

	name := *sc.Properties.Value("name")
	name.MinLength = &one
	sc.Properties.Set("name", &name)

	accessKey := *sc.Properties.Value("access_key")
	accessKey.MinLength = &one
	sc.Properties.Set("access_key", &accessKey)

	secretKey := *sc.Properties.Value("secret_key")
	secretKey.MinLength = &one
	sc.Properties.Set("secret_key", &secretKey)

	regions := *sc.Properties.Value("regions").OneOf[0]
	regions.MinItems = &one
	sc.Properties.Set("regions", &regions)

	sc.Required = append(sc.Required, "name", "access_key", "secret_key", "regions")
}

//go:embed schema.json
var JSONSchema string
