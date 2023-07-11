package workspaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDirectories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var directory types.WorkspaceDirectory
	require.NoError(t, faker.FakeObject(&directory))

	mock.EXPECT().DescribeWorkspaceDirectories(
		gomock.Any(),
		&workspaces.DescribeWorkspaceDirectoriesInput{},
		gomock.Any(),
	).Return(
		&workspaces.DescribeWorkspaceDirectoriesOutput{Directories: []types.WorkspaceDirectory{directory}},
		nil,
	)

	return client.Services{Workspaces: mock}
}

func TestWorkspacesDirectories(t *testing.T) {
	client.AwsMockTestHelper(t, Directories(), buildDirectories, client.TestOptions{})
}
