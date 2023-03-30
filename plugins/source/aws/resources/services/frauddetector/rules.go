package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_rules",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html`,
		Resolver:    fetchFrauddetectorRules,
		Transform:   transformers.TransformWithStruct(&types.RuleDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
