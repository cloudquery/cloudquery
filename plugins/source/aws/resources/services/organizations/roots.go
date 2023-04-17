package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Roots() *schema.Table {
	tableName := "aws_organizations_roots"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Root.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:  fetchOrganizationsRoots,
		Transform: transformers.TransformWithStruct(&types.Root{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			{
				Name:            "request_account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRootTags,
			},
		},
	}
}
func fetchOrganizationsRoots(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Organizations
	var input organizations.ListRootsInput
	paginator := organizations.NewListRootsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Roots
	}
	return nil
}

func resolveRootTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	cl := meta.(*client.Client)
	root := resource.Item.(types.Root)
	var tags []types.Tag
	input := organizations.ListTagsForResourceInput{
		ResourceId: root.Id,
	}
	paginator := organizations.NewListTagsForResourcePaginator(cl.Services().Organizations, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set("tags", client.TagsToMap(tags))
}
