package client

import "fmt"

type Spec struct {
	Accounts          []AccountSpec `json:"accounts,omitempty"`
	BillHistoryMonths int           `json:"bill_history_months,omitempty"`
}

type AccountSpec struct {
	Name      string   `json:"name,omitempty"`
	AccessKey string   `json:"access_key,omitempty"`
	SecretKey string   `json:"secret_key,omitempty"`
	RegionIDs []string `json:"region_ids,omitempty"`
}

func (s Spec) Validate() error {
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
		if len(account.RegionIDs) == 0 {
			return fmt.Errorf("missing region_ids in account configuration for account %s", account.Name)
		}
	}
	return nil
}
