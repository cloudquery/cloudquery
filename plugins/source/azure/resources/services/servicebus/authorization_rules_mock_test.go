// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func createAuthorizationRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusAuthorizationRulesClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			AuthorizationRules: mockClient,
			AccessKeys:         createAccessKeysMock(t, ctrl).Servicebus.AccessKeys,
		},
	}

	data := servicebus.SBAuthorizationRule{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBAuthorizationRuleListResultPage(servicebus.SBAuthorizationRuleListResult{Value: &[]servicebus.SBAuthorizationRule{data}}, func(ctx context.Context, result servicebus.SBAuthorizationRuleListResult) (servicebus.SBAuthorizationRuleListResult, error) {
		return servicebus.SBAuthorizationRuleListResult{}, nil
	})

	mockClient.EXPECT().ListAuthorizationRules(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
