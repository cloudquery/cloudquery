package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Accounts() *schema.Table {
	tableName := "aws_organizations_accounts"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:  fetchOrganizationsAccounts,
		Transform: transformers.TransformWithStruct(&types.Account{}, transformers.WithPrimaryKeyComponents("Arn")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveAccountTags,
			},
		},
		Relations: []*schema.Table{
			organizationalAccountParents(),
		},
	}
}

func fetchOrganizationsAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	var input organizations.ListAccountsInput
	paginator := organizations.NewListAccountsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
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
	paginator := organizations.NewListTagsForResourcePaginator(cl.Services(client.AWSServiceOrganizations).Organizations, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set("tags", client.TagsToMap(tags))
}
