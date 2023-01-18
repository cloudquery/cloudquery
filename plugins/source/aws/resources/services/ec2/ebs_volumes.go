package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EbsVolumes() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_ebs_volumes",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Volume.html`,
		Resolver:    fetchEc2EbsVolumes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.Volume{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveEbsVolumeArn,
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
