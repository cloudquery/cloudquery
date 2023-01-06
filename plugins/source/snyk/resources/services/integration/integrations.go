package integration

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Integrations() *schema.Table {
	return &schema.Table{
		Name:                "snyk_integrations",
		Description:         `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Integration`,
		Resolver:            fetchIntegrations,
		PreResourceResolver: getIntegration,
		Multiplex:           client.ByOrganization,
		Transform:           transformers.TransformWithStruct(&snyk.Integration{}),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganizationID,
			},
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: getIntegrationSettings,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
