package cloudiot

import (
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func DeviceConfigs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudiot_device_configs",
		Description: `https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.configVersions#DeviceConfig`,
		Resolver:    fetchDeviceConfigs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudiot.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.DeviceConfig{}, transformers.WithPrimaryKeys("Version")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
			{
				Name:       "device_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
		},
	}
}
