package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDistributions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var d lightsail.GetDistributionsOutput
	if err := faker.FakeData(&d); err != nil {
		t.Fatal(err)
	}
	d.NextPageToken = nil
	mock.EXPECT().GetDistributions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&d, nil)

	var r lightsail.GetDistributionLatestCacheResetOutput
	if err := faker.FakeData(&r); err != nil {
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
