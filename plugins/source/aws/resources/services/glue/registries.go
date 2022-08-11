package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource registries --config registries.hcl --output .
func Registries() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_registries",
		Description:  "A structure containing the details for a registry.",
		Resolver:     fetchGlueRegistries,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueRegistryTags,
			},
			{
				Name:        "created_time",
				Description: "The data the registry was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of the registry.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the registry.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegistryArn"),
			},
			{
				Name:        "registry_name",
				Description: "The name of the registry.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the registry.",
				Type:        schema.TypeString,
			},
			{
				Name:        "updated_time",
				Description: "The date the registry was updated.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_registry_schemas",
				Description: "An object that contains minimal details for a schema",
				Resolver:    fetchGlueRegistrySchemas,
				Columns: []schema.Column{
					{
						Name:        "registry_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_registries table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "tags",
						Description: "Resource tags.",
						Type:        schema.TypeJSON,
						Resolver:    resolveGlueRegistrySchemaTags,
					},
					{
						Name:        "compatibility",
						Description: "The compatibility mode of the schema.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_time",
						Description: "The date and time the schema was created.",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_format",
						Description: "The data format of the schema definition",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "A description of schema if specified when created",
						Type:        schema.TypeString,
					},
					{
						Name:        "latest_schema_version",
						Description: "The latest version of the schema associated with the returned schema definition.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "next_schema_version",
						Description: "The next version of the schema associated with the returned schema definition.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "registry_arn",
						Description: "The Amazon Resource Name (ARN) of the registry.",
						Type:        schema.TypeString,
					},
					{
						Name:        "registry_name",
						Description: "The name of the registry.",
						Type:        schema.TypeString,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the schema.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SchemaArn"),
					},
					{
						Name:        "schema_checkpoint",
						Description: "The version number of the checkpoint (the last time the compatibility mode was changed).",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "schema_name",
						Description: "The name of the schema.",
						Type:        schema.TypeString,
					},
					{
						Name:        "schema_status",
						Description: "The status of the schema.",
						Type:        schema.TypeString,
					},
					{
						Name:        "updated_time",
						Description: "The date and time the schema was updated.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_glue_registry_schema_versions",
						Description: "An object containing the details about a schema version",
						Resolver:    fetchGlueRegistrySchemaVersions,
						Columns: []schema.Column{
							{
								Name:        "registry_schema_cq_id",
								Description: "Unique CloudQuery ID of aws_glue_registry_schemas table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "metadata",
								Type:     schema.TypeJSON,
								Resolver: resolveGlueRegistrySchemaVersionMetadata,
							},
							{
								Name:        "created_time",
								Description: "The date and time the schema version was created.",
								Type:        schema.TypeString,
							},
							{
								Name:        "data_format",
								Description: "The data format of the schema definition",
								Type:        schema.TypeString,
							},
							{
								Name:        "schema_arn",
								Description: "The Amazon Resource Name (ARN) of the schema.",
								Type:        schema.TypeString,
							},
							{
								Name:        "schema_definition",
								Description: "The schema definition for the schema ID.",
								Type:        schema.TypeString,
							},
							{
								Name:        "id",
								Description: "The SchemaVersionId of the schema version.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("SchemaVersionId"),
							},
							{
								Name:        "status",
								Description: "The status of the schema version.",
								Type:        schema.TypeString,
							},
							{
								Name:        "version_number",
								Description: "The version number of the schema.",
								Type:        schema.TypeBigInt,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListRegistriesInput{MaxResults: aws.Int32(100)}
	for {
		result, err := svc.ListRegistries(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
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
			return diag.WrapError(err)
		}
		for _, item := range result.Schemas {
			s, err := svc.GetSchema(ctx, &glue.GetSchemaInput{SchemaId: &types.SchemaId{SchemaArn: item.SchemaArn}})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- s
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
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
			return diag.WrapError(err)
		}
		for _, item := range result.Schemas {
			s, err := svc.GetSchemaVersion(ctx, &glue.GetSchemaVersionInput{
				SchemaVersionId: item.SchemaVersionId,
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- s
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
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
			return diag.WrapError(err)
		}

		for k, v := range result.MetadataInfoMap {
			metadata[k] = v
		}

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, metadata))
}
