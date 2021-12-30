//go:build !integration

package ad

import (
	"context"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildADGroups(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockADGroupsClient(ctrl)
	var group graphrbac.ADGroup
	faker.SetIgnoreInterface(true)
	defer faker.SetIgnoreInterface(false)
	if err := faker.FakeData(&group); err != nil {
		t.Fatal(err)
	}
	group.AdditionalProperties = map[string]interface{}{"test": "value"}

	groupListPage := graphrbac.NewGroupListResultPage(
		graphrbac.GroupListResult{Value: &[]graphrbac.ADGroup{group}},
		func(ctx context.Context, list graphrbac.GroupListResult) (graphrbac.GroupListResult, error) {
			return graphrbac.GroupListResult{}, nil
		},
	)
	m.EXPECT().List(gomock.Any(), "").Return(groupListPage, nil)
	return services.Services{
		AD: services.AD{Groups: m},
	}
}

func TestADGroups(t *testing.T) {
	client.AzureMockTestHelper(t, Groups(), buildADGroups, client.TestOptions{})
}
