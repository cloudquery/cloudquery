package baremetalsolution

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
)

func VolumeLuns() *schema.Table {
	return &schema.Table{
		Name:        "gcp_baremetalsolution_volume_luns",
		Description: `https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes.luns#Lun`,
		Resolver:    fetchVolumeLuns,
		Multiplex:   client.ProjectMultiplexEnabledServices("baremetalsolution.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Lun{}, client.Options()...),
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
	}
}

func fetchVolumeLuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListLunsRequest{
		Parent: parent.Item.(*pb.Volume).Name,
	}
	gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListLuns(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
