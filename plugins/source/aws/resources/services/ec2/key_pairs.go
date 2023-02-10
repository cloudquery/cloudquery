package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func KeyPairs() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_key_pairs",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_KeyPairInfo.html`,
		Resolver:    fetchEc2KeyPairs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.KeyPairInfo{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveKeyPairArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
