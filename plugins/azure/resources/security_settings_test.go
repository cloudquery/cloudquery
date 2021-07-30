package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecuritySettings(s security.BasicSetting) func(*testing.T, *gomock.Controller) services.Services {
	return func(t *testing.T, ctrl *gomock.Controller) services.Services {
		m := mocks.NewMockSecuritySettingsClient(ctrl)

		result := security.NewSettingsListPage(
			security.SettingsList{Value: &[]security.BasicSetting{s}},
			func(context.Context, security.SettingsList) (security.SettingsList, error) {
				return security.SettingsList{}, nil
			},
		)
		m.EXPECT().List(gomock.Any()).Return(result, nil)
		return services.Services{
			Security: services.SecurityClient{Settings: m},
		}
	}
}

func TestSecuritySettings(t *testing.T) {
	t.Run("data_export_settings", func(t *testing.T) {
		var setting security.DataExportSettings
		if err := faker.FakeData(&setting); err != nil {
			t.Fatal(err)
		}
		id := fakeResourceGroup + "/" + *setting.ID
		setting.ID = &id
		azureTestHelper(t, resources.SecuritySettings(), buildSecuritySettings(setting))
	})
}
