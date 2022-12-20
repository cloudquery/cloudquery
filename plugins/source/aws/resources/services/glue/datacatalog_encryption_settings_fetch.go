package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueDatacatalogEncryptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- result.DataCatalogEncryptionSettings
	return nil
}
