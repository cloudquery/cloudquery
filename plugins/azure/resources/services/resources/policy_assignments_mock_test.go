package resources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildResourcesPolicyAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	a := mocks.NewMockAssignmentsClient(ctrl)

	faker.SetIgnoreInterface(true)
	as := policy.Assignment{}
	if err := faker.FakeData(&as); err != nil {
		t.Fatal(err)
	}
	as.Metadata = make(map[string]string)

	page := policy.NewAssignmentListResultPage(policy.AssignmentListResult{Value: &[]policy.Assignment{as}}, func(ctx context.Context, result policy.AssignmentListResult) (policy.AssignmentListResult, error) {
		return policy.AssignmentListResult{}, nil
	})
	a.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(page, nil)
	return services.Services{
		Resources: services.ResourcesClient{Assignments: a},
	}
}

func TestResourcesPolicyAssignments(t *testing.T) {
	client.AzureMockTestHelper(t, ResourcesPolicyAssignments(), buildResourcesPolicyAssignmentsMock, client.TestOptions{})
}
