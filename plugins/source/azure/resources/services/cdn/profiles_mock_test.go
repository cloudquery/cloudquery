// Auto generated code - DO NOT EDIT.

package cdn

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

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func TestCDNProfiles(t *testing.T) {
	client.MockTestHelper(t, Profiles(), createProfilesMock)
}

func createProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNProfilesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Profiles: mockClient,
		},
	}

	data := cdn.Profile{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := cdn.NewProfileListResultPage(cdn.ProfileListResult{Value: &[]cdn.Profile{data}}, func(ctx context.Context, result cdn.ProfileListResult) (cdn.ProfileListResult, error) {
		return cdn.ProfileListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
