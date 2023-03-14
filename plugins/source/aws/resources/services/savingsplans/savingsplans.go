package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SavingsPlanArn"),
				Description: `The Amazon Resource Name (ARN) of the Savings Plan.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSavingsPlanTags,
			},
		},
	}
}

func fetchSavingsPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Savingsplans
	config := savingsplans.DescribeSavingsPlansInput{
		MaxResults: aws.Int32(1000),
	}
	for {
		response, err := svc.DescribeSavingsPlans(ctx, &config)
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

func resolveSavingsPlanTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	sp := resource.Item.(types.SavingsPlan)
	cl := meta.(*client.Client)
	svc := cl.Services().Savingsplans
	out, err := svc.ListTagsForResource(ctx, &savingsplans.ListTagsForResourceInput{ResourceArn: sp.SavingsPlanArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
