package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func CustomCost() *schema.Table {
	tableName := "aws_costexplorer_cost_custom"
	return &schema.Table{
		Name:     tableName,
		Resolver: fetchCustom,
		Title:    "CUSTOM",
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. `,
		Transform: transformers.TransformWithStruct(&types.ResultByTime{}),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "start_date",
				Description: `The start date covered by the forecast.`,
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePeriod.Start"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "end_date",
				Description: `The end date covered by the forecast.`,
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePeriod.End"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
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
		for {
			resp, err := svc.GetCostAndUsage(ctx, &input.GetCostAndUsageInput, func(options *costexplorer.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- resp.ResultsByTime
			if aws.ToString(resp.NextPageToken) == "" {
				break
			}
		}
	}
	return nil
}
