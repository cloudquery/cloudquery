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

func createRuleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRuleSetsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			RuleSets: mockClient,
		},
	}

	data := cdn.RuleSet{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := cdn.NewRuleSetListResultPage(cdn.RuleSetListResult{Value: &[]cdn.RuleSet{data}}, func(ctx context.Context, result cdn.RuleSetListResult) (cdn.RuleSetListResult, error) {
		return cdn.RuleSetListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
