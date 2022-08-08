package organizations

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource accounts --config gen.hcl --output .
func Accounts() *schema.Table {
	return &schema.Table{
		Name:          "aws_organizations_accounts",
		Description:   "Contains information about an AWS account that is a member of an organization",
		Resolver:      fetchOrganizationsAccounts,
		Multiplex:     client.AccountMultiplex,
		DeleteFilter:  client.DeleteAccountFilter,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		IgnoreInTests: true,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "tags",
				Description: "The AWS tags of the resource.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveOrganizationsAccountTags,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the account",
				Type:        schema.TypeString,
			},
			{
				Name:        "email",
				Description: "The email address associated with the AWS account",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier (ID) of the account",
				Type:        schema.TypeString,
			},
			{
				Name:        "joined_method",
				Description: "The method by which the account joined the organization",
				Type:        schema.TypeString,
			},
			{
				Name:        "joined_timestamp",
				Description: "The date the account became a part of the organization",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The friendly name of the account",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the account in the organization",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchOrganizationsAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Organizations
	var input organizations.ListAccountsInput
	for {
		response, err := svc.ListAccounts(ctx, &input)
		var ae smithy.APIError
		if errors.As(err, &ae) {
			switch ae.ErrorCode() {
			case "AWSOrganizationsNotInUseException", "AccessDeniedException":
				// This is going to happen most probably due to account not being the root organizational account
				// so it's better to ignore it completely as it happens basically on every account
				// otherwise it screws up with dev experience and with our tests
				meta.Logger().Warn("account is probably not the root organization account https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccounts.html")
				return nil
			}
		}
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Accounts
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveOrganizationsAccountTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	cl := meta.(*client.Client)
	account := resource.Item.(types.Account)
	allTags := make(map[string]*string)
	input := organizations.ListTagsForResourceInput{
		ResourceId: account.Id,
		NextToken:  nil,
	}
	for {
		response, err := cl.Services().Organizations.ListTagsForResource(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, t := range response.Tags {
			allTags[*t.Key] = t.Value
		}
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}
	return diag.WrapError(resource.Set("tags", allTags))
}
