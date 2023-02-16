package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_network_interfaces",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterface.html`,
		Resolver:    fetchEc2NetworkInterfaces,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.NetworkInterface{}, transformers.WithSkipFields("TagSet")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveNetworkInterfaceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagSet"),
			},
		},
	}
}
