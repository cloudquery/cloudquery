package dependency

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Dependencies() *schema.Table {
	return &schema.Table{
		Name:        "snyk_dependencies",
		Description: `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Dependency`,
		Resolver:    fetchDependencies,
		Multiplex:   client.ByOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Dependency{}),
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
