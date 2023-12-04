package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

type arnProvider func(cl *client.Client, resource *schema.Resource) string

func tagsCol(arnProvider arnProvider) schema.Column {
	return schema.Column{
		Name:     "tags",
		Type:     sdkTypes.ExtensionTypes.JSON,
		Resolver: resolveTags(arnProvider),
	}
}

func resolveTags(arnProvider arnProvider) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		cl := meta.(*client.Client)
		svc := cl.Services(client.AWSServiceGlue).Glue
		input := glue.GetTagsInput{ResourceArn: aws.String(arnProvider(cl, resource))}

		response, err := svc.GetTags(ctx, &input, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		return resource.Set(c.Name, response.Tags)
	}
}
