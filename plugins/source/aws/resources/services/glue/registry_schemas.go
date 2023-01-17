package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegistrySchemas() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_registry_schemas",
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_GetSchema.html`,
		Resolver:            fetchGlueRegistrySchemas,
		PreResourceResolver: getRegistrySchema,
		Transform:           transformers.TransformWithStruct(&glue.GetSchemaOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SchemaArn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueRegistrySchemaTags,
			},
		},

		Relations: []*schema.Table{
			RegistrySchemaVersions(),
		},
	}
}
