package glue

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Registries() *schema.Table {
	tableName := "aws_glue_registries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_RegistryListItem.html`,
		Resolver:    fetchGlueRegistries,
		Transform:   transformers.TransformWithStruct(&types.RegistryListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGlueRegistryTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("RegistryArn"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			registrySchemas(),
		},
	}
}

func fetchGlueRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	paginator := glue.NewListRegistriesPaginator(svc, &glue.ListRegistriesInput{
		MaxResults: aws.Int32(100),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Registries
	}
	return nil
}

func resolveGlueRegistryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.RegistryListItem)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: r.RegistryArn,
	}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
