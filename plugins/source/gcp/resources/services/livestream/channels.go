package livestream

import (
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Channels() *schema.Table {
	return &schema.Table{
		Name:        "gcp_livestream_channels",
		Description: `https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.channels`,
		Resolver:    fetchChannels,
		Multiplex:   client.ProjectMultiplexEnabledServices("livestream.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Channel{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
