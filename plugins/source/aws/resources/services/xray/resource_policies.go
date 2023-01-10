package xray

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourcePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_resource_policies",
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_ResourcePolicy.html`,
		Resolver:    fetchXrayResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: createXrayResourcePolicyArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
