package securityhub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildEnabledStandards(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	standardsSubscription := types.StandardsSubscription{}
	err := faker.FakeObject(&standardsSubscription)
	if err != nil {
		t.Fatal(err)
	}

	shMock.EXPECT().GetEnabledStandards(
		gomock.Any(),
		&securityhub.GetEnabledStandardsInput{MaxResults: 100},
		gomock.Any(),
	).Return(
		&securityhub.GetEnabledStandardsOutput{StandardsSubscriptions: []types.StandardsSubscription{standardsSubscription}},
		nil,
	)

	return client.Services{Securityhub: shMock}
}

func TestEnabledStandards(t *testing.T) {
	client.AwsMockTestHelper(t, EnabledStandards(), buildEnabledStandards, client.TestOptions{})
}
