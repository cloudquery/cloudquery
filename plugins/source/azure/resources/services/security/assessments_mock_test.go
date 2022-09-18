// Auto generated code - DO NOT EDIT.

package security

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func TestSecurityAssessments(t *testing.T) {
	client.MockTestHelper(t, Assessments(), createAssessmentsMock)
}

func createAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityAssessmentsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Assessments: mockClient,
		},
	}

	data := security.Assessment{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewAssessmentListPage(security.AssessmentList{Value: &[]security.Assessment{data}}, func(ctx context.Context, result security.AssessmentList) (security.AssessmentList, error) {
		return security.AssessmentList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/test_sub").Return(result, nil)
	return s
}
