package cloudiot

import (
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func DeviceConfigs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudiot_device_configs",
		Description: `https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.configVersions#DeviceConfig`,
		Resolver:    fetchDeviceConfigs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudiot.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.DeviceConfig{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "device_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
