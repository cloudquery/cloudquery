package functions

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/functions/apiv1/functionspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	functions "cloud.google.com/go/functions/apiv1"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:        "gcp_functions_functions",
		Description: `https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions#CloudFunction`,
		Resolver:    fetchFunctions,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudfunctions.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.CloudFunction{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListFunctionsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := functions.NewCloudFunctionsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListFunctions(ctx, req, c.CallOptions...)
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
