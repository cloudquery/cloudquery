package ses

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ActiveReceiptRuleSets() *schema.Table {
	tableName := "aws_ses_active_receipt_rule_sets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ses/latest/APIReference/API_DescribeActiveReceiptRuleSet.html`,
		Resolver:    fetchSesActiveReceiptRuleSets,
		Transform:   transformers.TransformWithStruct(&ses.DescribeActiveReceiptRuleSetOutput{}, transformers.WithSkipFields("Metadata", "ResultMetadata")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Metadata.Name"),
				PrimaryKey: true,
			},
			{
				Name:     "created_timestamp",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("Metadata.CreatedTimestamp"),
			},
		},
	}
}

// Supported regions based on https://docs.aws.amazon.com/ses/latest/dg/regions.html#region-receive-email
// We hard code as there isn't a good way to automatically fetch this list
var supportedRegions = []string{"us-east-1", "us-west-2", "eu-west-1"}

func isRegionSupported(region string) bool {
	for _, r := range supportedRegions {
		if r == region {
			return true
		}
	}
	return false
}

func fetchSesActiveReceiptRuleSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ses

	set, err := svc.DescribeActiveReceiptRuleSet(ctx, nil, func(o *ses.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if !isRegionSupported(cl.Region) && client.IgnoreWithInvalidAction(err) {
			return nil
		}
		return err
	}

	if set.Metadata != nil && set.Metadata.Name != nil {
		res <- set
	}

	return nil
}
