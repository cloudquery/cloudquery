package ec2

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func KeyPairs() *schema.Table {
	tableName := "aws_ec2_key_pairs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_KeyPairInfo.html`,
		Resolver:    fetchEc2KeyPairs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.KeyPairInfo{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveKeyPairArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2KeyPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeKeyPairsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	output, err := svc.DescribeKeyPairs(ctx, &config, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.KeyPairs
	return nil
}

func resolveKeyPairArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.KeyPairInfo)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "key-pair/" + aws.ToString(item.KeyPairId),
	}
	return resource.Set(c.Name, a.String())
}
