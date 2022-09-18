// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func createRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRulesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Rules: mockClient,
		},
	}

	data := cdn.Rule{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewRuleListResultPage(cdn.RuleListResult{Value: &[]cdn.Rule{data}}, func(ctx context.Context, result cdn.RuleListResult) (cdn.RuleListResult, error) {
		return cdn.RuleListResult{}, nil
	})

	data.Actions = &[]cdn.BasicDeliveryRuleAction{}
	data.Conditions = &[]cdn.BasicDeliveryRuleCondition{}
	mockClient.EXPECT().ListByRuleSet(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
