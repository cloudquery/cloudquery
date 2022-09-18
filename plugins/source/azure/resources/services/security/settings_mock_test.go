// Auto generated code - DO NOT EDIT.

package security

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func TestSecuritySettings(t *testing.T) {
	client.MockTestHelper(t, Settings(), createSettingsMock)
}

func createSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecuritySettingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Settings: mockClient,
		},
	}

	data := security.Setting{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := security.NewSettingsListPage(security.SettingsList{Value: &[]security.BasicSetting{data}}, func(ctx context.Context, result security.SettingsList) (security.SettingsList, error) {
		return security.SettingsList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
