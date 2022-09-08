// Auto generated code - DO NOT EDIT.

package monitor

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
)

func TestMonitorLogProfiles(t *testing.T) {
	client.MockTestHelper(t, LogProfiles(), createLogProfilesMock)
}

func createLogProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorLogProfilesClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			LogProfiles: mockClient,
		},
	}

	data := insights.LogProfileResource{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := insights.LogProfileCollection{Value: &[]insights.LogProfileResource{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
