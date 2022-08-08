package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildInstanceSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var is lightsail.GetInstanceSnapshotsOutput
	if err := faker.FakeData(&is); err != nil {
		t.Fatal(err)
	}
	is.NextPageToken = nil

	mock.EXPECT().GetInstanceSnapshots(
		gomock.Any(),
		&lightsail.GetInstanceSnapshotsInput{},
		gomock.Any(),
	).Return(&is, nil)

	return client.Services{Lightsail: mock}
}

func TestLightsailInstanceSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceSnapshots(), buildInstanceSnapshots, client.TestOptions{})
}
