package cloudwatch

import (
	"context"
	"errors"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/mitchellh/hashstructure/v2"
)

type statOutput struct {
	types.Datapoint
	Label     *string
	InputJSON tableoptions.CloudwatchGetMetricStatisticsInput `json:"input_json"`
	InputHash string                                          `json:"input_hash"`
}

func metricStatistics() *schema.Table {
	tableName := "aws_alpha_cloudwatch_metric_statistics"
	return &schema.Table{
		Name:  tableName,
		Title: `Cloudwatch Metric Statistics (Alpha)`,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

Please note that this table is considered **alpha** (experimental) and may have breaking changes or be removed in the future.
`,
		Resolver:  fetchCloudwatchMetricStatistics,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform: transformers.TransformWithStruct(&statOutput{},
			transformers.WithPrimaryKeys("Timestamp", "Label"),
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:        "parent_input_hash",
				Description: `The hash of the parent input used to generate this result.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.ParentColumnResolver("input_hash"),
				PrimaryKey:  true,
			},
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
			},
		},
	}
}

func fetchCloudwatchMetricStatistics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	item := parent.Item.(metricOutput)

	if len(item.getStatsInputs) == 0 {
		return errors.New("skipping `aws_alpha_cloudwatch_metric_statistics` because `get_metric_statistics` is not specified in `table_options`")
	}

	svc := cl.Services().Cloudwatch
	for _, input := range item.getStatsInputs {
		input := input

		input.Dimensions = item.Dimensions
		input.Namespace = item.Namespace
		input.MetricName = item.MetricName

		hash, err := hashstructure.Hash(input, hashstructure.FormatV2, nil)
		if err != nil {
			return err
		}

		data, err := svc.GetMetricStatistics(ctx, &input.GetMetricStatisticsInput, func(options *cloudwatch.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for i := range data.Datapoints {
			res <- statOutput{
				Datapoint: data.Datapoints[i],
				Label:     data.Label,
				InputJSON: input,
				InputHash: strconv.FormatUint(hash, 10),
			}
		}
	}
	return nil
}
