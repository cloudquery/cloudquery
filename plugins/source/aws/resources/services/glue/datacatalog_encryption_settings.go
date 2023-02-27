package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatacatalogEncryptionSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_datacatalog_encryption_settings",
		Description: "https://docs.aws.amazon.com/glue/latest/webapi/API_GetDataCatalogEncryptionSettings.html",
		Resolver:    fetchGlueDatacatalogEncryptionSettings,
		Transform:   transformers.TransformWithStruct(&types.DataCatalogEncryptionSettings{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
