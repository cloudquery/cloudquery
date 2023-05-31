package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func BlockPublicAccessConfigs() *schema.Table {
	tableName := "aws_emr_block_public_access_configs"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchEmrBlockPublicAccessConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform: transformers.TransformWithStruct(&emr.GetBlockPublicAccessConfigurationOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchEmrBlockPublicAccessConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Emr
	out, err := svc.GetBlockPublicAccessConfiguration(ctx, &emr.GetBlockPublicAccessConfigurationInput{}, func(options *emr.Options) {
		options.Region = c.Region
	})
	if err != nil {
		if client.IgnoreNotAvailableRegion(err) {
			return nil
		}
		return err
	}
	res <- out
	return nil
}
