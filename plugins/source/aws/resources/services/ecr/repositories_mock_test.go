package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	l := types.Repository{}
	require.NoError(t, faker.FakeObject(&l))

	i := types.ImageDetail{}
	require.NoError(t, faker.FakeObject(&i))

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRepositoriesOutput{
			Repositories: []types.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	tagResponse := ecr.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	iF := ecr.DescribeImageScanFindingsOutput{}
	require.NoError(t, faker.FakeObject(&iF))

	iF.NextToken = nil
	m.EXPECT().DescribeImageScanFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iF, nil)

	repoResponse := ecr.GetRepositoryPolicyOutput{}
	require.NoError(t, faker.FakeObject(&repoResponse))

	policyText := "{}"
	repoResponse.PolicyText = &policyText
	m.EXPECT().GetRepositoryPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&repoResponse, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrRepositoriesMock, client.TestOptions{})
}
