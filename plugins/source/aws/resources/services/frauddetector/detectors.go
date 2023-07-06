package frauddetector

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Detectors() *schema.Table {
	tableName := "aws_frauddetector_detectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_Detector.html`,
		Resolver:    fetchFrauddetectorDetectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "frauddetector"),
		Transform:   transformers.TransformWithStruct(&types.Detector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveResourceTags,
			},
		},

		Relations: []*schema.Table{
			rules(),
		},
	}
}

func fetchFrauddetectorDetectors(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := frauddetector.NewGetDetectorsPaginator(meta.(*client.Client).Services().Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Detectors
	}
	return nil
}
