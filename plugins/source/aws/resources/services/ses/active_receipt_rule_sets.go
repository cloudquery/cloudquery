package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ActiveReceiptRuleSets() *schema.Table {
	return &schema.Table{
		Name:        "aws_ses_active_receipt_rule_sets",
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html`,
		Resolver:    fetchSesActiveReceiptRuleSets,
		Transform:   transformers.TransformWithStruct(&ses.DescribeActiveReceiptRuleSetOutput{}, transformers.WithSkipFields("Metadata", "ResultMetadata")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Metadata.Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Metadata.CreatedTimestamp"),
			},
		},
	}
}
