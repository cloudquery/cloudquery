package resources_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildNetworkWatchersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	watcherSvc := mocks.NewMockWatchersClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			Watchers: watcherSvc,
		},
	}
	watcher := network.Watcher{}
	err := faker.FakeData(&watcher)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	watcher.ID = &id
	page := network.WatcherListResult{Value: &[]network.Watcher{watcher}}
	watcherSvc.EXPECT().ListAll(gomock.Any()).Return(page, nil)

	flowLogInfo := network.FlowLogInformation{}
	err = faker.FakeData(&flowLogInfo)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	status := network.WatchersGetFlowLogStatusFuture{Result: func(client network.WatchersClient) (network.FlowLogInformation, error) {
		return flowLogInfo, nil
	}}
	watcherSvc.EXPECT().GetFlowLogStatus(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(status, nil)
	return s
}

func TestNetworkWatchers(t *testing.T) {
	azureTestHelper(t, resources.NetworkWatchers(), buildNetworkWatchersMock)
}
