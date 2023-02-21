package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var checksSupportedLanguageCodes = []string{"en", "fr", "de", "id", "it", "ja", "ko", "pt_BR", "es", "zh", "zh_TW"}

func TrustedAdvisorChecks() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_trusted_advisor_checks",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorChecks.html`,
		Resolver:    fetchTrustedAdvisorChecks,
		Transform:   transformers.TransformWithStruct(&types.TrustedAdvisorCheckDescription{}, transformers.WithPrimaryKeys("Id")),
		Multiplex:   client.ServiceAccountRegionsLanguageCodeMultiplex("support", checksSupportedLanguageCodes),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
		Relations: []*schema.Table{trustedAdvisorCheckSummaries(), trustedAdvisorCheckResults()},
	}
}

func fetchTrustedAdvisorChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Support
	input := support.DescribeTrustedAdvisorChecksInput{Language: aws.String(c.LanguageCode)}

	response, err := svc.DescribeTrustedAdvisorChecks(ctx, &input)
	if err != nil {
		return err
	}

	res <- response.Checks

	return nil
}
