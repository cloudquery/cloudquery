// Auto generated code - DO NOT EDIT.

package web

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createPublishingProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebPublishingProfilesClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			PublishingProfiles: mockClient,
		},
	}

	data := recipes.publishProfile{}
	require.Nil(t, faker.FakeObject(&data))

	result := recipes.NewpublishProfileListResultPage(recipes.publishProfileListResult{Value: &[]recipes.publishProfile{data}}, func(ctx context.Context, result recipes.publishProfileListResult) (recipes.publishProfileListResult, error) {
		return recipes.publishProfileListResult{}, nil
	})

	mockClient.EXPECT().ListPublishingProfileXMLWithSecrets(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
