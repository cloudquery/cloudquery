package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ByoipCidrs() *schema.Table {
	tableName := "aws_ec2_byoip_cidrs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html`,
		Resolver:    fetchEc2ByoipCidrs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.ByoipCidr{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "cidr",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
	}
}

func fetchEc2ByoipCidrs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	// DescribeByoipCidrs does not work in next regions, so we ignore them.
	if _, ok := map[string]struct{}{
		"cn-north-1":     {},
		"cn-northwest-1": {},
	}[cl.Region]; ok {
		return nil
	}
	svc := cl.Services().Ec2
	config := ec2.DescribeByoipCidrsInput{
		MaxResults: aws.Int32(100),
	}
	paginator := ec2.NewDescribeByoipCidrsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ByoipCidrs
	}
	return nil
}
