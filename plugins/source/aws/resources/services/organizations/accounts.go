package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "aws_organizations_accounts",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html`,
		Resolver:    fetchOrganizationsAccounts,
		Transform:   transformers.TransformWithStruct(&types.Account{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
	for {
		response, err := svc.ListAccounts(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Accounts
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func resolveAccountTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	cl := meta.(*client.Client)
	account := resource.Item.(types.Account)
	var tags []types.Tag
	input := organizations.ListTagsForResourceInput{
		ResourceId: account.Id,
		NextToken:  nil,
	}
	for {
		response, err := cl.Services().Organizations.ListTagsForResource(ctx, &input)
		if err != nil {
			return err
		}
		tags = append(tags, response.Tags...)
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set("tags", client.TagsToMap(tags))
}
