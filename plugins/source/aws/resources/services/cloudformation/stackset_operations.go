package cloudformation

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func stackSetOperations() *schema.Table {
	return &schema.Table{
		Name:                "aws_cloudformation_stack_set_operations",
		Description:         `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html`,
		Resolver:            fetchCloudformationStackSetOperations,
		PreResourceResolver: getStackSetOperation,
		Multiplex:           client.ServiceAccountRegionMultiplexer("cloudformation"),
		Transform:           transformers.TransformWithStruct(&types.StackSetOperation{}, transformers.WithPrimaryKeys("StackSetId", "OperationId", "CreationTimestamp")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationId"),
			},
			{
				Name:     "stack_set_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("stack_set_arn"),
			},
		},

		Relations: []*schema.Table{
			stackSetOperationResults(),
		},
	}
}
