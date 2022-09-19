// Auto generated code - DO NOT EDIT.

package web

import (
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

	data := services.PublishingProfiles{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().ListPublishingProfiles(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
