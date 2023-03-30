package dms

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReplicationInstances() *schema.Table {
	tableName := "aws_dms_replication_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dms"),
		Transform:   client.TransformWithStruct(&models.ReplicationInstanceWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
