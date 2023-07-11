package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func spotFleetInstances() *schema.Table {
	tableName := "aws_ec2_spot_fleet_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ActiveInstance.html`,
		Resolver:    fetchEC2SpotFleetInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.ActiveInstance{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveActiveInstanceArn,
				PrimaryKey: true,
			},
			{
				Name:     "spot_fleet_request_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("spot_fleet_request_id"),
			},
		},
	}
}

func fetchEC2SpotFleetInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.SpotFleetRequestConfig)

	config := ec2.DescribeSpotFleetInstancesInput{
		SpotFleetRequestId: p.SpotFleetRequestId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	// No paginator available
	for {
		output, err := svc.DescribeSpotFleetInstances(ctx, &config, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.ActiveInstances
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveActiveInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ActiveInstance)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "instance/" + aws.ToString(item.InstanceId),
	}
	return resource.Set(c.Name, a.String())
}
