package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ResourcePolicies() *schema.Table {
	tableName := "aws_cloudwatchlogs_resource_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html`,
		Resolver:    fetchCloudwatchlogsResourcePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PolicyDocument"),
			},
		},
	}
}

func fetchCloudwatchlogsResourcePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeResourcePoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	// No paginator available
	for {
		response, err := svc.DescribeResourcePolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ResourcePolicies
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
