package savingsplans

import (
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
		},
	}
}
