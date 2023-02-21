package organizations

import (
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourcePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_organizations_resource_policies",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html`,
		Resolver:    fetchOrganizationsResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
