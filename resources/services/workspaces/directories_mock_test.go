package workspaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDirectories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var directory types.WorkspaceDirectory
	if err := faker.FakeData(&directory); err != nil {
		t.Fatal(err)
	}

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
