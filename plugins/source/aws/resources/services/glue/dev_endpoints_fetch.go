package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueDevEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetDevEndpointsInput{}
	for {
		result, err := svc.GetDevEndpoints(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- result.DevEndpoints
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueDevEndpointArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName)))
	return resource.Set(c.Name, arn)
}
func resolveGlueDevEndpointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func devEndpointARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "devEndpoint", name)
}
