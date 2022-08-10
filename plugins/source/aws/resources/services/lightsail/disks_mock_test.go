package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDisks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var disks lightsail.GetDisksOutput
	if err := faker.FakeData(&disks); err != nil {
		t.Fatal(err)
	}
	disks.NextPageToken = nil
	mock.EXPECT().GetDisks(
		gomock.Any(),
		&lightsail.GetDisksInput{},
		gomock.Any(),
	).Return(
		&disks,
		nil,
	)

	var diskSnapshots lightsail.GetDiskSnapshotsOutput
	if err := faker.FakeData(&diskSnapshots); err != nil {
		t.Fatal(err)
	}
	diskSnapshots.NextPageToken = nil
	mock.EXPECT().GetDiskSnapshots(
		gomock.Any(),
		&lightsail.GetDiskSnapshotsInput{},
		gomock.Any(),
	).Return(
		&diskSnapshots,
		nil,
	)

	return client.Services{Lightsail: mock}
}

func TestLightsailDisks(t *testing.T) {
	client.AwsMockTestHelper(t, Disks(), buildDisks, client.TestOptions{})
}
