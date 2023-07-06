package ssm

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Instances() *schema.Table {
	tableName := "aws_ssm_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html`,
		Resolver:    fetchSsmInstances,
		Transform:   transformers.TransformWithStruct(&types.InstanceInformation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveInstanceARN,
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			instanceComplianceItems(),
			instancePatches(),
		},
	}
}

func fetchSsmInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm
	paginator := ssm.NewDescribeInstanceInformationPaginator(svc, &ssm.DescribeInstanceInformationInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.InstanceInformationList
	}
	return nil
}

func resolveInstanceARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "ssm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("managed-instance/%s", aws.ToString(instance.InstanceId)),
	}.String())
}
