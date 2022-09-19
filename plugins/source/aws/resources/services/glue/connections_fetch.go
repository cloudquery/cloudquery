package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.GetConnectionsInput{}
	for {
		output, err := svc.GetConnections(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.ConnectionList

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveGlueConnectionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Connection)
	arn := aws.String(connectionARN(cl, &r))
	return resource.Set(c.Name, arn)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func connectionARN(cl *client.Client, c *types.Connection) string {
	return cl.ARN(client.GlueService, "connection", *c.Name)
}
