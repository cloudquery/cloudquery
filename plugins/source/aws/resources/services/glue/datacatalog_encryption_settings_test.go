package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDatacatalogEncryptionSettingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var s glue.GetDataCatalogEncryptionSettingsOutput
	require.NoError(t, faker.FakeObject(&s))
	m.EXPECT().GetDataCatalogEncryptionSettings(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&s, nil)

	return client.Services{
		Glue: m,
	}
}

func TestDatacatalogEncryptionSettings(t *testing.T) {
	client.AwsMockTestHelper(t, DatacatalogEncryptionSettings(), buildDatacatalogEncryptionSettingsMock, client.TestOptions{})
}
