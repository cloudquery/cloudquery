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

func buildAppstreamFleetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.Fleet{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeFleets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeFleetsOutput{
			Fleets: []types.Fleet{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}
func TestAppstreamFleets(t *testing.T) {
	client.AwsMockTestHelper(t, Fleets(), buildAppstreamFleetsMock, client.TestOptions{})
}
