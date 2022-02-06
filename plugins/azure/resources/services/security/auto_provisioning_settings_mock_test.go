package security

import (
	"context"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecurityAutoProvisioningSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSecurityAutoProvisioningSettingsClient(ctrl)

	var setting security.AutoProvisioningSetting
	if err := faker.FakeData(&setting); err != nil {
		t.Fatal(err)
	}

	result := security.NewAutoProvisioningSettingListPage(
		security.AutoProvisioningSettingList{Value: &[]security.AutoProvisioningSetting{setting}},
		func(context.Context, security.AutoProvisioningSettingList) (security.AutoProvisioningSettingList, error) {
			return security.AutoProvisioningSettingList{}, nil
		},
	)
	m.EXPECT().List(gomock.Any()).Return(result, nil)
	return services.Services{
		Security: services.SecurityClient{AutoProvisioningSettings: m},
	}
}

func TestSecurityAutoProvisioningSettings(t *testing.T) {
	client.AzureMockTestHelper(t, SecurityAutoProvisioningSettings(), buildSecurityAutoProvisioningSettingsMock, client.TestOptions{})
}
