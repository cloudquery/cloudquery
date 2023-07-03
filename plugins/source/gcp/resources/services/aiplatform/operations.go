package aiplatform

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
)

func Operations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_operations",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.operations#Operation`,
		Resolver:    fetchOperations,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Operation{}, transformers.WithPrimaryKeys("Name")),
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

func fetchOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListOperationsRequest{
		Name: "projects/" + c.ProjectId + "/locations/-",
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewPipelineClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListOperations(ctx, req, c.CallOptions...)
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
