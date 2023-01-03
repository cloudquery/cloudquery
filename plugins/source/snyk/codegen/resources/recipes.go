package resources

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func recipes() []*Resource {
	return []*Resource{
		{
			Name:       "snyk_dependencies",
			Struct:     new(snyk.Dependency),
			Service:    "dependency",
			SubService: "dependencies",
			Multiplex:  "client.ByOrganization",
		},
		{
			Name:      "snyk_organizations",
			Struct:    new(snyk.Organization),
			Service:   "organization",
			Multiplex: "client.SingleOrganization",
		},
		{
			Name:    "snyk_integrations",
			Struct:  new(snyk.Integration),
			Service: "integration",
			ExtraColumns: codegen.ColumnDefinitions{
				organizationIDCol,
				codegen.ColumnDefinition{
					Name:     "settings",
					Type:     schema.TypeJSON,
					Resolver: "getIntegrationSettings",
				},
			},
			Multiplex:   "client.ByOrganization",
			PreResolver: "getIntegration",
		},
		{
			Name:         "snyk_projects",
			Struct:       new(snyk.Project),
			Service:      "project",
			ExtraColumns: codegen.ColumnDefinitions{organizationIDCol},
			Multiplex:    "client.ByOrganization",
		},
	}
}
