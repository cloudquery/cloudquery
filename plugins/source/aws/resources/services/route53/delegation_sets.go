package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DelegationSets() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_delegation_sets",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_DelegationSet.html`,
		Resolver:    fetchRoute53DelegationSets,
		Transform:   transformers.TransformWithStruct(&types.DelegationSet{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveDelegationSetArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
