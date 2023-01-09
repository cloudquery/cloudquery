package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegistrySchemaVersions() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_registry_schema_versions",
		Resolver:            fetchGlueRegistrySchemaVersions,
		PreResourceResolver: getRegistrySchemaVersion,
		Transform:           transformers.TransformWithStruct(&glue.GetSchemaVersionOutput{}),
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
				Name:     "registry_schema_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueRegistrySchemaVersionMetadata,
			},
		},
	}
}
