// Auto generated code - DO NOT EDIT.

package logic

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"
)

func createDiagnosticSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockLogicDiagnosticSettingsClient(ctrl)
	s := services.Services{
		Logic: services.LogicClient{
			DiagnosticSettings: mockClient,
		},
	}

	data := insights.DiagnosticSettingsResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{data}}

	mockClient.EXPECT().List(gomock.Any(), "test").Return(result, nil)
	return s
}
