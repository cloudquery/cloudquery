package ec2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Eips() *schema.Table {
	tableName := "aws_ec2_eips"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html`,
		Resolver:    fetchEc2Eips,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.Address{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2Eips(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	output, err := svc.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{
		Filters: []types.Filter{{Name: aws.String("domain"), Values: []string{"vpc"}}},
	}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.Addresses
	return nil
}
