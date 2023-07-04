package deployment

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func DeploymentChecks() *schema.Table {
	return &schema.Table{
		Name:          "vercel_deployment_checks",
		Resolver:      fetchDeploymentChecks,
		Transform:     client.TransformWithStruct(&vercel.DeploymentCheck{}),
		Multiplex:     client.TeamMultiplex,
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name:       "deployment_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("uid"),
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
