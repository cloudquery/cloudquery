package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildBackupRegionSettingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var settings backup.DescribeRegionSettingsOutput
	if err := faker.FakeData(&settings); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeRegionSettings(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&settings,
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestRegionSettings(t *testing.T) {
	client.AwsMockTestHelper(t, RegionSettings(), buildBackupRegionSettingsMock, client.TestOptions{})
}
