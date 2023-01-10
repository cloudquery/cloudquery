package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrafficPolicyVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_traffic_policy_versions",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html`,
		Resolver:    fetchRoute53TrafficPolicyVersions,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicy{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "traffic_policy_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "document",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Document"),
			},
		},
	}
}
