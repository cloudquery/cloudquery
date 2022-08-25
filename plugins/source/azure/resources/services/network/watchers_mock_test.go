package network

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildNetworkWatchersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	watcherSvc := mocks.NewMockNetworkWatchersClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			Watchers: watcherSvc,
		},
	}
	watcher := network.Watcher{}
	err := faker.FakeData(&watcher)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	id := client.FakeResourceGroup + "/" + *watcher.ID
	watcher.ID = &id
	page := network.WatcherListResult{Value: &[]network.Watcher{watcher}}
	watcherSvc.EXPECT().ListAll(gomock.Any()).Return(page, nil)

	return s
}

func TestNetworkWatchers(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkWatchers(), buildNetworkWatchersMock, client.TestOptions{})
}
