package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_projects",
		Description: `https://apidocs.launchdarkly.com/tag/Projects#operation/getProjects`,
		Resolver:    fetchProjects,
		Transform:   client.TransformWithStruct(&ldapi.Project{}, transformers.WithPrimaryKeys("Id"), transformers.WithSkipFields("Environments", "Links")),
		Relations: []*schema.Table{
			ProjectEnvironments(),
			ProjectFlagDefaults(),
			ProjectFlags(),
			ProjectMetrics(),
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	const limit = 20
	ofs := int64(0)
	for {
		list, b, err := cl.Services.ProjectsApi.GetProjects(ctx).Sort("createdOn").Expand("environments").Offset(ofs).Limit(limit).Execute()
		if err != nil {
			return err
		}
		b.Body.Close()

		res <- list.Items

		if len(list.Items) < limit {
			break
		}

		ofs += limit
	}

	return nil
}
