package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	"reflect"
	"strings"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "alicloud_ecs_instances",
		Title:       "Alibaba Cloud Elastic Compute Service (ECS) Instances",
		Resolver:    fetchEcsInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(client.ServiceECS),
		Description: "https://www.alibabacloud.com/help/en/elastic-compute-service/latest/describeinstances#t9865.html",
		Transform: transformers.TransformWithStruct(
			&ecs.Instance{},
			transformers.WithPrimaryKeys(
				"RegionId", "InstanceId",
			),
			transformers.WithTypeTransformer(func(f reflect.StructField) (arrow.DataType, error) {
				if strings.HasSuffix(f.Name, "Time") {
					return arrow.FixedWidthTypes.Timestamp_us, nil
				}
				return nil, nil
			}),
			transformers.WithResolverTransformer(func(f reflect.StructField, path string) schema.ColumnResolver {
				if strings.HasSuffix(f.Name, "Time") {
					return client.TimestampResolver("2006-01-02T15:04Z", path)
				}
				return transformers.DefaultResolverTransformer(f, path)
			}),
		),
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccount,
				PrimaryKey: true,
			},
		},
	}
}
