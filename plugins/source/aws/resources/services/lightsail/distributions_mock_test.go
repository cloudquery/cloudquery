package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDistributions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var d lightsail.GetDistributionsOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	d.NextPageToken = nil
	mock.EXPECT().GetDistributions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&d, nil)

	var r lightsail.GetDistributionLatestCacheResetOutput
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
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
