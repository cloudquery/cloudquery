package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func ProjectFlagDefaults() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_project_flag_defaults",
		Description: `https://apidocs.launchdarkly.com/tag/Projects#operation/getFlagDefaultsByProject`,
		Resolver:    fetchProjectFlagDefaults,
		Transform:   transformers.TransformWithStruct(&ldapi.FlagDefaultsRep{}, client.SharedTransformers(transformers.WithSkipFields("Links"))...),
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

func fetchProjectFlagDefaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(ldapi.Project)

	list, b, err := cl.Services.ProjectsApi.GetFlagDefaultsByProject(ctx, p.Key).Execute()
	if err != nil {
		return err
	}
	b.Body.Close()

	res <- list

	return nil
}
