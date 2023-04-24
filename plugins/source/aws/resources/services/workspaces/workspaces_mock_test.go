package workspaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildWorkspaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var workspace types.Workspace
	if err := faker.FakeObject(&workspace); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeWorkspaces(
		gomock.Any(),
		&workspaces.DescribeWorkspacesInput{},
		gomock.Any(),
	).Return(
		&workspaces.DescribeWorkspacesOutput{Workspaces: []types.Workspace{workspace}},
		nil,
	)

	return client.Services{Workspaces: mock}
}

func TestWorkspacesWorkspaces(t *testing.T) {
	client.AwsMockTestHelper(t, Workspaces(), buildWorkspaces, client.TestOptions{})
}
