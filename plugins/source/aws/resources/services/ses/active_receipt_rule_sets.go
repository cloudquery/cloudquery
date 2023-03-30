package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ActiveReceiptRuleSets() *schema.Table {
	tableName := "aws_ses_active_receipt_rule_sets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html`,
		Resolver:    fetchSesActiveReceiptRuleSets,
		Transform:   client.TransformWithStruct(&ses.DescribeActiveReceiptRuleSetOutput{}, transformers.WithSkipFields("Metadata", "ResultMetadata")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "email"),
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
