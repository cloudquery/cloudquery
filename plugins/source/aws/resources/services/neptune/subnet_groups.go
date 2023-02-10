package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_neptune_subnet_groups",
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-subnets.html#DescribeDBSubnetGroups`,
		Resolver:    fetchNeptuneSubnetGroups,
		Transform:   transformers.TransformWithStruct(&types.DBSubnetGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneSubnetGroupTags,
			},
			{
				Name:     "db_subnet_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupDescription"),
			},
			{
				Name:     "db_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupName"),
			},
			{
				Name:     "subnet_group_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetGroupStatus"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
