package livestream

import (
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
