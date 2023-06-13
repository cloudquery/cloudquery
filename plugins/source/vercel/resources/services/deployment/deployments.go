package deployment

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Deployments() *schema.Table {
	return &schema.Table{
		Name:          "vercel_deployments",
		Resolver:      fetchDeployments,
		Transform:     client.TransformWithStruct(&vercel.Deployment{}),
		Multiplex:     client.TeamMultiplex,
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name:       "uid",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			DeploymentChecks(),
		},
	}
}
