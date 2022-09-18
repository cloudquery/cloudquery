// Auto generated code - DO NOT EDIT.

package eventhub

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
)

func createNetworkRuleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockEventHubNetworkRuleSetsClient(ctrl)
	s := services.Services{
		EventHub: services.EventHubClient{
			NetworkRuleSets: mockClient,
		},
	}

	data := eventhub.NetworkRuleSet{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().GetNetworkRuleSet(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
