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

func registrySchemaVersions() *schema.Table {
	tableName := "aws_glue_registry_schema_versions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_GetSchemaVersion.html`,
		Resolver:            fetchGlueRegistrySchemaVersions,
		PreResourceResolver: getRegistrySchemaVersion,
		Transform:           transformers.TransformWithStruct(&glue.GetSchemaVersionOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "registry_schema_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "metadata",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGlueRegistrySchemaVersionMetadata,
			},
		},
	}
}

func fetchGlueRegistrySchemaVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := parent.Item.(*glue.GetSchemaOutput)
	svc := cl.Services().Glue
	input := glue.ListSchemaVersionsInput{
		SchemaId: &types.SchemaId{
			SchemaArn: s.SchemaArn,
		},
		MaxResults: aws.Int32(100),
	}
	paginator := glue.NewListSchemaVersionsPaginator(svc, &input)
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

func getRegistrySchemaVersion(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	item := resource.Item.(types.SchemaVersionListItem)

	s, err := svc.GetSchemaVersion(ctx, &glue.GetSchemaVersionInput{
		SchemaVersionId: item.SchemaVersionId,
	}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = s
	return nil
}

func resolveGlueRegistrySchemaVersionMetadata(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	s := resource.Item.(*glue.GetSchemaVersionOutput)
	input := &glue.QuerySchemaVersionMetadataInput{
		SchemaVersionId: s.SchemaVersionId,
	}
	metadata := make(map[string]types.MetadataInfo)
	// No paginator available
	for {
		result, err := svc.QuerySchemaVersionMetadata(ctx, input, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		for k, v := range result.MetadataInfoMap {
			metadata[k] = v
		}

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return resource.Set(c.Name, metadata)
}
