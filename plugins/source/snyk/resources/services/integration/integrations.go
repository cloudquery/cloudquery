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
		Transform:           transformers.TransformWithStruct(&snyk.Integration{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrganizationID,
			{
				Name:     "settings",
				Type:     schema.TypeJSON,
				Resolver: getIntegrationSettings,
			},
		},
	}
}
