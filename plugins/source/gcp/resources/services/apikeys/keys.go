// Code generated by codegen; DO NOT EDIT.

package apikeys

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/api/apikeys/v2"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/apikeys/apiv2"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:      "gcp_apikeys_keys",
		Resolver:  fetchKeys,
		Multiplex: client.ProjectMultiplex,
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
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uid"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "key_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyString"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "delete_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeleteTime"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "restrictions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Restrictions"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
		},
	}
}

func fetchKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListKeysRequest{
		Parent: "projects/" + c.ProjectId + "/locations/global",
	}
	gcpClient, err := apikeys.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListKeys(ctx, req, c.CallOptions...)
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
