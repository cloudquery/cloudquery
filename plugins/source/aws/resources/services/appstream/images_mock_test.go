package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAppstreamImagesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.Image{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeImagesOutput{
			Images: []types.Image{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}
func TestAppstreamImages(t *testing.T) {
	client.AwsMockTestHelper(t, Images(), buildAppstreamImagesMock, client.TestOptions{})
}
