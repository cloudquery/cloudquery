package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func InstanceCreditSpecifications() *schema.Table {
	tableName := "aws_ec2_instance_credit_specifications"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.DescribeInstanceCreditSpecifications`,
		Resolver:    fetchEc2InstanceCreditSpecifications,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.InstanceCreditSpecification{}, transformers.WithPrimaryKeys("InstanceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchEc2InstanceCreditSpecifications(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config ec2.DescribeInstanceCreditSpecificationsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	paginator := ec2.NewDescribeInstanceCreditSpecificationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.InstanceCreditSpecifications
	}
	return nil
}
