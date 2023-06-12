package cloudwatch

import (
	"context"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	cqtypes "github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/mitchellh/hashstructure/v2"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func MetricStatisticsCustom() *schema.Table {
	tableName := "aws_cloudwatch_metric_statistics_custom"
	return &schema.Table{
		Name:  tableName,
		Title: "Cloudwatch Metric Statistics custom query",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.
`,
		Resolver:  fetchCloudwatchMetricStatisticsCustom,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform: transformers.TransformWithStruct(&statOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
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
			},
		},
	}
}

func fetchCloudwatchMetricStatisticsCustom(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var allConfigs []tableoptions.CloudwatchGetMetricStatisticsInput
	if cl.Spec.TableOptions.CloudwatchCustomMetricStats != nil {
		allConfigs = cl.Spec.TableOptions.CloudwatchCustomMetricStats.GetMetricStatisticsOpts
	}

	if len(allConfigs) > 0 && !cl.Spec.UsePaidAPIs {
		return client.ErrPaidAPIsNotEnabled
	}

	svc := cl.Services().Cloudwatch
	for _, input := range allConfigs {
		input := input
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
		res <- statOutput{
			GetMetricStatisticsOutput: data,
			InputJSON:                 input,
			InputHash:                 strconv.FormatUint(hash, 10),
		}
	}
	return nil
}
