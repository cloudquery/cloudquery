package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SchemaArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueRegistrySchemaTags,
			},
		},

		Relations: []*schema.Table{
			registrySchemaVersions(),
		},
	}
}
