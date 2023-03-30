package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegistrySchemaVersions() *schema.Table {
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
