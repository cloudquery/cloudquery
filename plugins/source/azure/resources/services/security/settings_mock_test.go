package security

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/security"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/security"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSettings(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockSettingsClient := mocks.NewMockSettingsClient(ctrl)

	var response api.SettingsClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	var setting api.DataExportSettings
	require.NoError(t, faker.FakeObject(&setting))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	setting.ID = to.Ptr(id)
	response.Value = []api.SettingClassification{&setting}

	mockSettingsClient.EXPECT().NewListPager(gomock.Any()).
		Return(runtime.NewPager(runtime.PagingHandler[api.SettingsClientListResponse]{
			More: func(api.SettingsClientListResponse) bool {
				// it'll be called ONLY after the 1st page was processed
				return false
			},
			Fetcher: func(context.Context, *api.SettingsClientListResponse) (api.SettingsClientListResponse, error) {
				return response, nil
			},
		})).MinTimes(1)

	securityClient := &service.SecurityClient{
		SettingsClient: mockSettingsClient,
	}

	c := &client.Services{Security: securityClient}

	return c
}

func TestSettings(t *testing.T) {
	client.MockTestHelper(t, Settings(), buildSettings)
}
