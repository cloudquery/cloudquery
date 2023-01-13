package savingsplans

import (
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:        "aws_savingsplans_plans",
		Description: `https://docs.aws.amazon.com/savingsplans/latest/APIReference/API_SavingsPlan.html`,
		Resolver:    fetchSavingsPlans,
		Transform:   transformers.TransformWithStruct(&types.SavingsPlan{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
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
