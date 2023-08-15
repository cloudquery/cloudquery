package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DatacatalogEncryptionSettings() *schema.Table {
	tableName := "aws_glue_datacatalog_encryption_settings"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/glue/latest/webapi/API_GetDataCatalogEncryptionSettings.html",
		Resolver:    fetchGlueDatacatalogEncryptionSettings,
		Transform:   transformers.TransformWithStruct(&types.DataCatalogEncryptionSettings{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchGlueDatacatalogEncryptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- result.DataCatalogEncryptionSettings
	return nil
}
