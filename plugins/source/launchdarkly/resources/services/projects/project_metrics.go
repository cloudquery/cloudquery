package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func ProjectMetrics() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_project_metrics",
		Description: `https://apidocs.launchdarkly.com/tag/Metrics#operation/getMetrics`,
		Resolver:    fetchProjectMetrics,
		Transform:   transformers.TransformWithStruct(&ldapi.MetricListingRep{}, client.SharedTransformers(transformers.WithPrimaryKeys("Id"), transformers.WithSkipFields("Links", "Site"))...),
		Columns: schema.ColumnList{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchProjectMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(ldapi.Project)

	list, b, err := cl.Services.MetricsApi.GetMetrics(ctx, p.Key).Expand("experimentCount").Execute()
	if err != nil {
		return err
	}
	b.Body.Close()

	res <- list.Items

	return nil
}
