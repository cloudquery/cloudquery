package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func ProjectEnvironments() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_project_environments",
		Description: `https://apidocs.launchdarkly.com/tag/Environments#operation/getEnvironment`,
		Resolver:    fetchProjectEnvironments,
		Transform:   transformers.TransformWithStruct(&ldapi.Environment{}, client.SharedTransformers(transformers.WithPrimaryKeys("Id"), transformers.WithSkipFields("Links"))...),
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

func fetchProjectEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(ldapi.Project)

	if p.Environments == nil {
		return nil
	}

	res <- p.Environments.Items

	return nil
}
