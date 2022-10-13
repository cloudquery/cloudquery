package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildBackupGlobalSettingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var settings backup.DescribeGlobalSettingsOutput
	if err := faker.FakeObject(&settings); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeGlobalSettings(
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

func TestGlobalSettings(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalSettings(), buildBackupGlobalSettingsMock, client.TestOptions{})
}
