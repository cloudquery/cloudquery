// Auto generated code - DO NOT EDIT.

package network

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkWatchers(t *testing.T) {
	client.MockTestHelper(t, Watchers(), createWatchersMock)
}

func createWatchersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkWatchersClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			Watchers: mockClient,
		},
	}

	data := network.Watcher{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := network.WatcherListResult{Value: &[]network.Watcher{data}}

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
