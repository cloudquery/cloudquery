package organizations

import (
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourcePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_organization_resource_policies",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html`,
		Resolver:    fetchOrganizationsResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
