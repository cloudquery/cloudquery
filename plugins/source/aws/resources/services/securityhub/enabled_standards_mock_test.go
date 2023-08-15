package securityhub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEnabledStandards(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	standardsSubscription := types.StandardsSubscription{}
	require.NoError(t, faker.FakeObject(&standardsSubscription))

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
