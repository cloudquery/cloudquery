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
)

func buildSeverityLevels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	levels := []types.SeverityLevel{}
	err := faker.FakeObject(&levels)
	if err != nil {
		t.Fatal(err)
	}

	for _, languageCode := range severitySupportedLanguageCodes {
		m.EXPECT().DescribeSeverityLevels(gomock.Any(), &support.DescribeSeverityLevelsInput{Language: aws.String(languageCode)}, gomock.Any()).Return(&support.DescribeSeverityLevelsOutput{SeverityLevels: levels}, nil)
	}

	return client.Services{
		Support: m,
	}
}

func TestSeverityLevels(t *testing.T) {
	client.AwsMockTestHelper(t, SeverityLevels(), buildSeverityLevels, client.TestOptions{})
}
