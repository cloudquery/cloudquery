package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func ProjectFlags() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_project_flags",
		Description: `https://apidocs.launchdarkly.com/tag/Feature-flags#operation/getFeatureFlags`,
		Resolver:    fetchProjectFlags,
		Transform:   transformers.TransformWithStruct(&ldapi.FeatureFlag{}, client.SharedTransformers(transformers.WithPrimaryKeys("Key"), transformers.WithSkipFields("Links"))...),
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

func fetchProjectFlags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(ldapi.Project)

	for _, arch := range []bool{false, true} {
		list, b, err := cl.Services.FeatureFlagsApi.GetFeatureFlags(ctx, p.Key).Archived(arch).Execute()
		if err != nil {
			return err
		}
		b.Body.Close()

		res <- list.Items
	}

	return nil
}
