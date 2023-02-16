package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildTrustedAdvisorChecks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	checks := []types.TrustedAdvisorCheckDescription{}
	err := faker.FakeObject(&checks)
	if err != nil {
		t.Fatal(err)
	}

	for _, languageCode := range checksSupportedLanguageCodes {
		m.EXPECT().DescribeTrustedAdvisorChecks(gomock.Any(), &support.DescribeTrustedAdvisorChecksInput{Language: aws.String(languageCode)}).
			Return(&support.DescribeTrustedAdvisorChecksOutput{Checks: checks}, nil)
	}

	err = mockCheckSummaries(checks[0], m)
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
