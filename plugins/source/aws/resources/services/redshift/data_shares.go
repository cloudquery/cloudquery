package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DataShares() *schema.Table {
	tableName := "aws_redshift_data_shares"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_DataShare.html`,
		Resolver:    fetchDataShares,
		Transform:   transformers.TransformWithStruct(&types.DataShare{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchDataShares(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift

	config := redshift.DescribeDataSharesInput{
		MaxRecords: aws.Int32(100),
	}
	paginator := redshift.NewDescribeDataSharesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DataShares
	}

	return nil
}
