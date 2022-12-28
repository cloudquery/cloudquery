package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Organizations() []*Resource {
	alert := dependabotAlert()
	alert.Service = "organizations"
	alert.TableName = "organization_dependabot_alerts"

	sec := dependabotSecret()
	sec.Service = "organizations"
	sec.TableName = "organization_dependabot_secrets"

	return []*Resource{
		{
			TableName:    "organizations",
			Service:      "organizations",
			SubService:   "organizations",
			Struct:       new(github.Organization),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
			Relations:    []string{"Alerts()", "Secrets()", "Members()"},
		},
		alert,
		sec,
		{
			TableName:  "organization_members",
			Service:    "organizations",
			SubService: "members",
			Struct:     new(github.User),
			PKColumns:  []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{
				orgColumn, // we can use orgColumn here
				{
					Name:     "membership",
					Type:     schema.TypeJSON,
					Resolver: "resolveMembership",
				},
			},
			Multiplex: "", // we skip multiplexing here as it's a relation
		},
	}
}
