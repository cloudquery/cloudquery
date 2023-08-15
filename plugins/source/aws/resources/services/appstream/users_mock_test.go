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

func buildAppstreamUsersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.User{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeUsersOutput{
			Users: []types.User{object},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildAppstreamUsersMock, client.TestOptions{})
}
