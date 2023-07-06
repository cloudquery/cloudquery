package savingsplans

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Plans() *schema.Table {
	tableName := "aws_savingsplans_plans"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/savingsplans/latest/APIReference/API_SavingsPlan.html`,
		Resolver:    fetchSavingsPlans,
		Transform:   transformers.TransformWithStruct(&types.SavingsPlan{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "savingsplans"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:        "arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("SavingsPlanArn"),
				Description: `The Amazon Resource Name (ARN) of the Savings Plan.`,
				PrimaryKey:  true,
			},
		},
	}
}

func fetchSavingsPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Savingsplans
	config := savingsplans.DescribeSavingsPlansInput{
		MaxResults: aws.Int32(1000),
	}
	// no paginator available
	for {
		response, err := svc.DescribeSavingsPlans(ctx, &config, func(o *savingsplans.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.SavingsPlans
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
