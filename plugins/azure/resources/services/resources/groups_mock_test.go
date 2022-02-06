package resources

import (
	"context"
	"testing"

	resourcegroup "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildResourceGroupMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockGroupsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{Groups: m},
	}
	l := resourcegroup.Group{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	groupPage := resourcegroup.NewGroupListResultPage(resourcegroup.GroupListResult{Value: &[]resourcegroup.Group{l}}, func(ctx context.Context, result resourcegroup.GroupListResult) (resourcegroup.GroupListResult, error) {
		return resourcegroup.GroupListResult{}, nil
	})
	m.EXPECT().List(gomock.Any(), "", nil).Return(groupPage, nil)
	return s
}

func TestResourceGroups(t *testing.T) {
	client.AzureMockTestHelper(t, ResourcesGroups(), buildResourceGroupMock, client.TestOptions{})
}
