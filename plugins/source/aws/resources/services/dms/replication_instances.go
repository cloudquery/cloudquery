package dms

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReplicationInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_dms_replication_instances",
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dms"),
		Transform:  transformers.TransformWithStruct(&models.ReplicationInstanceWrapper{}),
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
				Resolver: schema.PathResolver("ReplicationInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
