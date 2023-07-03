package secretmanager

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_secretmanager_secrets",
		Description: `https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret`,
		Resolver:    fetchSecrets,
		Multiplex:   client.ProjectMultiplexEnabledServices("secretmanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Secret{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListSecretsRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := secretmanager.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListSecrets(ctx, req, c.CallOptions...)
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
