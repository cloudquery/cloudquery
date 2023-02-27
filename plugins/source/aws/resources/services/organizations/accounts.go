package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "aws_organizations_accounts",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html`,
		Resolver:    fetchOrganizationsAccounts,
		Transform:   transformers.TransformWithStruct(&types.Account{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveAccountTags,
			},
		},
	}
}

func fetchOrganizationsAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Organizations
	var input organizations.ListAccountsInput
	paginator := organizations.NewListAccountsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Accounts
	}
	return nil
}
func resolveAccountTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	cl := meta.(*client.Client)
	account := resource.Item.(types.Account)
	var tags []types.Tag
	input := organizations.ListTagsForResourceInput{
		ResourceId: account.Id,
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
