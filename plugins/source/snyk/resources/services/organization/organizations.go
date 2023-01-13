package organization

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organizations",
		Description: `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.SingleOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Organization{}),
		Columns: []schema.Column{
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
