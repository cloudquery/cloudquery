// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
)

func TestResourcesPolicyAssignments(t *testing.T) {
	client.MockTestHelper(t, PolicyAssignments(), createPolicyAssignmentsMock)
}

func createPolicyAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesPolicyAssignmentsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			PolicyAssignments: mockClient,
		},
	}

	data := policy.Assignment{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := policy.NewAssignmentListResultPage(policy.AssignmentListResult{Value: &[]policy.Assignment{data}}, func(ctx context.Context, result policy.AssignmentListResult) (policy.AssignmentListResult, error) {
		return policy.AssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), gomock.Any(), "", nil).Return(result, nil)
	return s
}
