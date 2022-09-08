// Auto generated code - DO NOT EDIT.

package security

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

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func TestSecurityJitNetworkAccessPolicies(t *testing.T) {
	client.MockTestHelper(t, JitNetworkAccessPolicies(), createJitNetworkAccessPoliciesMock)
}

func createJitNetworkAccessPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityJitNetworkAccessPoliciesClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			JitNetworkAccessPolicies: mockClient,
		},
	}

	data := security.JitNetworkAccessPolicy{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := security.NewJitNetworkAccessPoliciesListPage(security.JitNetworkAccessPoliciesList{Value: &[]security.JitNetworkAccessPolicy{data}}, func(ctx context.Context, result security.JitNetworkAccessPoliciesList) (security.JitNetworkAccessPoliciesList, error) {
		return security.JitNetworkAccessPoliciesList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
