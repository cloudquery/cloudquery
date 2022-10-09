package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListRegistriesInput{MaxResults: aws.Int32(100)}
	for {
		result, err := svc.ListRegistries(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Registries
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueRegistryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.RegistryListItem)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: r.RegistryArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
func fetchGlueRegistrySchemas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RegistryListItem)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListSchemasInput{
		RegistryId: &types.RegistryId{RegistryArn: r.RegistryArn},
		MaxResults: aws.Int32(100),
	}
	for {
		result, err := svc.ListSchemas(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Schemas

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func getRegistrySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	item := resource.Item.(types.SchemaListItem)

	s, err := svc.GetSchema(ctx, &glue.GetSchemaInput{SchemaId: &types.SchemaId{SchemaArn: item.SchemaArn}})
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
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

func fetchGlueRegistrySchemaVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	s := parent.Item.(*glue.GetSchemaOutput)
	svc := cl.Services().Glue
	schemaId := types.SchemaId{
		SchemaArn: s.SchemaArn,
	}
	input := glue.ListSchemaVersionsInput{
		SchemaId:   &schemaId,
		MaxResults: aws.Int32(100),
	}
	for {
		result, err := svc.ListSchemaVersions(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Schemas

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func getRegistrySchemaVersion(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	item := resource.Item.(types.SchemaVersionListItem)

	s, err := svc.GetSchemaVersion(ctx, &glue.GetSchemaVersionInput{
		SchemaVersionId: item.SchemaVersionId,
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
	for {
		result, err := svc.QuerySchemaVersionMetadata(ctx, input)
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
