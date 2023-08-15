package ec2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AccountAttributes() *schema.Table {
	tableName := "aws_ec2_account_attributes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccountAttribute.html`,
		Resolver:    fetchAccountAttributes,
		Multiplex:   client.AccountMultiplex(tableName),
		Transform:   transformers.TransformWithStruct(&types.AccountAttribute{}, transformers.WithPrimaryKeys("AttributeName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "partition",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSPartition,
			},
		},
	}
}
func fetchAccountAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	output, err := svc.DescribeAccountAttributes(ctx, &ec2.DescribeAccountAttributesInput{}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.AccountAttributes
	return nil
}
