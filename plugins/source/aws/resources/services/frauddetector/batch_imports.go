package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func BatchImports() *schema.Table {
	tableName := "aws_frauddetector_batch_imports"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html`,
		Resolver:    fetchFrauddetectorBatchImports,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "frauddetector"),
		Transform:   transformers.TransformWithStruct(&types.BatchImport{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchFrauddetectorBatchImports(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	paginator := frauddetector.NewGetBatchImportJobsPaginator(meta.(*client.Client).Services().Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.BatchImports
	}
	return nil
}
