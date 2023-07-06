package frauddetector

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func rules() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_rules",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html`,
		Resolver:    fetchFrauddetectorRules,
		Transform:   transformers.TransformWithStruct(&types.RuleDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchFrauddetectorRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := frauddetector.NewGetRulesPaginator(meta.(*client.Client).Services().Frauddetector,
		&frauddetector.GetRulesInput{DetectorId: parent.Item.(types.Detector).DetectorId})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.RuleDetails
	}
	return nil
}
