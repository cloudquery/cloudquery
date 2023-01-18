package cloudiot

import (
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudiot_devices",
		Description: `https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices#Device`,
		Resolver:    fetchDevices,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudiot.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Device{}, client.Options()...),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			DeviceConfigs(), DeviceStates(),
		},
	}
}
