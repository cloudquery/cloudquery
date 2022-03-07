package account

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAccountLocationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	v := mocks.NewMockSubscriptionsClient(ctrl)
	s := services.Services{
		Subscriptions: services.Subscriptions{
			SubscriptionID: client.TestSubscriptionID,
			Subscriptions:  v,
		},
	}
	location := subscription.Location{}
	err := faker.FakeData(&location)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	id := client.FakeResourceGroup + "/" + *location.ID
	location.ID = &id

	v.EXPECT().ListLocations(gomock.Any(), gomock.Any()).Return(subscription.LocationListResult{
		Value: &[]subscription.Location{location},
	}, nil)
	return s
}

func TestComputeAccountLocations(t *testing.T) {
	client.AzureMockTestHelper(t, AccountLocations(), buildAccountLocationsMock, client.TestOptions{})
}
