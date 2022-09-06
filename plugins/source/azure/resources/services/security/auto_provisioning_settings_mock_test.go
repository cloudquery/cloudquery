// Auto generated code - DO NOT EDIT.

package security

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func TestSecurityAutoProvisioningSettings(t *testing.T) {
	client.AzureMockTestHelper(t, AutoProvisioningSettings(), createAutoProvisioningSettingsMock, client.TestOptions{})
}

func createAutoProvisioningSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityAutoProvisioningSettingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			AutoProvisioningSettings: mockClient,
		},
	}

	data := security.AutoProvisioningSetting{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := security.NewAutoProvisioningSettingListPage(security.AutoProvisioningSettingList{Value: &[]security.AutoProvisioningSetting{data}}, func(ctx context.Context, result security.AutoProvisioningSettingList) (security.AutoProvisioningSettingList, error) {
		return security.AutoProvisioningSettingList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
