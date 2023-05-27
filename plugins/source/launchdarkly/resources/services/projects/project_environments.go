package projects

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func ProjectEnvironments() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_project_environments",
		Description: `https://apidocs.launchdarkly.com/tag/Environments#operation/getEnvironment`,
		Resolver:    fetchProjectEnvironments,
		Transform:   client.TransformWithStruct(&ldapi.Environment{}, transformers.WithPrimaryKeys("Id"), transformers.WithSkipFields("Links")),
		Columns: schema.ColumnList{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
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
