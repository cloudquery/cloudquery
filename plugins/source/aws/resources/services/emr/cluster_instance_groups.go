package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func clusterInstanceGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_emr_cluster_instance_groups",
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroup.html`,
		Resolver:    fetchClusterInstanceGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.InstanceGroup{}, transformers.WithPrimaryKeys("Id")),
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
