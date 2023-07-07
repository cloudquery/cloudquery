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

func registrySchemas() *schema.Table {
	tableName := "aws_glue_registry_schemas"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_GetSchema.html`,
		Resolver:            fetchGlueRegistrySchemas,
		PreResourceResolver: getRegistrySchema,
		Transform:           transformers.TransformWithStruct(&glue.GetSchemaOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("SchemaArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGlueRegistrySchemaTags,
			},
		},

		Relations: []*schema.Table{
			registrySchemaVersions(),
		},
	}
}

func fetchGlueRegistrySchemas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RegistryListItem)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListSchemasInput{
		RegistryId: &types.RegistryId{RegistryArn: r.RegistryArn},
		MaxResults: aws.Int32(100),
	}
	paginator := glue.NewListSchemasPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- result.Schemas
	}
	return nil
}

func getRegistrySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	item := resource.Item.(types.SchemaListItem)

	s, err := svc.GetSchema(ctx, &glue.GetSchemaInput{SchemaId: &types.SchemaId{SchemaArn: item.SchemaArn}}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = s
	return nil
}

func resolveGlueRegistrySchemaTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	s := resource.Item.(*glue.GetSchemaOutput)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: s.SchemaArn,
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
