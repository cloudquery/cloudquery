package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDistributions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var d lightsail.GetDistributionsOutput
	require.NoError(t, faker.FakeObject(&d))

	d.NextPageToken = nil
	mock.EXPECT().GetDistributions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&d, nil)

	var r lightsail.GetDistributionLatestCacheResetOutput
	require.NoError(t, faker.FakeObject(&r))

	mock.EXPECT().GetDistributionLatestCacheReset(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&r, nil)

	return client.Services{Lightsail: mock}
}

func TestLightsailDistributions(t *testing.T) {
	client.AwsMockTestHelper(t, Distributions(), buildDistributions, client.TestOptions{})
}
