package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTrustedAdvisorChecks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	checks := []types.TrustedAdvisorCheckDescription{}
	require.NoError(t, faker.FakeObject(&checks))

	for _, languageCode := range checksSupportedLanguageCodes {
		m.EXPECT().DescribeTrustedAdvisorChecks(gomock.Any(), &support.DescribeTrustedAdvisorChecksInput{Language: aws.String(languageCode)}, gomock.Any()).
			Return(&support.DescribeTrustedAdvisorChecksOutput{Checks: checks}, nil)
	}

	err := mockCheckSummaries(checks[0], m)
	if err != nil {
		t.Fatal(err)
	}

	err = mockCheckResults(checks[0], m)
	if err != nil {
		t.Fatal(err)
	}

	return client.Services{
		Support: m,
	}
}

func TestTrustedAdvisorChecks(t *testing.T) {
	client.AwsMockTestHelper(t, TrustedAdvisorChecks(), buildTrustedAdvisorChecks, client.TestOptions{})
}
