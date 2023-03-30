package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func clusterInstanceFleets() *schema.Table {
	tableName := "aws_emr_cluster_instance_fleets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleet.html`,
		Resolver:    fetchClusterInstanceFleets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:   client.TransformWithStruct(&types.InstanceFleet{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
