package cloudformation

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Stacks() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudformation_stacks",
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html`,
		Resolver:    fetchCloudformationStacks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudformation"),
		Transform:  transformers.TransformWithStruct(&types.Stack{}),
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			StackResources(),
		},
	}
}
