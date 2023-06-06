package costexplorer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/mitchellh/hashstructure/v2"
)

func CustomCost() *schema.Table {
	tableName := "aws_costexplorer_cost_custom"
	return &schema.Table{
		Name:     tableName,
		Resolver: fetchCustom,
		Title:    "AWS Cost Explorer costs based on custom inputs",
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration as well as specify the request parameters in the 'table_options' attribute. `,
		Transform: transformers.TransformWithStruct(&wrappedResultByTime{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "start_date",
				Description: `The start date covered by the forecast.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("TimePeriod.Start"),
				PrimaryKey:  true,
			},
			{
				Name:        "end_date",
				Description: `The end date covered by the forecast.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("TimePeriod.End"),
				PrimaryKey:  true,
			},
			{
				Name:        "input_hash",
				Description: `The hash of the input used to generate this result.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("inputHash"),
				PrimaryKey:  true,
			},
			{
				Name:        "input_json",
				Description: `The JSON of the input used to generate this result.`,
				Type:        cqtypes.ExtensionTypes.JSON,
				Resolver:    schema.PathResolver("inputJSON"),
			},
		},
	}
}

type wrappedResultByTime struct {
	types.ResultByTime
	inputJSON string
	inputHash string
}

func fetchCustom(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	if !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_costexplorer_cost_current_month` because `use_paid_apis` is set to false")
		return nil
	}

	if cl.Spec.TableOptions.CustomCostExplorer == nil {
		cl.Logger().Info().Msg("skipping `aws_costexplorer_cost_custom` because no inputs are set")
		return nil
	}

	svc := cl.Services().Costexplorer
	allConfigs := cl.Spec.TableOptions.CustomCostExplorer.GetCostAndUsageOpts
	for _, input := range allConfigs {
		hash, err := hashstructure.Hash(input, hashstructure.FormatV2, nil)
		if err != nil {
			return err
		}

		jsonInput, err := json.Marshal(input)
		if err != nil {
			return err
		}

		for {
			resp, err := svc.GetCostAndUsage(ctx, &input.GetCostAndUsageInput, func(options *costexplorer.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for _, r := range resp.ResultsByTime {
				res <- wrappedResultByTime{
					ResultByTime: r,
					inputHash:    fmt.Sprintf("%d", hash),
					inputJSON:    string(jsonInput),
				}
			}
			if aws.ToString(resp.NextPageToken) == "" {
				break
			}
		}
	}
	return nil
}
