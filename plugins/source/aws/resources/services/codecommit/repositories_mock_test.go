package codecommit

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodecommitClient(ctrl)

	repoMetadata := types.RepositoryMetadata{}
	require.NoError(t, faker.FakeObject(&repoMetadata))

	m.EXPECT().ListRepositories(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codecommit.ListRepositoriesOutput{
			Repositories: []types.RepositoryNameIdPair{{RepositoryName: repoMetadata.RepositoryName, RepositoryId: repoMetadata.RepositoryId}},
		},
		nil,
	)

	m.EXPECT().BatchGetRepositories(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codecommit.BatchGetRepositoriesOutput{
			Repositories: []types.RepositoryMetadata{repoMetadata},
		},
		nil,
	)

	tags := map[string]string{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codecommit.ListTagsForResourceOutput{
			Tags: tags,
		},
		nil,
	)
	return client.Services{Codecommit: m}
}

func TestRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}
