//go:build !integration

package subscription

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSubscriptionSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSubscriptionGetter(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model subscription.Model
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	rg := client.FakeResourceGroup
	model.SubscriptionID = &rg
	m.EXPECT().Get(gomock.Any(), subscriptionID).Return(model, nil)

	return services.Services{
		Subscriptions: services.SubscriptionsClient{
			SubscriptionID: subscriptionID,
			Subscriptions:  m,
		},
	}
}

func TestSubscriptionSubscriptions(t *testing.T) {
	client.AzureMockTestHelper(t, SubscriptionSubscriptions(), buildSubscriptionSubscriptionsMock, client.TestOptions{})
}
