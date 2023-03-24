package client

import (
	"fmt"
	"strconv"
	"strings"
)

type Spec struct {
	DeveloperToken  string   `json:"developer_token,omitempty"`
	LoginCustomerID string   `json:"login_customer_id,omitempty"`
	Customers       []string `json:"customers,omitempty"`

	OAuth *oauthSpec `json:"oauth,omitempty"`
}

func (s *Spec) validate() error {
	if len(s.DeveloperToken) == 0 {
		return fmt.Errorf("missing required \"developer_token\" value")
	}

	if len(s.LoginCustomerID) > 0 {
		id, err := cleanID(s.LoginCustomerID)
		if err != nil {
			return fmt.Errorf("failed to parse \"login_customer_id\": %w", err)
		}
		s.LoginCustomerID = id
	}

	for i, customer := range s.Customers {
		id, err := cleanID(customer)
		if err != nil {
			return fmt.Errorf("failed to parse \"customers\": %w", err)
		}
		s.Customers[i] = id
	}

	return s.OAuth.validate()
}

func cleanID(id string) (string, error) {
	clean := strings.ReplaceAll(id, "-", "")
	if _, err := strconv.ParseInt(clean, 10, 64); err != nil {
		return "", fmt.Errorf("failed to parse %q: %w", id, err)
	}
	return clean, nil
}
