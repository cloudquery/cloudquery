// Auto generated code - DO NOT EDIT.

package security

import (
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

func TestSecurityPricings(t *testing.T) {
	client.MockTestHelper(t, Pricings(), createPricingsMock)
}

func createPricingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityPricingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Pricings: mockClient,
		},
	}

	data := security.Pricing{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := security.PricingList{Value: &[]security.Pricing{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
