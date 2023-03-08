package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VpcEndpointServices() *schema.Table {
	tableName := "aws_ec2_vpc_endpoint_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html`,
		Resolver:    fetchEc2VpcEndpointServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.ServiceDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveVpcEndpointServiceArn,
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
