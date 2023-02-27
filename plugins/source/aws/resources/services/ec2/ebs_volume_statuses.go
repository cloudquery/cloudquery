package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EbsVolumesStatuses() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_ebs_volume_statuses",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VolumeStatusItem.html`,
		Resolver:    fetchEc2EbsVolumeStatuses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.VolumeStatusItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "volume_arn",
				Type:     schema.TypeString,
				Resolver: resolveEbsVolumeStatusArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
