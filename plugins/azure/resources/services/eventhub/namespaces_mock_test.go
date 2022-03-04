package eventhub

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEventHubNamespacesServices(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockEventHubClient(ctrl)
	var namespace eventhub.EHNamespace
	if err := faker.FakeData(&namespace); err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup + "/" + *namespace.ID
	namespace.ID = &id
	m.EXPECT().List(gomock.Any()).Return(
		eventhub.NewEHNamespaceListResultPage(
			eventhub.EHNamespaceListResult{Value: &[]eventhub.EHNamespace{namespace}},
			func(c context.Context, lr eventhub.EHNamespaceListResult) (eventhub.EHNamespaceListResult, error) {
				return eventhub.EHNamespaceListResult{}, nil
			},
		),
		nil,
	)

	var rs eventhub.NetworkRuleSet
	if err := faker.FakeData(&rs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetNetworkRuleSet(gomock.Any(), "test", *namespace.Name).Return(rs, nil)
	return services.Services{EventHub: m}
}

func TestEventHubNamespacesServices(t *testing.T) {
	table := EventHubNamespaces()
	table.IgnoreInTests = false
	client.AzureMockTestHelper(t, table, buildEventHubNamespacesServices, client.TestOptions{})
}
