package redshift

import (
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_subnet_groups",
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html`,
		Resolver:    fetchRedshiftSubnetGroups,
		Transform:   transformers.TransformWithStruct(&types.ClusterSubnetGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveSubnetGroupArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
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
