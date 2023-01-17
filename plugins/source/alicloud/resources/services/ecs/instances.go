package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_ecs_instances",
		Resolver:    fetchEcsInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(client.ServiceECS),
		Description: "https://www.alibabacloud.com/help/en/elastic-compute-service/latest/describeinstances#t9865.html",
		Transform: transformers.TransformWithStruct(
			&ecs.Instance{},
			transformers.WithPrimaryKeys(
				"RegionId", "InstanceId",
			),
		),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
