package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func OrganizationsAccounts() *schema.Table {
	return &schema.Table{
		Name:         "aws_organizations_accounts",
		Description:  "Contains information about an AWS account that is a member of an organization.",
		Resolver:     fetchOrganizationsAccounts,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "email",
				Description: "The email address associated with the AWS account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_id",
				Description: "The unique identifier (ID) of the account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "joined_method",
				Description: "The method by which the account joined the organization.",
				Type:        schema.TypeString,
			},
			{
				Name:        "joined_timestamp",
				Description: "The date the account became a part of the organization.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The friendly name of the account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the account in the organization.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchOrganizationsAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
