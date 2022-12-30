package resources

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/companies"
)

func recipes() []*Resource {
	return []*Resource{
		{
			Name:       "hubspot_companies",
			Service:    "companies",
			SubService: "companies",
			Struct:     new(companies.SimplePublicObject),
			PKColumns:  []string{"id"},
		},
	}
}
