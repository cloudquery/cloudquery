package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

var checksSupportedLanguageCodes = []string{"en", "fr", "de", "id", "it", "ja", "ko", "pt_BR", "es", "zh", "zh_TW"}

func TrustedAdvisorChecks() *schema.Table {
	tableName := "aws_support_trusted_advisor_checks"
	return &schema.Table{
		Name:        "aws_support_trusted_advisor_checks",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorChecks.html`,
		Resolver:    fetchTrustedAdvisorChecks,
		Transform:   transformers.TransformWithStruct(&types.TrustedAdvisorCheckDescription{}, transformers.WithPrimaryKeys("Id")),
		Multiplex:   client.ServiceAccountRegionsLanguageCodeMultiplex(tableName, "support", checksSupportedLanguageCodes),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
		Relations: []*schema.Table{trustedAdvisorCheckSummaries(), trustedAdvisorCheckResults()},
	}
}

func fetchTrustedAdvisorChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Support
	input := support.DescribeTrustedAdvisorChecksInput{Language: aws.String(cl.LanguageCode)}

	response, err := svc.DescribeTrustedAdvisorChecks(ctx, &input, func(o *support.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.Checks

	return nil
}
