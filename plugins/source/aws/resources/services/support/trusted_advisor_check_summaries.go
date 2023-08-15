package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/golang/mock/gomock"
)

func trustedAdvisorCheckSummaries() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_trusted_advisor_check_summaries",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckSummaries.html`,
		Resolver:    fetchTrustedAdvisorCheckSummaries,
		Transform:   transformers.TransformWithStruct(&types.TrustedAdvisorCheckSummary{}, transformers.WithPrimaryKeys("CheckId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
	}
}

func fetchTrustedAdvisorCheckSummaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	// No need to get the summary for each language, as those are the same have the same check id
	if cl.LanguageCode != "en" {
		return nil
	}
	svc := cl.Services().Support
	check := parent.Item.(types.TrustedAdvisorCheckDescription)
	input := support.DescribeTrustedAdvisorCheckSummariesInput{CheckIds: []*string{check.Id}}

	response, err := svc.DescribeTrustedAdvisorCheckSummaries(ctx, &input, func(o *support.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.Summaries

	return nil
}

func mockCheckSummaries(check types.TrustedAdvisorCheckDescription, m *mocks.MockSupportClient) error {
	summaries := []types.TrustedAdvisorCheckSummary{}
	err := faker.FakeObject(&summaries)
	if err != nil {
		return err
	}

	input := support.DescribeTrustedAdvisorCheckSummariesInput{CheckIds: []*string{check.Id}}
	m.EXPECT().DescribeTrustedAdvisorCheckSummaries(gomock.Any(), &input, gomock.Any()).Return(&support.DescribeTrustedAdvisorCheckSummariesOutput{Summaries: summaries}, nil)
	return nil
}
