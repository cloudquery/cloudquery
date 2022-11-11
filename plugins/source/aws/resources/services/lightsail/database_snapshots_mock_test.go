package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDatabaseSnapshotsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	s := lightsail.GetRelationalDatabaseSnapshotsOutput{}
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	s.NextPageToken = nil
	m.EXPECT().GetRelationalDatabaseSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestDatabaseSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, DatabaseSnapshots(), buildDatabaseSnapshotsMock, client.TestOptions{})
}
