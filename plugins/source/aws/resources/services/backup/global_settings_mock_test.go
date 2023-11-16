package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBackupGlobalSettingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var settings backup.DescribeGlobalSettingsOutput
	require.NoError(t, faker.FakeObject(&settings))

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
