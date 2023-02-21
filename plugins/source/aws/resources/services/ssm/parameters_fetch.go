package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm
	params := ssm.DescribeParametersInput{}
	for {
		output, err := svc.DescribeParameters(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.Parameters
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}

func resolveParameterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	parameter := resource.Item.(types.Parameter)

	svc := meta.(*client.Client).Services().SystemsManager
	response, err := svc.ListTagsForResource(ctx, &parameter.ListTagsForResourceInput{
		RessourceType: "Parameter",
		ResourceId: parameter.path,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.TagList))
}
