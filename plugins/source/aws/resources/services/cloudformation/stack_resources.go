package cloudformation

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func stackResources() *schema.Table {
	tableName := "aws_cloudformation_stack_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html`,
		Resolver:    fetchCloudformationStackResources,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform:   client.TransformWithStruct(&types.StackResourceSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
