package cloudwatch

import (
	"context"
	"errors"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/mitchellh/hashstructure/v2"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type metricOutput struct {
	types.Metric
	InputJSON tableoptions.CloudwatchListMetricsInput `json:"input_json"`
	InputHash string                                  `json:"input_hash"`

	getStatsInputs []tableoptions.CloudwatchGetMetricStatisticsInput
}

func Metrics() *schema.Table {
	tableName := "aws_alpha_cloudwatch_metrics"
	return &schema.Table{
		Name:  tableName,
		Title: `Cloudwatch Metrics (Alpha)`,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

Please note that this table is considered **alpha** (experimental) and may have breaking changes or be removed in the future.
`,
		Resolver:  fetchCloudwatchMetrics,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform: transformers.TransformWithStruct(&metricOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:        "input_hash",
				Description: `The hash of the input used to generate this result.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("InputHash"),
				PrimaryKey:  true,
			},
			{
				Name:        "input_json",
				Description: `The JSON of the input used to generate this result.`,
				Type:        cqtypes.ExtensionTypes.JSON,
				Resolver:    schema.PathResolver("InputJSON"),
			}},
		Relations: []*schema.Table{
			metricStatistics(),
		},
	}
}

func fetchCloudwatchMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	if len(cl.Spec.TableOptions.CloudwatchMetrics) > 0 && !cl.Spec.UsePaidAPIs {
		return client.ErrPaidAPIsNotEnabled
	}

	if len(cl.Spec.TableOptions.CloudwatchMetrics) == 0 {
		return errors.New("skipping `aws_alpha_cloudwatch_metrics` because `list_metrics` is not specified in `table_options`")
	}

	svc := cl.Services().Cloudwatch
	for _, input := range cl.Spec.TableOptions.CloudwatchMetrics {
		input := input

		hash, err := hashstructure.Hash(input, hashstructure.FormatV2, nil)
		if err != nil {
			return err
		}

		paginator := cloudwatch.NewListMetricsPaginator(svc, &input.ListMetricsOpts.ListMetricsInput)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *cloudwatch.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for i := range page.Metrics {
				res <- metricOutput{
					Metric:         page.Metrics[i],
					InputJSON:      input.ListMetricsOpts,
					InputHash:      strconv.FormatUint(hash, 10),
					getStatsInputs: input.GetMetricStatisticsOpts,
				}
			}
		}
	}
	return nil
}
